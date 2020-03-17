package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/soluty/x/qp/config"
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/pb"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strconv"
	"time"
)

type RoomData struct {
	Id        string      // 唯一id
	Key       string      // 6位key, 同一时间段内唯一
	Type      pb.GameType // 游戏类型
	Creator   string      // 创建者
	CreatedAt time.Time   // 创建时间
	Expire    time.Time   // 过期时间，0值代表永不过期
	players   []*Player   // 加入房间的玩家
	Round     byte        // 当前游戏的round，为0代表游戏没开始
}

func NewRoom(typ pb.GameType, id string) *Room {
	cfg := GetGameConfig(typ)
	room := &Room{
		blackBoard: entity.BlackBoard{},
		RoomData: RoomData{
			Id:      id,
			Type:    typ,
			players: make([]*Player, cfg.Max),
		},
	}
	return room
}

func init() {
	var _ entity.Room = &Room{}
}

// Room代表一个房间，玩家在房间内进行一场Game   room game player side
type Room struct {
	RoomData
	game       entity.Game // 正在进行的游戏，为""代表没有进行游戏
	blackBoard entity.BlackBoard

	playerItfs []entity.Player

	createTimer  *time.Timer
	createCancel context.CancelFunc
	msgChan      chan *pb.C2SMsg
	roomHandles  map[pb.C2SMsg_Type]Handler
	cancel       context.CancelFunc
}

func (r *Room) Receive(msg *pb.C2SMsg) error {
	isGameMsg, err := r.preCheckMsg(msg)
	if err != nil {
		return err
	}
	if isGameMsg {
		idx := r.getPlayerIndex(msg.PlayerId)
		if idx == -1 {
			return errors.New("no player")
		}
		if r.game == nil {
			return errors.New("todo")
		} else {
			return r.game.Receive(idx, msg.Head, msg.Body)
		}
	}
	switch msg.Head {
	case pb.C2SMsg_EnterRoom:
		req := &pb.C2SMsg_EnterRoomReq{}
		err := ptypes.UnmarshalAny(msg.Body, req)
		if err != nil {
			return err
		}
		return r.enterPlayer(msg.PlayerId)
	case pb.C2SMsg_LeaveRoom:
		req := &pb.C2SMsg_LeaveRoomReq{}
		err := ptypes.UnmarshalAny(msg.Body, req)
		if err != nil {
			return err
		}
		return r.leavePlayer(msg.PlayerId)
	case pb.C2SMsg_StartGame:
		return r.startGame()
	default:
		return errors.New("no room handler")
	}
}

func (r *Room) GetPlayer(idx int) entity.Player {
	return r.players[idx]
}

func (r *Room) GameType() pb.GameType {
	return r.Type
}

func (r *Room) SetGame(game entity.Game) {
	r.game = game
	game.SetRoom(r)
}

func (r *Room) Game() entity.Game {
	return r.game
}

func (r *Room) BlackBoard() entity.BlackBoard {
	return r.blackBoard
}

func (r *Room) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &r.RoomData)
	if err != nil {
		return err
	}
	r.game, err = CreateGame(r.RoomData.Type)
	if err != nil {
		return err
	}
	rs := gjson.GetBytes(bytes, "game")
	err = r.game.Unmarshal(&rs)
	return err
}

func (r *Room) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(r.RoomData)
	if err != nil {
		return nil, err
	}
	return sjson.SetBytes(bs, "game", r.game)
}

type Handler func(any *any.Any) error

func (r *Room) getPlayerIndex(playerId string) int {
	for idx, value := range r.players {
		if value != nil && playerId == value.Id {
			return idx
		}
	}
	return -1
}

func (r *Room) getPlayer(id string) *Player {
	for _, value := range r.players {
		if value != nil && id == value.Id {
			return value
		}
	}
	return nil
}

func (r *Room) PlayerCount() int {
	var count int
	for _, value := range r.players {
		if value != nil {
			count++
		}
	}
	return count
}

func (r *Room) inputLoop(ctx context.Context) {
	for {
		select {
		case msg := <-r.msgChan:
			err := r.Receive(msg)
			if err != nil {
				r.getPlayer(msg.PlayerId).respChan <- &pb.S2CMsg{
					Cid:  msg.Cid,
					Head: pb.S2CMsg_Confirm,
					Err: &pb.Error{
						Code: 1,
					},
				}
			}
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func (r *Room) SendTo(playerId string, head pb.S2CMsg_Type, body proto.Message) {
	player := r.getPlayer(playerId)
	if player == nil {
		// todo
		return
	}
	msg, _ := ptypes.MarshalAny(body)
	player.respChan <- &pb.S2CMsg{
		Head:     head,
		Body:     msg,
		PlayerId: playerId,
	}
}

type GameConfig struct {
	Min int
	Max int
}

func GetGameConfig(gameType pb.GameType) GameConfig {
	return GameConfig{
		Min: 2,
		Max: 3,
	}
}

// handle C2SMsg_EnterRoom
func (r *Room) enterPlayer(playerId string) error {
	cfg := GetGameConfig(r.GameType())
	currentPlayerCount := r.PlayerCount()
	if currentPlayerCount >= cfg.Max {
		return errors.New("人满了")
	}
	if r.getPlayerIndex(playerId) != -1 {
		return errors.New("已经在房间")
	}
	for idx, value := range r.players {
		if value == nil {
			r.players[idx] = NewPlayer(playerId)
			break
		}
	}
	return nil
}

func (r *Room) Players() []entity.Player {
	if r.playerItfs != nil {
		return r.playerItfs
	}
	for _, value := range r.players {
		r.playerItfs = append(r.playerItfs, value)
	}
	return r.playerItfs
}

// handle C2SMsg_LeaveRoom
func (r *Room) leavePlayer(playerId string) error {
	if idx := r.getPlayerIndex(playerId); idx == -1 {
		return errors.New("todo")
	} else {
		close(r.players[idx].respChan)
		r.players[idx] = nil
	}
	return nil
}

// handle C2SMsg_StartGame
func (r *Room) startGame() error {
	if r.game != nil {
		return errors.New("game exist")
	}
	if r.PlayerCount() < 2 {
		return errors.New("玩家未到最小值")
	}
	game, err := CreateGame(r.Type)
	if err != nil {
		return err
	}
	game.SetRoom(r)
	game.Init()
	var srcSide entity.Side
	var lastSide entity.Side
	var firstSide entity.Side
	for idx, value := range r.players {
		if value == nil {
			continue
		}
		side, err := CreateSide(r.Type)
		if err != nil {
			return err
		}
		if srcSide != nil {
			side.SetPrev(srcSide)
			srcSide.SetNext(side)
		}
		if firstSide == nil {
			firstSide = side
		}
		srcSide = side
		lastSide = side
		value.index = idx
		value.SetSide(side)
		side.Init()
	}
	lastSide.SetNext(firstSide)
	firstSide.SetPrev(lastSide)
	r.game = game
	r.Round = 1
	game.Reset(1)
	return nil
}

// 检查是房间消息还是具体游戏的消息
func (r *Room) preCheckMsg(msg *pb.C2SMsg) (isGame bool, err error) {
	if msg.Head != pb.C2SMsg_EnterRoom && !r.playerIdIsExist(msg.PlayerId) {
		return false, errors.New("不存在的玩家怎么把消息发过来了" + msg.PlayerId)
	}
	if msg.RoomId != r.Id {
		return false, errors.New("房间id不对怎么把消息发过来了" + msg.RoomId + "  <->  " + r.Id)
	}
	head := int32(msg.Head)
	if head < config.GateServerMsgEnd || head >= config.GameServerMsgEnd {
		return false, errors.New("无法处理的消息头" + strconv.Itoa(int(head)))
	}
	if head >= config.GameMsgStart {
		return true, nil
	}
	return false, nil
}

func (r *Room) playerIdIsExist(PlayerId string) bool {
	for _, p := range r.players {
		if p == nil {
			continue
		}
		if p.Id == PlayerId {
			return true
		}
	}
	return false
}
