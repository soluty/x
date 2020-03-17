package app

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/tomb.v2"
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
	*tomb.Tomb
	moduleCtx context.Context
	once      sync.Once
	state     int // app的状态
}

func New() *App {
	app := &App{}
	app.Tomb, app.moduleCtx = tomb.WithContext(context.Background())
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
	for _, m := range ms {
		if m.StartOnAppInit() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				m.OnCreate()
			}()
		}
		typ := reflect.TypeOf(m)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		app.moduleSchema.Store(m.Name(), typ)
	}
	wg.Wait()

	// sync.WaitGroup errgroup.Group 都没有tomb包好用
	// g,c  := errgroup.WithContext()
	// g.Go()
	// g.Wait()
	for _, m := range ms {
		ctx, cancel := context.WithCancel(app.moduleCtx)
		app.current.Store(m.Name(), canceledModule{m, cancel})
		app.Go(func() (err error) {
			defer func() {
				if err := recover(); err != nil {
					// todo logger
					err = fmt.Errorf("%v", err)
				}
			}()
			// start 封装了OnStart和OnStop
			err = m.start(ctx)
			// app.Tomb.Err() 还是没有reason
			return
		})
	}
	go app.shutdownListen()
	return app.Wait()
}

func (app *App) shutdownListen() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	<-ch
	app.Kill(nil)
}

// 停止一个模块
func (app *App) Stop(module string) error {
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

func (app *App) Kill(reason error) {
	app.once.Do(func() {
		//for _, m := range app.modules {
		//	m.OnStop()
		//}
		app.Tomb.Kill(reason)
	})
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
