package simple

import (
	"errors"
	"fmt"
	"github.com/soluty/x/qp/common/poker_alg"
	"github.com/soluty/x/qp/entity"
)

// rule里面函数必须要进行单元测试
func canChu(shoupai []entity.Poker, pokerId uint32, shangshou entity.Poker) error {
	poker, err := entity.GetPoker(pokerId)
	if err != nil {
		return fmt.Errorf("canchu: %w", err)
	}
	if poker_alg.IndexPoker(shoupai, poker) < 0 {
		return errors.New("poker %v nod in pokers")
	}
	if shangshou == nil {
		return nil
	}
	if match(poker, shangshou) {
		return nil
	}
	return errors.New("poker not match")
}

// simple game 的核心规则，只有同花色或则数字+1才能要的起
func match(shoupai, chupai entity.Poker) bool {
	if shoupai.Color() == chupai.Color() || shoupai.Number() == chupai.Number()+1 {
		return true
	}
	return false
}

// 手牌中是不是要的起别人出的牌
func canYao(shoupai []entity.Poker, chu entity.Poker) bool {
	for _, value := range shoupai {
		if match(value, chu) {
			return true
		}
	}
	return false
}
