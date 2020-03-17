package main

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type m int
func (m)Name()string{
	return "m"
}
func (m)Start(ctx context.Context)  error{
	fmt.Println("i am deamon")
	for {
		select {
		case <-time.After(time.Second):
		}
	}
	return nil
}

type S struct {
	s int
}

func (s *S) String() string {
	return "1"
}

var s fmt.Stringer = &S{1}

func main() {
	var wg sync.WaitGroup
	wg.Wait()
	typ := reflect.TypeOf(s)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	v := reflect.New(typ).Interface().(fmt.Stringer)
	fmt.Println(v.String())
}
