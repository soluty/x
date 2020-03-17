package service

import (
	"github.com/golang/protobuf/ptypes/any"
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/pb"
	"github.com/tidwall/gjson"
)

var _ entity.Side = &SideBase{}
var _ entity.Game = &GameBase{}
type SideBase struct {
	player entity.Player
	next   entity.Side
	prev   entity.Side
}

func (s *SideBase) SetPlayer(player entity.Player) {
	s.player = player
}

func (s *SideBase) Player() entity.Player {
	return s.player
}

func (s *SideBase) Init() {
}

func (s *SideBase) Next() entity.Side {
	return s.next
}

func (s *SideBase) Prev() entity.Side {
	return s.prev
}
func (s *SideBase) SetNext(side entity.Side) {
	s.next = side
}

func (s *SideBase) SetPrev(side entity.Side) {
	s.prev = side
}

type GameBase struct {
	room  entity.Room
}

func (g *GameBase) Reset(idx int) {
	panic("implement me")
}

func (g *GameBase) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (g *GameBase) Unmarshal(*gjson.Result) error {
	panic("implement me")
}

func (g *GameBase) GetSide(idx int) entity.Side {
	player := g.Room().GetPlayer(idx)
	if player == nil {
		return nil
	}
	return player.Side()
}

func (g *GameBase) Init() {
}


func (g *GameBase) Receive(idx int, head pb.C2SMsg_Type, body *any.Any) error {
	panic("GameBase implements Receive")
}

func (g *GameBase) SetRoom(room entity.Room) {
	g.room = room
}

func (g *GameBase) Room() entity.Room {
	return g.room
}
