package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/pb"
	"strconv"
)

var _ entity.Player = &Player{}

type Player struct {
	Id         string
	Ready      byte
	Disconnect bool

	index int
	side  entity.Side

	// 响应的channel chan 1 即可
	respChan chan *pb.S2CMsg
}

func NewPlayer(id string) *Player {
	player := &Player{
		Id:       id,
		respChan: make(chan *pb.S2CMsg, 1),
	}
	go player.respLoop()
	return player
}

func (p *Player) respLoop() {
	for msg := range p.respChan {
		p.send(msg)
	}
}

func (p *Player) SetSide(side entity.Side) {
	p.side = side
	side.SetPlayer(p)
}

func (p *Player) Side() entity.Side {
	return p.side
}

func (p *Player) Index() int {
	return p.index
}

var tm = proto.TextMarshaler{Compact: true, ExpandAny: true}

func (p *Player) send(msg *pb.S2CMsg) {
	//var v interface{}
	//err := json.Unmarshal(msg.Body.Value, v)
	//if err != nil {
	//	panic(err)
	//}
	// 搞死人了，这个u8转换
	//s := msg.Body.String()
	s := tm.Text(msg.Body)
	s, _ = strconv.Unquote(`"` + s + `"`)
	fmt.Printf("给index=%v玩家%v发送消息%v: %s\n", p.index, p.Id, msg.Head, s)

}

func (p *Player) Send(head pb.S2CMsg_Type, body proto.Message) {
	any, err := ptypes.MarshalAny(body)
	if err != nil {
		panic(err)
	}
	p.respChan <- &pb.S2CMsg{
		Head: head,
		Body: any,
	}
}
