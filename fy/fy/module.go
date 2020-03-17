package fy

import (
	"context"
)

// Module 定义基本的生命周期
type Module interface {
	Name() string
	OnCreate()
	OnStart(ctx context.Context) error
	OnStop(reason error)
}

type ModuleBase struct {

}

func (ModuleBase) OnCreate() {

}
func (ModuleBase) OnStop(reason error) {

}
// debug module pprof
// config module


//type RoomModule struct {
//	ctx     *actor.RootContext
//	idToPid sync.Map
//}
//
//func (r *RoomModule) Name() string {
//	return "room"
//}
//
//func (r *RoomModule) Send(roomId string, msg interface{}) {
//	s, ok := r.idToPid.Load(roomId)
//	if !ok {
//		return
//	}
//	pid := s.(*actor.PID)
//	r.ctx.Send(pid, msg)
//}
//
//func (r *RoomModule) OnCreate() {
//	r.ctx = actor.NewRootContext(nil)
//}
//
//func (r *RoomModule) OnStart(ctx context.Context) error {
//	r.CreateRoom("abc")
//	<-ctx.Done()
//	return ctx.Err()
//}
//
//func (r *RoomModule) OnStop(reason error) {
//
//}
//
//func (r *RoomModule) CreateRoom(id string) {
//	room := NewRoom(id)
//	pid := r.ctx.Spawn(actor.PropsFromFunc(room.Receive))
//	r.idToPid.Store(id, pid)
//
//	r.Send("abc", 1111)
//}
//
//var _ Module = &RoomModule{}
//
//func NewRoom(id string) *Room {
//	return &Room{
//		id: id,
//	}
//}
//
//type Room struct {
//	id string
//}
//
//func (r *Room) Receive(ctx actor.Context) {
//	switch msg := ctx.Message().(type) {
//	case int:
//		fmt.Println(msg)
//	case *actor.Started:
//		fmt.Println(r.id, "start")
//		time.Sleep(time.Second*5)
//	case *actor.Stopped:
//		fmt.Println(r.id, "stop")
//	default:
//		log.Println("msg not register",msg)
//	}
//}
