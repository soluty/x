package ecs

import (
	"sync"
	"sync/atomic"
)

type World struct {
	entities        map[uint64]*Entity
	groupedEntities map[string]*entityGroup // group来访问实体列表, key为所有的group的集合

	systems              []System
	systemsDirtyEntities []*Entity

	idGen uint64
	pool  *sync.Pool
}

func NewWorld() *World {
	w := &World{
		entities:        map[uint64]*Entity{},
		groupedEntities: map[string]*entityGroup{},
		idGen:           0,
	}
	w.pool = &sync.Pool{
		New: func() interface{} {
			return &Entity{
				world: w,
			}
		},
	}
	return w
}

type entityGroup struct {
	entities   []*Entity
	components []string
	name       string
}

func (w *World) CreateEntity() *Entity {
	id := atomic.AddUint64(&w.idGen, 1)
	e := w.pool.Get().(*Entity)
	e.id = id
	e.components = map[string]Component{}
	w.entities[id] = e
	e.setSystemsDirty()
	return e
}

func (w *World) RemoveEntity(id uint64) {
	if entity, ok := w.entities[id]; !ok {
		return
	} else {
		entity.RemoveAllComponents()
		entity.systemsDirty = false
		w.pool.Put(entity)
		delete(w.entities, id)
	}
}

// 拥有多个component的所有实体
func (w *World) QueryComponents(base string, other ...string) []*Entity {
	group, ok := w.groupedEntities[groupKey(base, other...)]

	if !ok {
		group = w._indexGroup(base, other...)
	}

	return group.entities
}

func (w *World) Count() int {
	return len(w.entities)
}

func (w *World) entityAddComponent(e *Entity, name string) {
	for _, group := range w.groupedEntities {
		// Component已经在group中, e 有所有group中的组件，并且不在value中
		if containsComponent(group.components, name) {
			continue
		}
		if !e.HasComponents(group.components) {
			continue
		}
		if containsEntity(group.entities, e) {
			continue
		}
		group.entities = append(group.entities, e)
	}
}

func (w *World) entityRemoveComponent(e *Entity, name string) {
	for _, group := range w.groupedEntities {
		// Component已经在group中, e 有所有group中的组件，并且不在value中
		if !containsComponent(group.components, name) {
			continue
		}
		if !e.HasComponents(group.components) {
			continue
		}
		for idx, value := range group.entities {
			if value == e {
				group.entities = append(group.entities[:idx], group.entities[idx+1:]...)
				break
			}
		}
	}
}

func (w *World) _indexGroup(c string, other ...string) *entityGroup {
	key := groupKey(c, other...)
	e := &entityGroup{
		name:       key,
		components: append([]string{c}, other...),
	}
	w.groupedEntities[key] = e
	for _, entity := range w.entities {
		if entity.HasComponents(e.components) {
			e.entities = append(e.entities, entity)
		}
	}
	return e
}

func groupKey(c string, other ...string) string {
	ret := c
	for _, value := range other {
		ret += "-" + value
	}
	return ret
}

func containsComponent(cs []string, c string) bool {
	for _, v := range cs {
		if v == c {
			return true
		}
	}
	return false
}

func containsEntity(cs []*Entity, c *Entity) bool {
	for _, v := range cs {
		if v == c {
			return true
		}
	}
	return false
}
