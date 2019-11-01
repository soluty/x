package ecs

type System interface {
	Name() string
	World() *World
	Init(name string, w *World)
}

type SystemBase struct {

}
