package ecs

type Component interface {
	Init(name string, o *Entity)
	Node() *Entity
	Name() string
}

type BaseComponent struct {
	node *Entity
	name string
}

func (c *BaseComponent) Init(name string, o *Entity) {
	c.name = name
	c.node = o
}

func (c *BaseComponent) Node() *Entity {
	return c.node
}

func (c *BaseComponent) Name() string {
	return c.name
}
