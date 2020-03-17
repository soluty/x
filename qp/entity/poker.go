package entity

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/soluty/x/qp/pb"
	"strconv"
)

type PokerColor byte
type PokerNumber byte

const (
	PokerColor_Heitao   PokerColor = 1
	PokerColor_Hongtao  PokerColor = 2
	PokerColor_Meihua   PokerColor = 3
	PokerColor_Fangpian PokerColor = 4
	PokerColor_Gui      PokerColor = 5
	PokerColor_Blank    PokerColor = 6
)

const (
	PokerNumber_A  PokerNumber = 1
	PokerNumber_2  PokerNumber = 2
	PokerNumber_3  PokerNumber = 3
	PokerNumber_4  PokerNumber = 4
	PokerNumber_5  PokerNumber = 5
	PokerNumber_6  PokerNumber = 6
	PokerNumber_7  PokerNumber = 7
	PokerNumber_8  PokerNumber = 8
	PokerNumber_9  PokerNumber = 9
	PokerNumber_10 PokerNumber = 10
	PokerNumber_J  PokerNumber = 11
	PokerNumber_Q  PokerNumber = 12
	PokerNumber_K  PokerNumber = 13

	PokerNumber_Xiaogui PokerNumber = 101
	PokerNumber_Dagui   PokerNumber = 102

	PokerNumber_Any PokerNumber = 255
)

// 享元模式
type poker struct {
	id     uint32
	group  byte // 第x副牌
	color  PokerColor
	number PokerNumber
}

func GetPoker(id uint32) (Poker, error) {
	if p, ok := pokers[id]; ok {
		return p, nil
	} else {
		return nil, errors.New("没有的扑克牌")
	}
}

func MustGetPoker(id uint32) Poker {
	if p, ok := pokers[id]; ok {
		return p
	} else {
		panic("MustGetPoker")
	}
}

func newPoker(g byte, c PokerColor, n PokerNumber) *poker {
	p := &poker{
		group:  g,
		color:  c,
		number: n,
	}
	p.id = binary.BigEndian.Uint32([]byte{0, p.group, byte(p.color), byte(p.number)})
	pokers[p.id] = p
	return p
}

func (p *poker) Proto() *pb.Poker {
	return &pb.Poker{
		Id:   p.id,
		Desc: p.String(),
	}
}

func (p *poker) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(p.id))), nil
}

func (p *poker) String() string {
	lang := GetLang()
	var ret string
	switch lang {
	case Zh_cn:
		ret = p.chinaStr()
	case En_us:
		ret = p.englishStr()
	default:
		ret = p.chinaStr()
	}
	return ret
}

func (p *poker) chinaStr() string {
	dis := p.Display()
	ret := fmt.Sprintf("%d_%v", p.group, dis)
	return ret
}

func (p *poker) Display() string {
	color := ""
	switch p.color {
	case PokerColor_Heitao:
		color = "♠"
	case PokerColor_Hongtao:
		color = "♥"
	case PokerColor_Meihua:
		color = "♣"
	case PokerColor_Fangpian:
		color = "♦"
	}
	num := ""
	switch p.number {
	case PokerNumber_Xiaogui:
		num = "小王"
	case PokerNumber_Dagui:
		num = "大王"
	case PokerNumber_Any:
		num = "花牌"
	case PokerNumber_A:
		num = "A"
	case PokerNumber_J:
		num = "J"
	case PokerNumber_Q:
		num = "Q"
	case PokerNumber_K:
		num = "K"
	case PokerNumber_2, PokerNumber_3, PokerNumber_4, PokerNumber_5, PokerNumber_6, PokerNumber_7, PokerNumber_8, PokerNumber_9:
		num = fmt.Sprintf("%d", p.number)
	}
	return fmt.Sprintf("%v%v", color, num)
}

func (p *poker) englishStr() string {
	return p.chinaStr()
}

type Poker interface {
	fmt.Stringer
	json.Marshaler
	Id() uint32
	Group() byte
	Color() PokerColor
	Number() PokerNumber
	Proto() *pb.Poker
	Display() string
}

var pokers = map[uint32]*poker{}

func (p *poker) Id() uint32 {
	return p.id
}
func (p *poker) Group() byte {
	return p.group
}
func (p *poker) Color() PokerColor {
	return p.color
}
func (p *poker) Number() PokerNumber {
	return p.number
}

var (
	G1_Xiaogui = newPoker(1, PokerColor_Gui, PokerNumber_Xiaogui)
	G1_Dagui   = newPoker(1, PokerColor_Gui, PokerNumber_Dagui)
	G1_Blank   = newPoker(1, PokerColor_Blank, PokerNumber_Any)
	G2_Xiaogui = newPoker(2, PokerColor_Gui, PokerNumber_Xiaogui)
	G2_Dagui   = newPoker(2, PokerColor_Gui, PokerNumber_Dagui)
	G2_Blank   = newPoker(2, PokerColor_Blank, PokerNumber_Any)
	G3_Xiaogui = newPoker(3, PokerColor_Gui, PokerNumber_Xiaogui)
	G3_Dagui   = newPoker(3, PokerColor_Gui, PokerNumber_Dagui)
	G3_Blank   = newPoker(3, PokerColor_Blank, PokerNumber_Any)
	G4_Xiaogui = newPoker(4, PokerColor_Gui, PokerNumber_Xiaogui)
	G4_Dagui   = newPoker(4, PokerColor_Gui, PokerNumber_Dagui)
	G4_Blank   = newPoker(4, PokerColor_Blank, PokerNumber_Any)

	G1_Heitao_A    = newPoker(1, PokerColor_Heitao, PokerNumber_A)
	G1_Heitao_2    = newPoker(1, PokerColor_Heitao, PokerNumber_2)
	G1_Heitao_3    = newPoker(1, PokerColor_Heitao, PokerNumber_3)
	G1_Heitao_4    = newPoker(1, PokerColor_Heitao, PokerNumber_4)
	G1_Heitao_5    = newPoker(1, PokerColor_Heitao, PokerNumber_5)
	G1_Heitao_6    = newPoker(1, PokerColor_Heitao, PokerNumber_6)
	G1_Heitao_7    = newPoker(1, PokerColor_Heitao, PokerNumber_7)
	G1_Heitao_8    = newPoker(1, PokerColor_Heitao, PokerNumber_8)
	G1_Heitao_9    = newPoker(1, PokerColor_Heitao, PokerNumber_9)
	G1_Heitao_10   = newPoker(1, PokerColor_Heitao, PokerNumber_10)
	G1_Heitao_J    = newPoker(1, PokerColor_Heitao, PokerNumber_J)
	G1_Heitao_Q    = newPoker(1, PokerColor_Heitao, PokerNumber_Q)
	G1_Heitao_K    = newPoker(1, PokerColor_Heitao, PokerNumber_K)
	G1_Hongtao_A   = newPoker(1, PokerColor_Hongtao, PokerNumber_A)
	G1_Hongtao_2   = newPoker(1, PokerColor_Hongtao, PokerNumber_2)
	G1_Hongtao_3   = newPoker(1, PokerColor_Hongtao, PokerNumber_3)
	G1_Hongtao_4   = newPoker(1, PokerColor_Hongtao, PokerNumber_4)
	G1_Hongtao_5   = newPoker(1, PokerColor_Hongtao, PokerNumber_5)
	G1_Hongtao_6   = newPoker(1, PokerColor_Hongtao, PokerNumber_6)
	G1_Hongtao_7   = newPoker(1, PokerColor_Hongtao, PokerNumber_7)
	G1_Hongtao_8   = newPoker(1, PokerColor_Hongtao, PokerNumber_8)
	G1_Hongtao_9   = newPoker(1, PokerColor_Hongtao, PokerNumber_9)
	G1_Hongtao_10  = newPoker(1, PokerColor_Hongtao, PokerNumber_10)
	G1_Hongtao_J   = newPoker(1, PokerColor_Hongtao, PokerNumber_J)
	G1_Hongtao_Q   = newPoker(1, PokerColor_Hongtao, PokerNumber_Q)
	G1_Hongtao_K   = newPoker(1, PokerColor_Hongtao, PokerNumber_K)
	G1_Meihua_A    = newPoker(1, PokerColor_Meihua, PokerNumber_A)
	G1_Meihua_2    = newPoker(1, PokerColor_Meihua, PokerNumber_2)
	G1_Meihua_3    = newPoker(1, PokerColor_Meihua, PokerNumber_3)
	G1_Meihua_4    = newPoker(1, PokerColor_Meihua, PokerNumber_4)
	G1_Meihua_5    = newPoker(1, PokerColor_Meihua, PokerNumber_5)
	G1_Meihua_6    = newPoker(1, PokerColor_Meihua, PokerNumber_6)
	G1_Meihua_7    = newPoker(1, PokerColor_Meihua, PokerNumber_7)
	G1_Meihua_8    = newPoker(1, PokerColor_Meihua, PokerNumber_8)
	G1_Meihua_9    = newPoker(1, PokerColor_Meihua, PokerNumber_9)
	G1_Meihua_10   = newPoker(1, PokerColor_Meihua, PokerNumber_10)
	G1_Meihua_J    = newPoker(1, PokerColor_Meihua, PokerNumber_J)
	G1_Meihua_Q    = newPoker(1, PokerColor_Meihua, PokerNumber_Q)
	G1_Meihua_K    = newPoker(1, PokerColor_Meihua, PokerNumber_K)
	G1_Fangpian_A  = newPoker(1, PokerColor_Fangpian, PokerNumber_A)
	G1_Fangpian_2  = newPoker(1, PokerColor_Fangpian, PokerNumber_2)
	G1_Fangpian_3  = newPoker(1, PokerColor_Fangpian, PokerNumber_3)
	G1_Fangpian_4  = newPoker(1, PokerColor_Fangpian, PokerNumber_4)
	G1_Fangpian_5  = newPoker(1, PokerColor_Fangpian, PokerNumber_5)
	G1_Fangpian_6  = newPoker(1, PokerColor_Fangpian, PokerNumber_6)
	G1_Fangpian_7  = newPoker(1, PokerColor_Fangpian, PokerNumber_7)
	G1_Fangpian_8  = newPoker(1, PokerColor_Fangpian, PokerNumber_8)
	G1_Fangpian_9  = newPoker(1, PokerColor_Fangpian, PokerNumber_9)
	G1_Fangpian_10 = newPoker(1, PokerColor_Fangpian, PokerNumber_10)
	G1_Fangpian_J  = newPoker(1, PokerColor_Fangpian, PokerNumber_J)
	G1_Fangpian_Q  = newPoker(1, PokerColor_Fangpian, PokerNumber_Q)
	G1_Fangpian_K  = newPoker(1, PokerColor_Fangpian, PokerNumber_K)

	G2_Heitao_A    = newPoker(2, PokerColor_Heitao, PokerNumber_A)
	G2_Heitao_2    = newPoker(2, PokerColor_Heitao, PokerNumber_2)
	G2_Heitao_3    = newPoker(2, PokerColor_Heitao, PokerNumber_3)
	G2_Heitao_4    = newPoker(2, PokerColor_Heitao, PokerNumber_4)
	G2_Heitao_5    = newPoker(2, PokerColor_Heitao, PokerNumber_5)
	G2_Heitao_6    = newPoker(2, PokerColor_Heitao, PokerNumber_6)
	G2_Heitao_7    = newPoker(2, PokerColor_Heitao, PokerNumber_7)
	G2_Heitao_8    = newPoker(2, PokerColor_Heitao, PokerNumber_8)
	G2_Heitao_9    = newPoker(2, PokerColor_Heitao, PokerNumber_9)
	G2_Heitao_10   = newPoker(2, PokerColor_Heitao, PokerNumber_10)
	G2_Heitao_J    = newPoker(2, PokerColor_Heitao, PokerNumber_J)
	G2_Heitao_Q    = newPoker(2, PokerColor_Heitao, PokerNumber_Q)
	G2_Heitao_K    = newPoker(2, PokerColor_Heitao, PokerNumber_K)
	G2_Hongtao_A   = newPoker(2, PokerColor_Hongtao, PokerNumber_A)
	G2_Hongtao_2   = newPoker(2, PokerColor_Hongtao, PokerNumber_2)
	G2_Hongtao_3   = newPoker(2, PokerColor_Hongtao, PokerNumber_3)
	G2_Hongtao_4   = newPoker(2, PokerColor_Hongtao, PokerNumber_4)
	G2_Hongtao_5   = newPoker(2, PokerColor_Hongtao, PokerNumber_5)
	G2_Hongtao_6   = newPoker(2, PokerColor_Hongtao, PokerNumber_6)
	G2_Hongtao_7   = newPoker(2, PokerColor_Hongtao, PokerNumber_7)
	G2_Hongtao_8   = newPoker(2, PokerColor_Hongtao, PokerNumber_8)
	G2_Hongtao_9   = newPoker(2, PokerColor_Hongtao, PokerNumber_9)
	G2_Hongtao_10  = newPoker(2, PokerColor_Hongtao, PokerNumber_10)
	G2_Hongtao_J   = newPoker(2, PokerColor_Hongtao, PokerNumber_J)
	G2_Hongtao_Q   = newPoker(2, PokerColor_Hongtao, PokerNumber_Q)
	G2_Hongtao_K   = newPoker(2, PokerColor_Hongtao, PokerNumber_K)
	G2_Meihua_A    = newPoker(2, PokerColor_Meihua, PokerNumber_A)
	G2_Meihua_2    = newPoker(2, PokerColor_Meihua, PokerNumber_2)
	G2_Meihua_3    = newPoker(2, PokerColor_Meihua, PokerNumber_3)
	G2_Meihua_4    = newPoker(2, PokerColor_Meihua, PokerNumber_4)
	G2_Meihua_5    = newPoker(2, PokerColor_Meihua, PokerNumber_5)
	G2_Meihua_6    = newPoker(2, PokerColor_Meihua, PokerNumber_6)
	G2_Meihua_7    = newPoker(2, PokerColor_Meihua, PokerNumber_7)
	G2_Meihua_8    = newPoker(2, PokerColor_Meihua, PokerNumber_8)
	G2_Meihua_9    = newPoker(2, PokerColor_Meihua, PokerNumber_9)
	G2_Meihua_10   = newPoker(2, PokerColor_Meihua, PokerNumber_10)
	G2_Meihua_J    = newPoker(2, PokerColor_Meihua, PokerNumber_J)
	G2_Meihua_Q    = newPoker(2, PokerColor_Meihua, PokerNumber_Q)
	G2_Meihua_K    = newPoker(2, PokerColor_Meihua, PokerNumber_K)
	G2_Fangpian_A  = newPoker(2, PokerColor_Fangpian, PokerNumber_A)
	G2_Fangpian_2  = newPoker(2, PokerColor_Fangpian, PokerNumber_2)
	G2_Fangpian_3  = newPoker(2, PokerColor_Fangpian, PokerNumber_3)
	G2_Fangpian_4  = newPoker(2, PokerColor_Fangpian, PokerNumber_4)
	G2_Fangpian_5  = newPoker(2, PokerColor_Fangpian, PokerNumber_5)
	G2_Fangpian_6  = newPoker(2, PokerColor_Fangpian, PokerNumber_6)
	G2_Fangpian_7  = newPoker(2, PokerColor_Fangpian, PokerNumber_7)
	G2_Fangpian_8  = newPoker(2, PokerColor_Fangpian, PokerNumber_8)
	G2_Fangpian_9  = newPoker(2, PokerColor_Fangpian, PokerNumber_9)
	G2_Fangpian_10 = newPoker(2, PokerColor_Fangpian, PokerNumber_10)
	G2_Fangpian_J  = newPoker(2, PokerColor_Fangpian, PokerNumber_J)
	G2_Fangpian_Q  = newPoker(2, PokerColor_Fangpian, PokerNumber_Q)
	G2_Fangpian_K  = newPoker(2, PokerColor_Fangpian, PokerNumber_K)

	G3_Heitao_A    = newPoker(3, PokerColor_Heitao, PokerNumber_A)
	G3_Heitao_2    = newPoker(3, PokerColor_Heitao, PokerNumber_2)
	G3_Heitao_3    = newPoker(3, PokerColor_Heitao, PokerNumber_3)
	G3_Heitao_4    = newPoker(3, PokerColor_Heitao, PokerNumber_4)
	G3_Heitao_5    = newPoker(3, PokerColor_Heitao, PokerNumber_5)
	G3_Heitao_6    = newPoker(3, PokerColor_Heitao, PokerNumber_6)
	G3_Heitao_7    = newPoker(3, PokerColor_Heitao, PokerNumber_7)
	G3_Heitao_8    = newPoker(3, PokerColor_Heitao, PokerNumber_8)
	G3_Heitao_9    = newPoker(3, PokerColor_Heitao, PokerNumber_9)
	G3_Heitao_10   = newPoker(3, PokerColor_Heitao, PokerNumber_10)
	G3_Heitao_J    = newPoker(3, PokerColor_Heitao, PokerNumber_J)
	G3_Heitao_Q    = newPoker(3, PokerColor_Heitao, PokerNumber_Q)
	G3_Heitao_K    = newPoker(3, PokerColor_Heitao, PokerNumber_K)
	G3_Hongtao_A   = newPoker(3, PokerColor_Hongtao, PokerNumber_A)
	G3_Hongtao_2   = newPoker(3, PokerColor_Hongtao, PokerNumber_2)
	G3_Hongtao_3   = newPoker(3, PokerColor_Hongtao, PokerNumber_3)
	G3_Hongtao_4   = newPoker(3, PokerColor_Hongtao, PokerNumber_4)
	G3_Hongtao_5   = newPoker(3, PokerColor_Hongtao, PokerNumber_5)
	G3_Hongtao_6   = newPoker(3, PokerColor_Hongtao, PokerNumber_6)
	G3_Hongtao_7   = newPoker(3, PokerColor_Hongtao, PokerNumber_7)
	G3_Hongtao_8   = newPoker(3, PokerColor_Hongtao, PokerNumber_8)
	G3_Hongtao_9   = newPoker(3, PokerColor_Hongtao, PokerNumber_9)
	G3_Hongtao_10  = newPoker(3, PokerColor_Hongtao, PokerNumber_10)
	G3_Hongtao_J   = newPoker(3, PokerColor_Hongtao, PokerNumber_J)
	G3_Hongtao_Q   = newPoker(3, PokerColor_Hongtao, PokerNumber_Q)
	G3_Hongtao_K   = newPoker(3, PokerColor_Hongtao, PokerNumber_K)
	G3_Meihua_A    = newPoker(3, PokerColor_Meihua, PokerNumber_A)
	G3_Meihua_2    = newPoker(3, PokerColor_Meihua, PokerNumber_2)
	G3_Meihua_3    = newPoker(3, PokerColor_Meihua, PokerNumber_3)
	G3_Meihua_4    = newPoker(3, PokerColor_Meihua, PokerNumber_4)
	G3_Meihua_5    = newPoker(3, PokerColor_Meihua, PokerNumber_5)
	G3_Meihua_6    = newPoker(3, PokerColor_Meihua, PokerNumber_6)
	G3_Meihua_7    = newPoker(3, PokerColor_Meihua, PokerNumber_7)
	G3_Meihua_8    = newPoker(3, PokerColor_Meihua, PokerNumber_8)
	G3_Meihua_9    = newPoker(3, PokerColor_Meihua, PokerNumber_9)
	G3_Meihua_10   = newPoker(3, PokerColor_Meihua, PokerNumber_10)
	G3_Meihua_J    = newPoker(3, PokerColor_Meihua, PokerNumber_J)
	G3_Meihua_Q    = newPoker(3, PokerColor_Meihua, PokerNumber_Q)
	G3_Meihua_K    = newPoker(3, PokerColor_Meihua, PokerNumber_K)
	G3_Fangpian_A  = newPoker(3, PokerColor_Fangpian, PokerNumber_A)
	G3_Fangpian_2  = newPoker(3, PokerColor_Fangpian, PokerNumber_2)
	G3_Fangpian_3  = newPoker(3, PokerColor_Fangpian, PokerNumber_3)
	G3_Fangpian_4  = newPoker(3, PokerColor_Fangpian, PokerNumber_4)
	G3_Fangpian_5  = newPoker(3, PokerColor_Fangpian, PokerNumber_5)
	G3_Fangpian_6  = newPoker(3, PokerColor_Fangpian, PokerNumber_6)
	G3_Fangpian_7  = newPoker(3, PokerColor_Fangpian, PokerNumber_7)
	G3_Fangpian_8  = newPoker(3, PokerColor_Fangpian, PokerNumber_8)
	G3_Fangpian_9  = newPoker(3, PokerColor_Fangpian, PokerNumber_9)
	G3_Fangpian_10 = newPoker(3, PokerColor_Fangpian, PokerNumber_10)
	G3_Fangpian_J  = newPoker(3, PokerColor_Fangpian, PokerNumber_J)
	G3_Fangpian_Q  = newPoker(3, PokerColor_Fangpian, PokerNumber_Q)
	G3_Fangpian_K  = newPoker(3, PokerColor_Fangpian, PokerNumber_K)

	G4_Heitao_A    = newPoker(4, PokerColor_Heitao, PokerNumber_A)
	G4_Heitao_2    = newPoker(4, PokerColor_Heitao, PokerNumber_2)
	G4_Heitao_3    = newPoker(4, PokerColor_Heitao, PokerNumber_3)
	G4_Heitao_4    = newPoker(4, PokerColor_Heitao, PokerNumber_4)
	G4_Heitao_5    = newPoker(4, PokerColor_Heitao, PokerNumber_5)
	G4_Heitao_6    = newPoker(4, PokerColor_Heitao, PokerNumber_6)
	G4_Heitao_7    = newPoker(4, PokerColor_Heitao, PokerNumber_7)
	G4_Heitao_8    = newPoker(4, PokerColor_Heitao, PokerNumber_8)
	G4_Heitao_9    = newPoker(4, PokerColor_Heitao, PokerNumber_9)
	G4_Heitao_10   = newPoker(4, PokerColor_Heitao, PokerNumber_10)
	G4_Heitao_J    = newPoker(4, PokerColor_Heitao, PokerNumber_J)
	G4_Heitao_Q    = newPoker(4, PokerColor_Heitao, PokerNumber_Q)
	G4_Heitao_K    = newPoker(4, PokerColor_Heitao, PokerNumber_K)
	G4_Hongtao_A   = newPoker(4, PokerColor_Hongtao, PokerNumber_A)
	G4_Hongtao_2   = newPoker(4, PokerColor_Hongtao, PokerNumber_2)
	G4_Hongtao_3   = newPoker(4, PokerColor_Hongtao, PokerNumber_3)
	G4_Hongtao_4   = newPoker(4, PokerColor_Hongtao, PokerNumber_4)
	G4_Hongtao_5   = newPoker(4, PokerColor_Hongtao, PokerNumber_5)
	G4_Hongtao_6   = newPoker(4, PokerColor_Hongtao, PokerNumber_6)
	G4_Hongtao_7   = newPoker(4, PokerColor_Hongtao, PokerNumber_7)
	G4_Hongtao_8   = newPoker(4, PokerColor_Hongtao, PokerNumber_8)
	G4_Hongtao_9   = newPoker(4, PokerColor_Hongtao, PokerNumber_9)
	G4_Hongtao_10  = newPoker(4, PokerColor_Hongtao, PokerNumber_10)
	G4_Hongtao_J   = newPoker(4, PokerColor_Hongtao, PokerNumber_J)
	G4_Hongtao_Q   = newPoker(4, PokerColor_Hongtao, PokerNumber_Q)
	G4_Hongtao_K   = newPoker(4, PokerColor_Hongtao, PokerNumber_K)
	G4_Meihua_A    = newPoker(4, PokerColor_Meihua, PokerNumber_A)
	G4_Meihua_2    = newPoker(4, PokerColor_Meihua, PokerNumber_2)
	G4_Meihua_3    = newPoker(4, PokerColor_Meihua, PokerNumber_3)
	G4_Meihua_4    = newPoker(4, PokerColor_Meihua, PokerNumber_4)
	G4_Meihua_5    = newPoker(4, PokerColor_Meihua, PokerNumber_5)
	G4_Meihua_6    = newPoker(4, PokerColor_Meihua, PokerNumber_6)
	G4_Meihua_7    = newPoker(4, PokerColor_Meihua, PokerNumber_7)
	G4_Meihua_8    = newPoker(4, PokerColor_Meihua, PokerNumber_8)
	G4_Meihua_9    = newPoker(4, PokerColor_Meihua, PokerNumber_9)
	G4_Meihua_10   = newPoker(4, PokerColor_Meihua, PokerNumber_10)
	G4_Meihua_J    = newPoker(4, PokerColor_Meihua, PokerNumber_J)
	G4_Meihua_Q    = newPoker(4, PokerColor_Meihua, PokerNumber_Q)
	G4_Meihua_K    = newPoker(4, PokerColor_Meihua, PokerNumber_K)
	G4_Fangpian_A  = newPoker(4, PokerColor_Fangpian, PokerNumber_A)
	G4_Fangpian_2  = newPoker(4, PokerColor_Fangpian, PokerNumber_2)
	G4_Fangpian_3  = newPoker(4, PokerColor_Fangpian, PokerNumber_3)
	G4_Fangpian_4  = newPoker(4, PokerColor_Fangpian, PokerNumber_4)
	G4_Fangpian_5  = newPoker(4, PokerColor_Fangpian, PokerNumber_5)
	G4_Fangpian_6  = newPoker(4, PokerColor_Fangpian, PokerNumber_6)
	G4_Fangpian_7  = newPoker(4, PokerColor_Fangpian, PokerNumber_7)
	G4_Fangpian_8  = newPoker(4, PokerColor_Fangpian, PokerNumber_8)
	G4_Fangpian_9  = newPoker(4, PokerColor_Fangpian, PokerNumber_9)
	G4_Fangpian_10 = newPoker(4, PokerColor_Fangpian, PokerNumber_10)
	G4_Fangpian_J  = newPoker(4, PokerColor_Fangpian, PokerNumber_J)
	G4_Fangpian_Q  = newPoker(4, PokerColor_Fangpian, PokerNumber_Q)
	G4_Fangpian_K  = newPoker(4, PokerColor_Fangpian, PokerNumber_K)
)
