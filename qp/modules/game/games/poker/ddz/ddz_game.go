package ddz

import (
	"github.com/golang/protobuf/proto"
	"github.com/soluty/x/qp/common/poker_alg"
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/modules/game/service"
	"github.com/soluty/x/qp/pb"
)

// 斗地主3张底牌
type Game struct {
	// 有限debug
	debugPaidui []entity.Poker
	zhuangIndex int
	turn        int // 当前出牌的人
	players     []*Player
	paidui      []entity.Poker
	dipai       []entity.Poker
}

func NewGame() *Game {
	return &Game{
		zhuangIndex: -1,
	}
}

func (g *Game) Init() {
	var toShuffle = make([]entity.Poker, 48)
	for _, value := range g.debugPaidui {
		if idx := poker_alg.IndexPoker(paidui, value); idx < 0 {
			panic("xx")
		} else {
			toShuffle = append(paidui[:idx], paidui[idx+1:]...)
			g.paidui = append(g.paidui, value)
		}
	}
	g.debugPaidui = nil
	g.paidui = append(g.paidui, poker_alg.Shuffle(toShuffle)...)
}

func (g *Game) SetDebug(debugPaidui interface{}) {
	if debugPaidui == nil {
		return
	}
	g.debugPaidui = debugPaidui.([]entity.Poker)
}

// 给3个人发牌
func (g *Game) Fapai() {
	for i := 0; i < 3; i++ {
		_ = g.paidui[:17] // todo 给玩家发牌
		g.paidui = g.paidui[17:]
	}
	g.dipai = g.paidui
}

func (g *Game) getPlayer(playerId string) *service.Player {
	return nil
}

func (g *Game) Handle(playerIndex int, head pb.C2SMsg_Type, body proto.Message) error {
	//player := g.players[playerIndex]
	//switch head {
	//case pb.C2SMsg_Ddz_Jiaofen: // 叫庄
	//	msg := body.(*pb.Ddz_Jiaofen)
	//case pb.C2SMsg_Ddz_Chupai: // 出牌
	//	msg := body.(*pb.Ddz_Chupai)
	//case pb.C2SMsg_Ddz_Pass: // pass
	//}
	return nil
}
