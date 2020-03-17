package qp

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
)

var ModuleNotExist = errors.New("模块不存在")

type appBase struct {
	moduleSchema sync.Map // map[string]reflect.Type  存储所有module的引用，方便重启module
	current      sync.Map // map[string]Module  存储当前正在运行中的模块
	rootCtx      context.Context
	rootCancel   context.CancelFunc

	// 客户端服务发现模块
	Discover Discover
}

func New(name string) *appBase {
	a := &appBase{}
	a.rootCtx, a.rootCancel = context.WithCancel(context.Background())
	return a
}

type canceledModule struct {
	m      Module
	cancel context.CancelFunc
}

func (a *appBase) Get(name string) Module {
	m, ok := a.current.Load(name)
	if !ok {
		return nil
	}
	return m.(Module)
}

// Run方法默认启动所有模块
func (a *appBase) Run(ms ...Module) {
	if err := a.checkModuleName(ms); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, mod := range ms {
		m := mod
		m.setApp(a)
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Create()
		}()
		typ := reflect.TypeOf(m)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		a.moduleSchema.Store(m.Name(), typ)
	}
	wg.Wait()

	for _, mod := range ms {
		m := mod
		wg.Add(1)
		ctx, cancel := context.WithCancel(a.rootCtx)
		a.current.Store(m.Name(), canceledModule{m, cancel})
		go func() {
			var startErr error
			defer func() {
				if err := recover(); err != nil {
					// todo logger
					err = fmt.Errorf("%v", err)
					recoverStop(m.Stop, err.(error))
				} else {
					err = startErr
					if err == nil {
						recoverStop(m.Stop, nil)
					} else {
						recoverStop(m.Stop, err.(error))
					}
				}
				wg.Done()
			}()
			startErr = m.Start(ctx, &wg)
			return
		}()
	}

	go a.shutdownListen()
	wg.Wait()
}

func recoverStop(f func(err error), err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("module stop error!")
		}
	}()
	f(err)
}

func (a *appBase) shutdownListen() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	<-ch
	_ = a.Stop("")
	//appBase.Kill(nil)
}

// 停止一个模快， module为空则是停止所有模块
func (app *appBase) Stop(module string) error {
	if module == "" {
		app.rootCancel()
		return nil
	}
	if _, ok := app.moduleSchema.Load(module); !ok {
		return ModuleNotExist
	} else {
		mod, ok := app.current.Load(module)
		if !ok {
			return errors.New("该模块已经停止")
		}
		app.current.Delete(module)
		m := mod.(canceledModule)
		m.cancel()
		return nil
	}
}

func (a *appBase) checkModuleName(ms []Module) error {
	set := map[string]struct{}{}
	for _, m := range ms {
		if strings.TrimSpace(m.Name()) != m.Name() {
			return fmt.Errorf("模块%v的名字不能包含空格", m.Name())
		}
		if m.Name() == "" {
			return errors.New("模块名字不能为空")
		}
		if _, ok := set[m.Name()]; ok {
			return fmt.Errorf("模块%v名字重复", m.Name())
		}
		set[m.Name()] = struct{}{}
	}
	return nil
}
