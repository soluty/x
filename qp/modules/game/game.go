package game

import (
	"context"
	"fmt"
	"github.com/soluty/x/qp"
	"github.com/soluty/x/qp/pb"
	"net/http"
	"sync"
	"time"
)

type ModuleGame struct {
	qp.ModuleBase
	rooms sync.Map
}

func (m *ModuleGame) Name() string {
	return "game"
}

func (m *ModuleGame) Create() {

}

func (m *ModuleGame) Stop(reason error) {

}

type Room struct {
}

func (m *ModuleGame) getRoom(roomId string) *Room {
	return nil
}

//
func (m *ModuleGame) ReceiveMsg(ctx context.Context, req *pb.C2SMsg) (*pb.S2CMsg, error) {
	fmt.Println("receive req")
	m.getRoom(req.RoomId)
	go func() {
		<-ctx.Done()
		fmt.Println("down context")
		//panic("hehe")
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("send res---")
	return &pb.S2CMsg{
		Head:pb.S2CMsg_Login,
	}, nil
}

//
func (m *ModuleGame) CreateRoom(context.Context, *pb.C2SMsg) (*pb.S2CMsg, error) {
	return nil, nil
}

func (m *ModuleGame) Start(ctx context.Context, wg *sync.WaitGroup) error {
	fmt.Println("game start")
	gameService := pb.NewGameServer(m, nil)
	server := &http.Server{Addr: ":8182", Handler: gameService}
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.ListenAndServe()
		fmt.Println("game service done")
	}()
	<-ctx.Done()
	server.Shutdown(context.Background())
	return nil
}
