package simple

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/soluty/x/qp/common/poker_alg"
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/modules/game/service"
	"github.com/soluty/x/qp/pb"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"math/rand"
	"time"
)

// 规则见rule.md
type Game struct {
	service.GameBase

	debugPaidui []entity.Poker
	zhuang      int
	turn        int // 当前要出牌的人的index
	players     []*Side
	paidui      []entity.Poker
	lastPoker   entity.Poker // 最新出的牌
	isOver      bool         // 游戏结束

	errWrapper
}

type errWrapper struct {
	err error
}

func (g *errWrapper) set(json, path string, value interface{}) (ret string) {
	if g.err != nil {
		return ""
	}
	ret, g.err = sjson.Set(json, path, value)
	return
}

func (g *Game) Unmarshal(res *gjson.Result) error {
	g.turn = int(res.Get("turn").Int())
	return nil
}

func (g *Game) Reset(idx int) {
	var toShuffle = make([]entity.Poker, 48)
	for _, value := range g.debugPaidui {
		if idx := poker_alg.IndexPoker(tablePoker, value); idx < 0 {
			panic("xx")
		} else {
			toShuffle = append(tablePoker[:idx], tablePoker[idx+1:]...)
			g.paidui = append(g.paidui, value)
		}
	}
	g.debugPaidui = nil
	g.paidui = append(g.paidui, poker_alg.Shuffle(toShuffle)...)
	g.paidui = tablePoker
	rand.Seed(time.Now().UnixNano())
	var zhuang = rand.Intn(g.Room().PlayerCount())
	for idx, value := range g.Room().Players() {
		if zhuang == 0 {
			g.zhuang = idx
			break
		}
		if value != nil {
			zhuang--
		}
	}
	g.Fapai()
}

func (g *Game) MarshalJSON() ([]byte, error) {
	ret := ""
	ret = g.set(ret, "turn", g.turn)
	ret = g.set(ret, "lastPoker", g.lastPoker)
	if g.err != nil {
		return nil, g.err
	}
	return []byte(ret), nil
}

func (g *Game) Init() {
	g.zhuang = -1
	g.turn = -1
}

func (g *Game) SetDebug(debugPaidui interface{}) {
	g.turn = 34
	if debugPaidui == nil {
		return
	}
	g.debugPaidui = debugPaidui.([]entity.Poker)
}

// 给3个人发牌
func (g *Game) Fapai() {
	g.turn = g.zhuang
	var side = g.GetSide(g.zhuang).(*Side)
	for s := side; ; s = s.Next().(*Side) {
		s.AddShoupai(g.paidui[:17])
		g.paidui = g.paidui[17:]
		if s.Next() == side {
			break
		}
	}
	for s := side; ; s = s.Next().(*Side) {
		g.sendShoupaiMsg(s.Player().Index())
		if s.Next() == side {
			break
		}
	}
	fmt.Println("aa")
}

// 从庄开始循环
func (g *Game) Range(cb func(side *Side)) {
	var side = g.GetSide(g.zhuang).(*Side)
	for s := side; ; s = s.Next().(*Side) {
		cb(s)
		if s.Next() == side {
			break
		}
	}
}

// 给某个人发送手牌消息
func (g *Game) sendShoupaiMsg(idx int) {
	msg := &pb.Game_Poker_Ddz_Shoupai{}
	g.Range(func(side *Side) {
		if side.Player().Index() == idx {
			msg.Shoupai = append(msg.Shoupai, side.msgShoupai(true))
		} else {
			msg.Shoupai = append(msg.Shoupai, side.msgShoupai(false))
		}
	})
	g.GetSide(idx).Player().Send(pb.S2CMsg_Login, msg)
}

func (g *Game) Handle(playerIndex int, head pb.C2SMsg_Type, body proto.Message) error {
	player := g.players[playerIndex]
	switch head {
	case pb.C2SMsg_Poker_Simple_Chupai: // 出牌
		msg, ok := body.(*pb.Game_Poker_Simple_Chupai)
		if !ok {
			return fmt.Errorf("handle Game_Poker_Simple_Chupai: 传入的pb不对，应该是 Game_Poker_Simple.Chupai")
		}
		if msg.Pai == nil {
			return fmt.Errorf("handle Game_Poker_Simple_Chupai: 没有传要出的牌")
		}
		return g.Chupai(player, msg.Pai)
	}
	return nil
}

// 玩家出一张牌，不能传nil
func (g *Game) Chupai(side *Side, msg *pb.Poker) error {
	if g.isOver {
		return fmt.Errorf("游戏已经结束")
	}
	if side.Player().Index() != g.turn {
		return fmt.Errorf("还没轮到%v出牌, turn=%v", side.Player().Index(), g.turn)
	}
	if err := canChu(side.shoupai, msg.Id, g.lastPoker); err != nil {
		return err
	}
	poker := entity.MustGetPoker(msg.Id)
	g.lastPoker = poker
	side.Chupai(poker)
	g.turn = (side.Player().Index() + 1) % 3
	nextPlayer := g.players[g.turn]
	if nextPlayer.RemainShoupai() == 0 {
		// 一家没手牌平局
		g.isOver = true
		g.CalScore(-1, int(poker.Number()))
	} else if !canYao(nextPlayer.shoupai, g.lastPoker) {
		// 下家要不起，上级赢
		g.isOver = true
		g.CalScore(side.Player().Index(), int(poker.Number()))
	}
	return nil
}

func (g *Game) CalScore(winner int, score int) {
	if winner == -1 {
		return
	}
	for idx, player := range g.players {
		if idx == winner {
			player.score = score
		} else {
			player.score = -score
		}
	}
}
