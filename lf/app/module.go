package app

import (
	"context"
	"net/rpc"
	"reflect"
	"time"
)

// Module 是app的一个子模块，和组件类似，每个Module的Start方法都会在app.Run时启动一个协程运行
// 一个Module代表app中一个正交的功能，可以热重启模块，模块名不可重复
// 模块内部一般是单线程, 模块之间通信一般通过rpc
type Module interface {
	App() *App // 获取app的实例
	Name() string
	OnCreate() // 启动, 顺序执行，允许panic, 不应该返回error
	OnStart(ctx context.Context) error
	OnStop()
	StartOnAppInit() bool // 模块是否在app启动时启动
	start(ctx context.Context) error
}

// room component
// game1 component
// game2 component
// 每一个world都会在etcd上注册

// server
// name   string                 // name of service
// rcvr   reflect.Value          // receiver of methods for the service
// typ    reflect.Type           // type of the receiver
// method map[string]*methodType // registered methods

type Call struct {
	Module string
	ServiceMethod string      // The name of the service and method to call.
	Args          interface{} // The argument to the function (*struct).
	Reply         interface{} // The reply from the function (*struct).
	Error         error       // After completion, the error status.
	Done          chan *Call  // Strobes when call is complete.
}

// 模块间通信, resp必须为指针
func (app *App) Call(module string, method string, req interface{}, resp interface{}) error {

	return nil
}

type Ticker interface {
	Tick()
}

var _ Module = &BaseModule{}

type BaseModule struct {
	ch     chan *Call
	ticker Ticker
	parent Module
}

func (m *BaseModule) OnCreate() {
}

func (m *BaseModule) OnStart(ctx context.Context) error {
	return nil
}

func (m *BaseModule) OnStop() {
}

func (m *BaseModule) StartOnAppInit() bool {
	return true
}

func (m *BaseModule) start(ctx context.Context) error {
	rpc.NewServer()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case rpcReq := <-m.ch:
			m.App().rpcReq.Module  // World  Entity  Component
			v := reflect.ValueOf(m.parent)
			method := v.MethodByName(rpcReq.method)
		}
	}

	// Tick ->

	err := m.OnStart(ctx)
	if err != nil {

	}
	m.OnStop()
	return err
}

type modRpc struct {
	method string
	args   interface{}
	reply  interface{}
}

func (m *BaseModule) Name() string {
	return ""
}

func (m *BaseModule) Start(ctx context.Context) error {
	for {
		select {
		case <-m.ch:
		default:
		}
		if m.ticker != nil {
			m.ticker.Tick()
		}
		time.Sleep(time.Millisecond)
	}
	return nil
}
