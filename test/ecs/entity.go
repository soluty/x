package ecs

import "reflect"

// Entity 不能通过平常的方法创建，必须通过World.CreateEntity来创建
type Entity struct {
	world      *World
	components []Component
	tags       []string
	id         uint64
}

func (this *Entity) AddComponent(c Component) *Entity{
	if this.HasComponent(c.Type()) {
		return this
	}
	c.Init(reflect.TypeOf(c), this)
	this.components = append(this.components, c)
	this.world.entityAddComponent(this, c)
	return this
}

func (this *Entity) RemoveComponent(c reflect.Type) *Entity{
	return this
}


func (this *Entity) HasComponent(c reflect.Type) bool {
	for _, value := range this.components {
		if value.Type() == c {
			return true
		}
	}
	return false
}

func (this *Entity) HasComponents(cs []reflect.Type) bool {
	for _, value := range cs {
		if !this.HasComponent(value) {
			return false
		}
	}
	return true
}



func (this *Entity) HasTag (tag string) bool {
	for _, value := range this.tags {
		if value == tag {
			return true
		}
	}
	return false
}

func (this *Entity) AddTag (tag string)  {

}

func (this *Entity) RemoveTag (tag string)  {

}

func (this *Entity) Destroy ()  {
	this.world.RemoveEntity(this.id)
}