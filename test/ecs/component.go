package ecs

import "reflect"

type Component interface {
	Init(typ reflect.Type, o *Entity)
	Node() *Entity
	Type() reflect.Type
}

var _ Component = &BaseComponent{}

type BaseComponent struct {
	node *Entity
	typ reflect.Type
}

func (this *BaseComponent) Init(typ reflect.Type, o *Entity) {
	this.typ = typ
	this.node = o
}

func (this *BaseComponent) Node() *Entity {
	return this.node
}

func (this *BaseComponent) Type() reflect.Type {
	return this.typ
}


