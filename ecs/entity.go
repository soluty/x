package ecs

import (
	"log"
	"reflect"
	"strings"
)

// Entity 不能通过平常的方法创建，必须通过World.CreateEntity来创建
type Entity struct {
	world      *World
	components map[string]Component
	id         uint64

	systems []string // 哪些system持有该entity的引用
	systemsDirty bool //
}

func (e *Entity) ID() uint64 {
	return e.id
}

func (e *Entity) AddComponent(c Component, name ...string) *Entity {
	cname := ""
	if len(name) > 0 {
		cname = name[0]
	}
	if cname == "" {
		cname = reflect.TypeOf(c).String()
		cname = strings.TrimPrefix(cname, "*")
		ss := strings.Split(cname, ".")
		cname = ss[len(ss)-1]
	}
	if e.HasComponent(cname) {
		log.Printf("[warn] ecs entity %v AddComponent %v duplicate!\n", e.id, cname)
		return e
	}
	c.Init(cname, e)
	e.components[cname] = c
	e.world.entityAddComponent(e, cname)
	return e
}

func (e *Entity) RemoveComponent(name string) *Entity {
	if !e.HasComponent(name) {
		return e
	}
	e.world.entityRemoveComponent(e, name)
	delete(e.components, name)
	return e
}

func (e *Entity) GetComponent(name string) Component {
	return e.components[name]
}

func (e *Entity) HasComponent(name string) bool {
	for _, value := range e.components {
		if value.Name() == name {
			return true
		}
	}
	return false
}

// 尝试通过reflect来确定组件
func (e *Entity) HasComponent2(name string) bool {
	// reflect.ValueOf(e).FieldByName("")
	return false
}

func (e *Entity) HasComponents(names []string) bool {
	for _, name := range names {
		if !e.HasComponent(name) {
			return false
		}
	}
	return true
}

func (e *Entity) RemoveAllComponents() {
	for key := range e.components {
		e.RemoveComponent(key)
	}
}

func (e *Entity) Destroy() {
	e.world.RemoveEntity(e.id)
}

func (e *Entity) setSystemsDirty() {
	if !e.systemsDirty && e.world != nil {
		e.systemsDirty = true
		e.world.systemsDirtyEntities = append(e.world.systemsDirtyEntities, e)
	}
}
