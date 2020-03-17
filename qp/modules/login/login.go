package login

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/soluty/x/qp"
	"github.com/soluty/x/qp/pb"
	"net/http"
	"sync"
)

type ModuleLogin struct {
	qp.ModuleBase
	selector qp.Selector
}

func (m *ModuleLogin) Name() string {
	msg := pb.S2CMsg{}
	if msg.Err == nil {

	}
	return "login"
}

func (m *ModuleLogin) Create() {
	fmt.Println("login create")
	m.selector = nil
}

func (m *ModuleLogin) Stop(err error) {
	fmt.Println("login stop", err)
}

func (m *ModuleLogin) Start(ctx context.Context, wg *sync.WaitGroup) error {
	fmt.Println("login start")
	loginService := pb.NewLoginServer(m, nil)
	server := &http.Server{Addr: ":8181", Handler: loginService}
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.ListenAndServe()
		fmt.Println("login service done")
	}()
	<-ctx.Done()
	server.Shutdown(context.Background())
	return nil
}

// todo real login
func (m *ModuleLogin) getUid(req *pb.LoginReq) (string, error) {
	return req.PlatId, nil
}

// 分配gateUrl
// 客户端去连接
// 发现之前有  kick掉
//

// 返回token uid 和 url，客户端通过uid记录自己的唯一id，通过gateUrl和token去找到网关登录
func (m *ModuleLogin) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	// 1. 确认登录的用户名和密码正确
	uid, err := m.getUid(req)
	if err != nil {
		return nil, err
	}
	discover := m.App().Discover
	mutex := concurrency.NewMutex(&concurrency.Session{}, "player/"+uid)
	mutex.Lock(ctx)
	// 分配给玩家一个外网的gateUrl
	gateUrl, err := discover.AssignPlayerGate(uid)
	if err != nil {
		return nil, err
	}
	// 内网url
	oldGate, err := discover.GetPlayerGateUrl(uid)
	if err != nil {
		return nil, err
	}
	if oldGate != "" {
		gateClient := pb.NewGateProtobufClient(oldGate, http.DefaultClient)
		_, err = gateClient.Kick(ctx, &pb.KickReq{
			Reason: pb.KickReq_LoginConnect,
		})
		if err != nil {

		}
	}
	return &pb.LoginRes{
		Token:   "abc",
		GateUrl: gateUrl,
	}, nil
	// /gateUrl/playerId  -> gateUrl
	// 首先查询etcd, 之前有没有连接，如果有连接还在连接，则踢掉

	//lastGate, err := m.selector.GetPlayerLastGate(req.PlatId)
	//if err != nil {
	//	return nil, err
	//}
	// 根据各个gateUrl的负载
	// /conf/
	// /payload/gate/${gateurl} /payload/room/${roomurl}
	// /conn/${playerid} -> url
	//

	//gateUrl, lastGate, err := m.selector.SelectGateServer(uid)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if lastGate != "" {
	//	gateClient := pb.NewGateJSONClient(lastGate, http.DefaultClient)
	//	_, err = gateClient.Kick(ctx, &pb.KickReq{
	//		Reason: pb.KickReq_LoginConnect,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//}
}

// ack  ->  uid,
func (m *ModuleLogin) LoginAck(context.Context, *pb.LoginAckReq) (*pb.LoginAckRes, error) {
	//discover := m.App().Discover
	return nil, nil
}
