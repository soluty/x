package fy

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
)

var ModuleNotExist = errors.New("模块不存在")

type App struct {
	moduleSchema sync.Map // map[string]reflect.Type  存储所有module的引用，方便重启module
	current      sync.Map // map[string]Module  存储当前正在运行中的模块
	rootCtx      context.Context
	rootCancel   context.CancelFunc
	state        int // app的状态
}

func New() *App {
	app := &App{}
	app.rootCtx, app.rootCancel = context.WithCancel(context.Background())
	//tomb.WithContext(context.Background())
	return app
}

type canceledModule struct {
	m      Module
	cancel context.CancelFunc
}

// Run方法默认启动所有模块
func (app *App) Run(appName string, ms ...Module) error {
	if err := app.checkModuleName(ms); err != nil {
		return err
	}
	// app onCreate
	var wg sync.WaitGroup
	for _, mod := range ms {
		m := mod
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.OnCreate()
		}()
		typ := reflect.TypeOf(m)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		app.moduleSchema.Store(m.Name(), typ)
	}
	wg.Wait()

	// start and stop
	for _, mod := range ms {
		m := mod
		wg.Add(1)
		ctx, cancel := context.WithCancel(app.rootCtx)
		app.current.Store(m.Name(), canceledModule{m, cancel})
		go func() {
			var err interface{}
			defer func() {
				if err = recover(); err != nil {
					// todo logger
					err = fmt.Errorf("%v", err)
					log.Println(err)
					m.OnStop(err.(error))
				} else {
					m.OnStop(nil)
				}
				wg.Done()
			}()
			err = m.OnStart(ctx)
			return
		}()
	}
	go app.shutdownListen()
	wg.Wait()
	return nil
}

func (app *App) shutdownListen() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	<-ch
	_ = app.Stop("")
	//app.Kill(nil)
}

// 停止一个模快， module为空则是停止所有模块
func (app *App) Stop(module string) error {
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

func (app *App) checkModuleName(ms []Module) error {
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
