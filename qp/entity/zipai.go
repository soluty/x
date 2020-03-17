package entity

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/soluty/x/qp/pb"
)

//黑牌：一、三、四、五、六、八、九各4张。壹、叁、肆、伍、陆、捌、玖各4张。
//红牌：二、七、十各4张。贰、柒、拾各4张。

type ZipaiColor byte
type ZipaiNumber byte
type ZipaiSize byte

const (
	ZipaiColor_Hong ZipaiColor = 1
	ZipaiColor_Hei  ZipaiColor = 2
)

const (
	ZipaiNumber_1  ZipaiNumber = 1
	ZipaiNumber_2  ZipaiNumber = 2
	ZipaiNumber_3  ZipaiNumber = 3
	ZipaiNumber_4  ZipaiNumber = 4
	ZipaiNumber_5  ZipaiNumber = 5
	ZipaiNumber_6  ZipaiNumber = 6
	ZipaiNumber_7  ZipaiNumber = 7
	ZipaiNumber_8  ZipaiNumber = 8
	ZipaiNumber_9  ZipaiNumber = 9
	ZipaiNumber_10 ZipaiNumber = 10
)

const (
	ZipaiSize_Xiao ZipaiSize = 1
	ZipaiSize_Da   ZipaiSize = 2
)

type zipai struct {
	id     uint32
	group  byte
	color  ZipaiColor
	number ZipaiNumber
	size   ZipaiSize
}

func newZipai(g byte, c ZipaiColor, n ZipaiNumber, s ZipaiSize) *zipai {
	p := &zipai{
		group:  g,
		color:  c,
		number: n,
		size:   s,
	}
	p.id = binary.BigEndian.Uint32([]byte{p.group, byte(p.size), byte(p.color), byte(p.number)})
	zipais[p.id] = p
	return p
}

var numStrings = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}

func (p *zipai) String() string {
	size := ""
	switch p.size {
	case ZipaiSize_Da:
		size = "大"
	case ZipaiSize_Xiao:
		size = "小"
	}
	num := numStrings[int(p.number)-1]
	return fmt.Sprintf("%d_%v%v", p.group, size, num)
}

type Zipai interface {
	Id() uint32
	Group() byte
	Color() ZipaiColor
	Number() ZipaiNumber
	Size() ZipaiSize
	From(id uint32) (Zipai, error)
	Proto() *pb.Zipai
}

var zipais = map[uint32]*zipai{}

func (p *zipai) From(id uint32) (Zipai, error) {
	if p, ok := zipais[id]; ok {
		return p, nil
	} else {
		return nil, errors.New("没有的字牌")
	}
}

func (p *zipai) Proto() *pb.Zipai {
	return &pb.Zipai{
		Id: p.id,
	}
}

func (p *zipai) Id() uint32 {
	return p.id
}

func (p *zipai) Group() byte {
	return p.group
}

func (p *zipai) Color() ZipaiColor {
	return p.color
}

func (p *zipai) Number() ZipaiNumber {
	return p.number
}

func (p *zipai) Size() ZipaiSize {
	return p.size
}

var (
	G1_Xiao_1  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Xiao)
	G1_Xiao_2  = newZipai(1, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Xiao)
	G1_Xiao_3  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Xiao)
	G1_Xiao_4  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Xiao)
	G1_Xiao_5  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Xiao)
	G1_Xiao_6  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Xiao)
	G1_Xiao_7  = newZipai(1, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Xiao)
	G1_Xiao_8  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Xiao)
	G1_Xiao_9  = newZipai(1, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Xiao)
	G1_Xiao_10 = newZipai(1, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Xiao)
	G1_Da_1    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Da)
	G1_Da_2    = newZipai(1, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Da)
	G1_Da_3    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Da)
	G1_Da_4    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Da)
	G1_Da_5    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Da)
	G1_Da_6    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Da)
	G1_Da_7    = newZipai(1, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Da)
	G1_Da_8    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Da)
	G1_Da_9    = newZipai(1, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Da)
	G1_Da_10   = newZipai(1, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Da)

	G2_Xiao_1  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Xiao)
	G2_Xiao_2  = newZipai(2, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Xiao)
	G2_Xiao_3  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Xiao)
	G2_Xiao_4  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Xiao)
	G2_Xiao_5  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Xiao)
	G2_Xiao_6  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Xiao)
	G2_Xiao_7  = newZipai(2, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Xiao)
	G2_Xiao_8  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Xiao)
	G2_Xiao_9  = newZipai(2, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Xiao)
	G2_Xiao_10 = newZipai(2, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Xiao)
	G2_Da_1    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Da)
	G2_Da_2    = newZipai(2, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Da)
	G2_Da_3    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Da)
	G2_Da_4    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Da)
	G2_Da_5    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Da)
	G2_Da_6    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Da)
	G2_Da_7    = newZipai(2, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Da)
	G2_Da_8    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Da)
	G2_Da_9    = newZipai(2, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Da)
	G2_Da_10   = newZipai(2, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Da)

	G3_Xiao_1  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Xiao)
	G3_Xiao_2  = newZipai(3, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Xiao)
	G3_Xiao_3  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Xiao)
	G3_Xiao_4  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Xiao)
	G3_Xiao_5  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Xiao)
	G3_Xiao_6  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Xiao)
	G3_Xiao_7  = newZipai(3, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Xiao)
	G3_Xiao_8  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Xiao)
	G3_Xiao_9  = newZipai(3, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Xiao)
	G3_Xiao_10 = newZipai(3, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Xiao)
	G3_Da_1    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Da)
	G3_Da_2    = newZipai(3, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Da)
	G3_Da_3    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Da)
	G3_Da_4    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Da)
	G3_Da_5    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Da)
	G3_Da_6    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Da)
	G3_Da_7    = newZipai(3, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Da)
	G3_Da_8    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Da)
	G3_Da_9    = newZipai(3, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Da)
	G3_Da_10   = newZipai(3, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Da)

	G4_Xiao_1  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Xiao)
	G4_Xiao_2  = newZipai(4, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Xiao)
	G4_Xiao_3  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Xiao)
	G4_Xiao_4  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Xiao)
	G4_Xiao_5  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Xiao)
	G4_Xiao_6  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Xiao)
	G4_Xiao_7  = newZipai(4, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Xiao)
	G4_Xiao_8  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Xiao)
	G4_Xiao_9  = newZipai(4, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Xiao)
	G4_Xiao_10 = newZipai(4, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Xiao)
	G4_Da_1    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_1, ZipaiSize_Da)
	G4_Da_2    = newZipai(4, ZipaiColor_Hong, ZipaiNumber_2, ZipaiSize_Da)
	G4_Da_3    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_3, ZipaiSize_Da)
	G4_Da_4    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_4, ZipaiSize_Da)
	G4_Da_5    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_5, ZipaiSize_Da)
	G4_Da_6    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_6, ZipaiSize_Da)
	G4_Da_7    = newZipai(4, ZipaiColor_Hong, ZipaiNumber_7, ZipaiSize_Da)
	G4_Da_8    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_8, ZipaiSize_Da)
	G4_Da_9    = newZipai(4, ZipaiColor_Hei, ZipaiNumber_9, ZipaiSize_Da)
	G4_Da_10   = newZipai(4, ZipaiColor_Hong, ZipaiNumber_10, ZipaiSize_Da)
)
