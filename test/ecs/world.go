package ecs

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type entityGroup struct {
	entities []*Entity
	components []Component
	name string
}

func getComponentsName(cs []Component) string{
	ret := ""
	for _, value := range cs {
		ret += value.Type().String()+"-"
	}
	return ret
}

type World struct {
	entities map[uint64]*Entity
	tagEntities map[string][]*Entity  // tag来访问实体列表
	groups map[string]*entityGroup    // group来访问实体列表, key为所有的group的集合

	systems  []System
	idGen    uint64
	pool     *sync.Pool
}

func NewWorld() *World {
	this := &World{
		entities: map[uint64]*Entity{},
		systems:  nil,
		idGen:    0,
	}
	this.pool = &sync.Pool{
		New: func() interface{} {
			return &Entity{
				world: this,
			}
		},
	}
	return this
}

func (this *World) CreateEntity() *Entity {
	id := atomic.AddUint64(&this.idGen, 1)
	e := this.pool.Get().(*Entity)
	e.id = id
	e.components = nil
	this.entities[id] = e
	return e
}

func (this *World) RemoveEntity(id uint64) {
	this.pool.Put(this.entities[id])

	delete(this.entities, id)
}

func (this *World) RegisterSystem(s System) error {
	for _, value := range this.systems {
		if value.Name() == s.Name() {
			return errors.New("system name cf: " + s.Name())
		}
	}
	this.systems = append(this.systems, s)
	return nil
}

// 拥有多个component的所有实体
func (this *World) QueryComponents(base Component, other ...Component) []*Entity {

	return nil
}

func (this *World) QueryTag(tag string) []*Entity {
	return nil
}

func (this *World) Count() int {
	return len(this.entities)
}

func (this *World) Run(t time.Duration) {

}

func (this *World) entityAddComponent (e *Entity, c Component) {

}