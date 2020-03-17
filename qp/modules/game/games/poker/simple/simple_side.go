package simple

import (
	"github.com/soluty/x/qp/entity"
	"github.com/soluty/x/qp/modules/game/service"
	"github.com/soluty/x/qp/pb"
)

type Side struct {
	service.SideBase

	shoupai []entity.Poker
	lastChu entity.Poker // 上轮出的牌
	score   int          // 本局分数
}

func (p *Side) Chupai(poker entity.Poker) {
	for idx, value := range p.shoupai {
		if value == poker {
			p.shoupai = append(p.shoupai[:idx], p.shoupai[idx+1:]...)
			p.lastChu = poker
			break
		}
	}
}

func (p *Side) RemainShoupai() int {
	return len(p.shoupai)
}

func (p *Side) AddShoupai(pokers []entity.Poker) {
	p.shoupai = append(p.shoupai, pokers...)
}

func (p *Side) msgShoupai(isSelf bool) *pb.Data_Shoupai {
	msg := &pb.Data_Shoupai{}
	for _, value := range p.shoupai {
		if isSelf {
			msg.Pai = append(msg.Pai, value.Proto())
		} else {
			msg.Pai = append(msg.Pai, &pb.Poker{})
		}
	}
	return msg
}
