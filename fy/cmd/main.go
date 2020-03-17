package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/gogf/gf/net/ghttp"
	"github.com/soluty/x/fy/fy"
	"github.com/soluty/x/fy/fy/modules"
	"github.com/soluty/x/fy/hello"
	"go.etcd.io/etcd/etcdserver/api/etcdhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type m1 struct {
	fy.ModuleBase
}

func (m1) Name() string {
	return "m1"
}


func (m m1) OnStart(ctx context.Context) error {
	cli, _ := clientv3.NewFromURL("localhost:2379")
	//cli.KV = namespace.NewKV(cli.KV, "a/")
	//cli.Watcher = namespace.NewWatcher(cli.Watcher, "a/")
	//cli.Lease = namespace.NewLease(cli.Lease, "a/")
	etcdhttp.HandleBasic()
	rsp, err := cli.Delete(ctx, "/a", clientv3.WithPrevKV())
	fmt.Println(err,	rsp.PrevKvs)
	fmt.Println("start m1")
	//go func() {
	//	for rsp := range cli.Watch(ctx, "s", clientv3.WithRev(1)) {
	//		fmt.Println(len(rsp.Events))
	//	}
	//}()
	// 但由QUIT字符（通常是Ctrl+\）来控制。进程因收到SIGQUIT退出时会产生core文件，在这个意义上类似于一个程序错误信号。
	// net,  ctrl+c ,  kill , ctrl+/
	// syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	// 配置中心 -> 配置

	// 事务
	//rsp2, _ := concurrency.NewSTM(cli, func(stm concurrency.STM) error {
	//	stm.Del("")
	//	return nil
	//})

	s := ghttp.GetServer()
	s.SetPort(8080)
	s.EnablePProf()
	ch := make(chan struct{})
	go func() {
		s.Run()
		ch <- struct{}{}
	}()
	<-ctx.Done()
	s.Shutdown()
	<-ch
	return ctx.Err()
}


type m2 string

func (m2) Name() string {
	return "m2"
}

func (m2) Say(ctx context.Context, req *hello.HelloReq) (res *hello.HelloRes, err error){
	res = &hello.HelloRes{
		Message:"hi " + req.Name,
	}
	return
}

func (m2) OnCreate() {
	fmt.Println("create m2")
}

func (m m2) OnStart(ctx context.Context) error {
	fmt.Println("start m2")
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, m)

	lis, err := net.Listen("tcp", ":7171")
	if err != nil {
		log.Fatal("err..")
	}
	go func() {
		log.Println("serve hello server in 7171")
		s.Serve(lis)
	}()

	//
	//mm := cmux.New(lis)
	//mm.Match(cmux.HTTP2HeaderFieldPrefix("content-type", "application/grpc"))
	//mm.Match(cmux.HTTP1HeaderField("Upgrade", "websocket"))
	//
	//mm.Match(cmux.HTTP2())
	//mm.Serve()
	//cmux.PrefixMatcher()

	time.Sleep(2 * time.Second)
	c , err := grpc.Dial("localhost:7171",grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("err1..")
	}
	cli := hello.NewHelloClient(c)
	res, err := cli.Say(ctx, &hello.HelloReq{Name:"lv"})
	if err != nil {
		log.Fatal("err2..")
	}
	fmt.Println(res.Message)
	<-ctx.Done()
	s.GracefulStop()
	// hello.NewHelloClient()
	return ctx.Err()
}

func (m2) OnStop(reason error) {
	fmt.Println("stop m2")
}

var _ fy.Module = new(m1)

func main() {
	app := fy.New()

	//go func() {
	//	time.Sleep(time.Second)
	//	app.Stop("m1")
	//}()

	err := app.Run("abc", new(m1), new(m2), new(modules.InputModule))
	if err != nil {
		fmt.Println(err)
	}
	return
}
