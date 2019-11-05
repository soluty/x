package ecs

import "reflect"

type AComp struct {
	*BaseComponent
}

type BComp struct {
	*BaseComponent
}

type MyEntity struct {
	*AComp
	*BComp
}

func _() {
	e := &MyEntity{}
	reflect.ValueOf(e).FieldByName("AComp")
}
