package qp

import (
	"context"
	"github.com/soluty/x/qp/pb"
	"golang.org/x/sync/semaphore"
	"sync"
)

// Module 定义基本的生命周期, 每一个模块都可以动态的停止, 模块可以是插件，也可以不是
type Module interface {
	Name() string
	Create()
	Start(ctx context.Context, wg *sync.WaitGroup) error
	Stop(reason error)
	App() *appBase
	setApp(*appBase)
}

type M2 interface {
	Run() error
	Exit(error)
}

type ModuleBase struct {
	app *appBase
}

func (m *ModuleBase) App() *appBase {
	return m.app
}

func (m *ModuleBase) setApp(app *appBase) {
	m.app = app
}

// 服务发现接口, 代理etcd
type Discover interface {
	Get(string) string
	GetPlayerGateUrl(uid string) (string, error) // 获取玩家的内网连接url
	GetPlayerGameUrl(uid string) (string, error) // 获取玩家所在游戏服的url
	// 给一个登录成功的玩家分配一个url, 分配完成以后，客户端拿着分配完成的url去连接服务器
	// 不管是否连接上，都给login服务器发一个LoginAck, 如果连上了，并且登录通过，login则更新PlayerGateUrl
	AssignPlayerGate(uid string) (string, error)
	AckPlayerGate(uid string)
	OnStop()

	// login服务器的分布式锁, 防止登录过程中有其它操作
	LockLogin(uid string) error
	GetLoginLocker() (sync.Locker, error)
	UnlockLogin(uid string)

	GetLocker(res string) (ContextLocker, error)
}

// 可超时锁
type ContextLocker interface {
	Lock(ctx context.Context) error  // 立刻返回，可以由ctx取消
	Unlock()  // 没有lock住的时候多次调用不崩溃
	IsLock() bool // 返回是否lock
}


type MemContextLocker struct {
	mu   sync.Mutex
	s    semaphore.Weighted
	lock int32
}

// (player, conn)
var playerConn sync.Map

func (l *MemContextLocker) Lock(ctx context.Context) error {
	//var wg sync.WaitGroup
	//ll, _ := net.ListenTCP()
	//for  {
	//	ll.SetDeadline(time.Now().Add(time.Second))
	//	conn, _ := ll.Accept()
	//	wg.Add(1)
	//	playerConn.Store(1,conn)
	//	go func() {
	//		defer wg.Done()
	//		for {
	//			conn.Read()
	//			wg.Add(1)
	//			// 通过id分段，找到游戏消息，还是其他消息  gate  createRoom
	//			// 路由，
	//			// router processRead()
	//		}
	//	}()
	//}
	//
	//context.WithTimeout()
	//
	//ll.Close()
	//
	//s,_ := concurrency.NewSession()
	//concurrency.NewMutex(s, "/aa")
	//concurrency.NewLocker()
	//b := l.s.TryAcquire(1)
	//if !b {
	//	return false
	//}
	//atomic.CompareAndSwapInt32(&l.lock, 0,1)
	//defer l.s.Release(1)
	//go func() {
	//	<-ctx.Done()
	//	l.Unlock()
	//}()
	//return true
	return nil
}

func (l *MemContextLocker) Unlock() {
	l.s.Release(1)
}

func (l *MemContextLocker) IsLock() bool {
	return false
}

var _ ContextLocker = &MemContextLocker{}

// 选择一个网关给客户端
type Selector interface {
	SelectGateServer(uid string) (newGate string, oldGate string, err error)
	// GetPlayerLastGate(uid string) (string, error)
	SelectGameServer(uid string, typ pb.GameType) (string, error)
}
