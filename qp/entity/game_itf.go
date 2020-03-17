package entity

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/soluty/x/qp/pb"
	"github.com/tidwall/gjson"
)

// 实际的游戏
type Game interface {
	Init() // 构造方法
	Reset(idx int) // 每一局开始游戏时调用，idx从1开始
	SetRoom(r Room)
	Room() Room
	// Handle(playerIndex int, head pb.C2SMsg_Type, body proto.Message) error
	Receive(idx int, head pb.C2SMsg_Type, body *any.Any) error // 处理外界接收到的消息
	GetSide(int) Side
	json.Marshaler
	Unmarshal(*gjson.Result) error
}

type Side interface {
	Init() // 构造方法
	SetPlayer(Player)
	Player() Player  // get player send msg
	Next() Side // 下家
	Prev() Side // 上家
	SetNext(Side)
	SetPrev(Side)
}

// 房间，与外界交流
type Room interface {
	SetGame(g Game)
	Game() Game
	BlackBoard() BlackBoard
	Receive(msg *pb.C2SMsg) error // 处理外界接收到的消息
	GetPlayer(int) Player
	Players() []Player // 获取所有玩家数组
	GameType() pb.GameType
	PlayerCount() int
}

// 玩家，与外界交流
type Player interface {
	SetSide(Side)
	Side() Side
	Index() int
	Send(head pb.S2CMsg_Type, body proto.Message) // 给外界发送消息
}

type BlackBoard map[string]interface{}

func (b BlackBoard) Get(key string) interface{} {
	return b[key]
}

func (b BlackBoard) Set(key string, v interface{}) {
	b[key] = v
}
