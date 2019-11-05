package ecs

type UpdateSystem interface {
	Update()
}

type System interface {
	Name() string
	World() *World
	Init(name string, w *World, super System)
	Update(e *Entity)
}

type BaseSystem struct {
	name string
	world *World
	entities []*Entity
	super System
}

func (s *BaseSystem) Test(e *Entity) bool {
	return false
}
func (s *BaseSystem) Name() string {
	return s.name
}

func (s *BaseSystem) World() *World {
	return s.world
}

func (s *BaseSystem) Init(name string, w *World, super System) {
	s.world = w
	s.name = name
	s.super = super
}

func (s *BaseSystem) AddEntity(e *Entity)  {
	s.entities= append(s.entities, e)
}

func (s *BaseSystem) RemoveEntity(e *Entity)  {
	// todo
}
func (s *BaseSystem) UpdateAll()  {
	for _, value := range s.entities {
		s.super.Update(value)
	}
}
func (s *BaseSystem) Update(e *Entity)  {

}