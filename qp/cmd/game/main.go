package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/soluty/x/qp"
	"github.com/soluty/x/qp/config"
	"github.com/soluty/x/qp/modules/game"
	"github.com/soluty/x/qp/modules/game/games/poker/simple"
	"github.com/soluty/x/qp/modules/game/service"
	"github.com/soluty/x/qp/modules/login"
	"github.com/soluty/x/qp/pb"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

func getAny(message proto.Message) *any.Any {
	a, err := ptypes.MarshalAny(message)
	if err != nil {
		panic(err)
	}
	return a
}

func main() {
	//funcMain()
	service.Register(pb.GameType_Poker_Simple, reflect.TypeOf(simple.Game{}), reflect.TypeOf(simple.Side{}))
	room := service.NewRoom(pb.GameType_Poker_Simple, "r")
	err := room.Receive(&pb.C2SMsg{
		PlayerId: "p1",
		RoomId:   "r",
		Head:     pb.C2SMsg_EnterRoom,
		Body: getAny(&pb.C2SMsg_EnterRoomReq{
			Id: "r",
		}),
	})
	if err != nil {
		panic(err)
	}
	err = room.Receive(&pb.C2SMsg{
		PlayerId: "p2",
		RoomId:   "r",
		Head:     pb.C2SMsg_EnterRoom,
		Body: getAny(&pb.C2SMsg_EnterRoomReq{
			Id: "r",
		}),
	})
	if err != nil {
		panic(err)
	}
	err = room.Receive(&pb.C2SMsg{
		PlayerId: "p1",
		RoomId:   "r",
		Head:     pb.C2SMsg_StartGame,
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
	//var m = map[string]interface{}{}
	//err = json.Unmarshal(testBs, &m)
	//fmt.Println(err)
	//fmt.Println(m)
}

func funcMain() {
	config.Init(nil)
	app := qp.New("")
	var loginModule = &login.ModuleLogin{}
	var gameModule = &game.ModuleGame{}
	go func() {
		time.Sleep(time.Second)
		urlpy, _ := url.Parse("http://127.0.0.1:9090")
		c := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(urlpy),
			}}
		gameC := pb.NewGameJSONClient("http://127.0.0.1:8080", c)
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		fmt.Println("start request game")
		res, err := gameC.ReceiveMsg(ctx, &pb.C2SMsg{
			Head: pb.C2SMsg_Login,
		})
		//go func() {
		//	time.Sleep(100*time.Millisecond)
		//	cl()
		//}()
		fmt.Println(res, err)
	}()
	app.Run(loginModule, gameModule)
}
