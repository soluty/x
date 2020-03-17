package entity

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/soluty/x/qp/pb"
)

type MahjongColor byte
type MahjongNumber byte

const (
	MahjongColor_Tiao MahjongColor = 1
	MahjongColor_Tong MahjongColor = 2
	MahjongColor_Wan  MahjongColor = 3

	MahjongColor_Feng MahjongColor = 11
	MahjongColor_Hua  MahjongColor = 21
)

const (
	MahjongNumber_1 MahjongNumber = 1
	MahjongNumber_2 MahjongNumber = 2
	MahjongNumber_3 MahjongNumber = 3
	MahjongNumber_4 MahjongNumber = 4
	MahjongNumber_5 MahjongNumber = 5
	MahjongNumber_6 MahjongNumber = 6
	MahjongNumber_7 MahjongNumber = 7
	MahjongNumber_8 MahjongNumber = 8
	MahjongNumber_9 MahjongNumber = 9

	MahjongNumber_Dong          MahjongNumber = 101
	MahjongNumber_Nan           MahjongNumber = 102
	MahjongNumber_Xi            MahjongNumber = 103
	MahjongNumber_Bei           MahjongNumber = 104
	MahjongNumber_Zhong         MahjongNumber = 105
	MahjongNumber_Fa            MahjongNumber = 106
	MahjongNumber_Bai           MahjongNumber = 107
	MahjongNumber_Plum          MahjongNumber = 201
	MahjongNumber_Orchid        MahjongNumber = 202
	MahjongNumber_Chrysanthemum MahjongNumber = 203
	MahjongNumber_Bamboo        MahjongNumber = 204
	MahjongNumber_Spring        MahjongNumber = 205
	MahjongNumber_Summer        MahjongNumber = 206
	MahjongNumber_Autumn        MahjongNumber = 207
	MahjongNumber_Winter        MahjongNumber = 208
)

// 享元模式
type mahjong struct {
	id     uint32
	group  byte // 第x副牌
	color  MahjongColor
	number MahjongNumber
}

func newMahjong(g byte, c MahjongColor, n MahjongNumber) *mahjong {
	p := &mahjong{
		group:  g,
		color:  c,
		number: n,
	}
	p.id = binary.BigEndian.Uint32([]byte{0, p.group, byte(p.color), byte(p.number)})
	return p
}


func (p *mahjong) String() string {
	color := ""
	switch p.color {
	case MahjongColor_Tiao:
		color = "条"
	case MahjongColor_Tong:
		color = "筒"
	case MahjongColor_Wan:
		color = "万"
	}
	num := ""
	switch p.number {
	case MahjongNumber_Dong:
		num = "东"
	case MahjongNumber_Nan:
		num = "南"
	case MahjongNumber_Xi:
		num = "西"
	case MahjongNumber_Bei:
		num = "北"
	case MahjongNumber_Zhong:
		num = "中"
	case MahjongNumber_Fa:
		num = "發"
	case MahjongNumber_Bai:
		num = "白"
	case MahjongNumber_Plum:
		num = "梅"
	case MahjongNumber_Orchid:
		num = "兰"
	case MahjongNumber_Chrysanthemum:
		num = "菊"
	case MahjongNumber_Bamboo:
		num = "竹"
	case MahjongNumber_Spring:
		num = "春"
	case MahjongNumber_Summer:
		num = "夏"
	case MahjongNumber_Autumn:
		num = "秋"
	case MahjongNumber_Winter:
		num = "东"
	case MahjongNumber_1:
		num = "一"
	case MahjongNumber_2:
		num = "二"
	case MahjongNumber_3:
		num = "三"
	case MahjongNumber_4:
		num = "四"
	case MahjongNumber_5:
		num = "五"
	case MahjongNumber_6:
		num = "六"
	case MahjongNumber_7:
		num = "七"
	case MahjongNumber_8:
		num = "八"
	case MahjongNumber_9:
		num = "九"
	}
	return fmt.Sprintf("%d_%v%v", p.group, num, color)
}

type Mahjong interface {
	Id() uint32
	Group() byte
	Color() MahjongColor
	Number() MahjongNumber
	From(id uint32) (Mahjong, error)
	Proto() *pb.Mahjong
}

var mahjongs = map[uint32]*mahjong{}

func (p *mahjong) From(id uint32) (Mahjong, error) {
	if p, ok := mahjongs[id]; ok {
		return p, nil
	} else {
		return nil, errors.New("没有的麻将牌")
	}
}

func (p *mahjong) Proto() *pb.Mahjong {
	return &pb.Mahjong{
		Id: p.id,
	}
}

func (p *mahjong) Id() uint32 {
	return p.id
}

func (p *mahjong) Group() byte {
	return p.group
}

func (p *mahjong) Color() MahjongColor {
	return p.color
}

func (p *mahjong) Number() MahjongNumber {
	return p.number
}

var (
	G1_Wan_1         = newMahjong(1, MahjongColor_Wan, MahjongNumber_1)
	G1_Wan_2         = newMahjong(1, MahjongColor_Wan, MahjongNumber_2)
	G1_Wan_3         = newMahjong(1, MahjongColor_Wan, MahjongNumber_3)
	G1_Wan_4         = newMahjong(1, MahjongColor_Wan, MahjongNumber_4)
	G1_Wan_5         = newMahjong(1, MahjongColor_Wan, MahjongNumber_5)
	G1_Wan_6         = newMahjong(1, MahjongColor_Wan, MahjongNumber_6)
	G1_Wan_7         = newMahjong(1, MahjongColor_Wan, MahjongNumber_7)
	G1_Wan_8         = newMahjong(1, MahjongColor_Wan, MahjongNumber_8)
	G1_Wan_9         = newMahjong(1, MahjongColor_Wan, MahjongNumber_9)
	G1_Tong_1        = newMahjong(1, MahjongColor_Tong, MahjongNumber_1)
	G1_Tong_2        = newMahjong(1, MahjongColor_Tong, MahjongNumber_2)
	G1_Tong_3        = newMahjong(1, MahjongColor_Tong, MahjongNumber_3)
	G1_Tong_4        = newMahjong(1, MahjongColor_Tong, MahjongNumber_4)
	G1_Tong_5        = newMahjong(1, MahjongColor_Tong, MahjongNumber_5)
	G1_Tong_6        = newMahjong(1, MahjongColor_Tong, MahjongNumber_6)
	G1_Tong_7        = newMahjong(1, MahjongColor_Tong, MahjongNumber_7)
	G1_Tong_8        = newMahjong(1, MahjongColor_Tong, MahjongNumber_8)
	G1_Tong_9        = newMahjong(1, MahjongColor_Tong, MahjongNumber_9)
	G1_Tiao_1        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_1)
	G1_Tiao_2        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_2)
	G1_Tiao_3        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_3)
	G1_Tiao_4        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_4)
	G1_Tiao_5        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_5)
	G1_Tiao_6        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_6)
	G1_Tiao_7        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_7)
	G1_Tiao_8        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_8)
	G1_Tiao_9        = newMahjong(1, MahjongColor_Tiao, MahjongNumber_9)
	G1_Dong          = newMahjong(1, MahjongColor_Feng, MahjongNumber_Dong)
	G1_Nan           = newMahjong(1, MahjongColor_Feng, MahjongNumber_Nan)
	G1_Xi            = newMahjong(1, MahjongColor_Feng, MahjongNumber_Xi)
	G1_Bei           = newMahjong(1, MahjongColor_Feng, MahjongNumber_Bei)
	G1_Zhong         = newMahjong(1, MahjongColor_Feng, MahjongNumber_Zhong)
	G1_Fa            = newMahjong(1, MahjongColor_Feng, MahjongNumber_Fa)
	G1_Bai           = newMahjong(1, MahjongColor_Feng, MahjongNumber_Bai)
	G1_Plum          = newMahjong(1, MahjongColor_Hua, MahjongNumber_Plum)
	G1_Orchid        = newMahjong(1, MahjongColor_Hua, MahjongNumber_Orchid)
	G1_Chrysanthemum = newMahjong(1, MahjongColor_Hua, MahjongNumber_Chrysanthemum)
	G1_Bamboo        = newMahjong(1, MahjongColor_Hua, MahjongNumber_Bamboo)
	G1_Spring        = newMahjong(1, MahjongColor_Hua, MahjongNumber_Spring)
	G1_Summer        = newMahjong(1, MahjongColor_Hua, MahjongNumber_Summer)
	G1_Autumn        = newMahjong(1, MahjongColor_Hua, MahjongNumber_Autumn)
	G1_Winter        = newMahjong(1, MahjongColor_Hua, MahjongNumber_Winter)

	G2_Wan_1         = newMahjong(2, MahjongColor_Wan, MahjongNumber_1)
	G2_Wan_2         = newMahjong(2, MahjongColor_Wan, MahjongNumber_2)
	G2_Wan_3         = newMahjong(2, MahjongColor_Wan, MahjongNumber_3)
	G2_Wan_4         = newMahjong(2, MahjongColor_Wan, MahjongNumber_4)
	G2_Wan_5         = newMahjong(2, MahjongColor_Wan, MahjongNumber_5)
	G2_Wan_6         = newMahjong(2, MahjongColor_Wan, MahjongNumber_6)
	G2_Wan_7         = newMahjong(2, MahjongColor_Wan, MahjongNumber_7)
	G2_Wan_8         = newMahjong(2, MahjongColor_Wan, MahjongNumber_8)
	G2_Wan_9         = newMahjong(2, MahjongColor_Wan, MahjongNumber_9)
	G2_Tong_1        = newMahjong(2, MahjongColor_Tong, MahjongNumber_1)
	G2_Tong_2        = newMahjong(2, MahjongColor_Tong, MahjongNumber_2)
	G2_Tong_3        = newMahjong(2, MahjongColor_Tong, MahjongNumber_3)
	G2_Tong_4        = newMahjong(2, MahjongColor_Tong, MahjongNumber_4)
	G2_Tong_5        = newMahjong(2, MahjongColor_Tong, MahjongNumber_5)
	G2_Tong_6        = newMahjong(2, MahjongColor_Tong, MahjongNumber_6)
	G2_Tong_7        = newMahjong(2, MahjongColor_Tong, MahjongNumber_7)
	G2_Tong_8        = newMahjong(2, MahjongColor_Tong, MahjongNumber_8)
	G2_Tong_9        = newMahjong(2, MahjongColor_Tong, MahjongNumber_9)
	G2_Tiao_1        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_1)
	G2_Tiao_2        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_2)
	G2_Tiao_3        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_3)
	G2_Tiao_4        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_4)
	G2_Tiao_5        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_5)
	G2_Tiao_6        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_6)
	G2_Tiao_7        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_7)
	G2_Tiao_8        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_8)
	G2_Tiao_9        = newMahjong(2, MahjongColor_Tiao, MahjongNumber_9)
	G2_Dong          = newMahjong(2, MahjongColor_Feng, MahjongNumber_Dong)
	G2_Nan           = newMahjong(2, MahjongColor_Feng, MahjongNumber_Nan)
	G2_Xi            = newMahjong(2, MahjongColor_Feng, MahjongNumber_Xi)
	G2_Bei           = newMahjong(2, MahjongColor_Feng, MahjongNumber_Bei)
	G2_Zhong         = newMahjong(2, MahjongColor_Feng, MahjongNumber_Zhong)
	G2_Fa            = newMahjong(2, MahjongColor_Feng, MahjongNumber_Fa)
	G2_Bai           = newMahjong(2, MahjongColor_Feng, MahjongNumber_Bai)
	G2_Plum          = newMahjong(2, MahjongColor_Hua, MahjongNumber_Plum)
	G2_Orchid        = newMahjong(2, MahjongColor_Hua, MahjongNumber_Orchid)
	G2_Chrysanthemum = newMahjong(2, MahjongColor_Hua, MahjongNumber_Chrysanthemum)
	G2_Bamboo        = newMahjong(2, MahjongColor_Hua, MahjongNumber_Bamboo)
	G2_Spring        = newMahjong(2, MahjongColor_Hua, MahjongNumber_Spring)
	G2_Summer        = newMahjong(2, MahjongColor_Hua, MahjongNumber_Summer)
	G2_Autumn        = newMahjong(2, MahjongColor_Hua, MahjongNumber_Autumn)
	G2_Winter        = newMahjong(2, MahjongColor_Hua, MahjongNumber_Winter)

	G3_Wan_1         = newMahjong(3, MahjongColor_Wan, MahjongNumber_1)
	G3_Wan_2         = newMahjong(3, MahjongColor_Wan, MahjongNumber_2)
	G3_Wan_3         = newMahjong(3, MahjongColor_Wan, MahjongNumber_3)
	G3_Wan_4         = newMahjong(3, MahjongColor_Wan, MahjongNumber_4)
	G3_Wan_5         = newMahjong(3, MahjongColor_Wan, MahjongNumber_5)
	G3_Wan_6         = newMahjong(3, MahjongColor_Wan, MahjongNumber_6)
	G3_Wan_7         = newMahjong(3, MahjongColor_Wan, MahjongNumber_7)
	G3_Wan_8         = newMahjong(3, MahjongColor_Wan, MahjongNumber_8)
	G3_Wan_9         = newMahjong(3, MahjongColor_Wan, MahjongNumber_9)
	G3_Tong_1        = newMahjong(3, MahjongColor_Tong, MahjongNumber_1)
	G3_Tong_2        = newMahjong(3, MahjongColor_Tong, MahjongNumber_2)
	G3_Tong_3        = newMahjong(3, MahjongColor_Tong, MahjongNumber_3)
	G3_Tong_4        = newMahjong(3, MahjongColor_Tong, MahjongNumber_4)
	G3_Tong_5        = newMahjong(3, MahjongColor_Tong, MahjongNumber_5)
	G3_Tong_6        = newMahjong(3, MahjongColor_Tong, MahjongNumber_6)
	G3_Tong_7        = newMahjong(3, MahjongColor_Tong, MahjongNumber_7)
	G3_Tong_8        = newMahjong(3, MahjongColor_Tong, MahjongNumber_8)
	G3_Tong_9        = newMahjong(3, MahjongColor_Tong, MahjongNumber_9)
	G3_Tiao_1        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_1)
	G3_Tiao_2        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_2)
	G3_Tiao_3        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_3)
	G3_Tiao_4        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_4)
	G3_Tiao_5        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_5)
	G3_Tiao_6        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_6)
	G3_Tiao_7        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_7)
	G3_Tiao_8        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_8)
	G3_Tiao_9        = newMahjong(3, MahjongColor_Tiao, MahjongNumber_9)
	G3_Dong          = newMahjong(3, MahjongColor_Feng, MahjongNumber_Dong)
	G3_Nan           = newMahjong(3, MahjongColor_Feng, MahjongNumber_Nan)
	G3_Xi            = newMahjong(3, MahjongColor_Feng, MahjongNumber_Xi)
	G3_Bei           = newMahjong(3, MahjongColor_Feng, MahjongNumber_Bei)
	G3_Zhong         = newMahjong(3, MahjongColor_Feng, MahjongNumber_Zhong)
	G3_Fa            = newMahjong(3, MahjongColor_Feng, MahjongNumber_Fa)
	G3_Bai           = newMahjong(3, MahjongColor_Feng, MahjongNumber_Bai)
	G3_Plum          = newMahjong(3, MahjongColor_Hua, MahjongNumber_Plum)
	G3_Orchid        = newMahjong(3, MahjongColor_Hua, MahjongNumber_Orchid)
	G3_Chrysanthemum = newMahjong(3, MahjongColor_Hua, MahjongNumber_Chrysanthemum)
	G3_Bamboo        = newMahjong(3, MahjongColor_Hua, MahjongNumber_Bamboo)
	G3_Spring        = newMahjong(3, MahjongColor_Hua, MahjongNumber_Spring)
	G3_Summer        = newMahjong(3, MahjongColor_Hua, MahjongNumber_Summer)
	G3_Autumn        = newMahjong(3, MahjongColor_Hua, MahjongNumber_Autumn)
	G3_Winter        = newMahjong(3, MahjongColor_Hua, MahjongNumber_Winter)

	G4_Wan_1         = newMahjong(4, MahjongColor_Wan, MahjongNumber_1)
	G4_Wan_2         = newMahjong(4, MahjongColor_Wan, MahjongNumber_2)
	G4_Wan_3         = newMahjong(4, MahjongColor_Wan, MahjongNumber_3)
	G4_Wan_4         = newMahjong(4, MahjongColor_Wan, MahjongNumber_4)
	G4_Wan_5         = newMahjong(4, MahjongColor_Wan, MahjongNumber_5)
	G4_Wan_6         = newMahjong(4, MahjongColor_Wan, MahjongNumber_6)
	G4_Wan_7         = newMahjong(4, MahjongColor_Wan, MahjongNumber_7)
	G4_Wan_8         = newMahjong(4, MahjongColor_Wan, MahjongNumber_8)
	G4_Wan_9         = newMahjong(4, MahjongColor_Wan, MahjongNumber_9)
	G4_Tong_1        = newMahjong(4, MahjongColor_Tong, MahjongNumber_1)
	G4_Tong_2        = newMahjong(4, MahjongColor_Tong, MahjongNumber_2)
	G4_Tong_3        = newMahjong(4, MahjongColor_Tong, MahjongNumber_3)
	G4_Tong_4        = newMahjong(4, MahjongColor_Tong, MahjongNumber_4)
	G4_Tong_5        = newMahjong(4, MahjongColor_Tong, MahjongNumber_5)
	G4_Tong_6        = newMahjong(4, MahjongColor_Tong, MahjongNumber_6)
	G4_Tong_7        = newMahjong(4, MahjongColor_Tong, MahjongNumber_7)
	G4_Tong_8        = newMahjong(4, MahjongColor_Tong, MahjongNumber_8)
	G4_Tong_9        = newMahjong(4, MahjongColor_Tong, MahjongNumber_9)
	G4_Tiao_1        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_1)
	G4_Tiao_2        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_2)
	G4_Tiao_3        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_3)
	G4_Tiao_4        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_4)
	G4_Tiao_5        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_5)
	G4_Tiao_6        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_6)
	G4_Tiao_7        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_7)
	G4_Tiao_8        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_8)
	G4_Tiao_9        = newMahjong(4, MahjongColor_Tiao, MahjongNumber_9)
	G4_Dong          = newMahjong(4, MahjongColor_Feng, MahjongNumber_Dong)
	G4_Nan           = newMahjong(4, MahjongColor_Feng, MahjongNumber_Nan)
	G4_Xi            = newMahjong(4, MahjongColor_Feng, MahjongNumber_Xi)
	G4_Bei           = newMahjong(4, MahjongColor_Feng, MahjongNumber_Bei)
	G4_Zhong         = newMahjong(4, MahjongColor_Feng, MahjongNumber_Zhong)
	G4_Fa            = newMahjong(4, MahjongColor_Feng, MahjongNumber_Fa)
	G4_Bai           = newMahjong(4, MahjongColor_Feng, MahjongNumber_Bai)
	G4_Plum          = newMahjong(4, MahjongColor_Hua, MahjongNumber_Plum)
	G4_Orchid        = newMahjong(4, MahjongColor_Hua, MahjongNumber_Orchid)
	G4_Chrysanthemum = newMahjong(4, MahjongColor_Hua, MahjongNumber_Chrysanthemum)
	G4_Bamboo        = newMahjong(4, MahjongColor_Hua, MahjongNumber_Bamboo)
	G4_Spring        = newMahjong(4, MahjongColor_Hua, MahjongNumber_Spring)
	G4_Summer        = newMahjong(4, MahjongColor_Hua, MahjongNumber_Summer)
	G4_Autumn        = newMahjong(4, MahjongColor_Hua, MahjongNumber_Autumn)
	G4_Winter        = newMahjong(4, MahjongColor_Hua, MahjongNumber_Winter)
)
