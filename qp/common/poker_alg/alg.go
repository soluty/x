package poker_alg

import (
	"github.com/soluty/x/qp/entity"
	"math/rand"
)

func Shuffle(pokers []entity.Poker) []entity.Poker {
	idxes := rand.Perm(len(pokers))
	ret := make([]entity.Poker, len(pokers))
	for _, value := range idxes {
		ret = append(ret, pokers[value])
	}
	return ret
}

func IndexPoker(pokers []entity.Poker, p entity.Poker) int {
	for idx, value := range pokers {
		if p == value {
			return idx
		}
	}
	return -1
}