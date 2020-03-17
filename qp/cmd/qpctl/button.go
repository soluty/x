package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/colornames"
	"image/color"
	"sync"
)

type Button struct {
	color      color.Color
	x, y, w, h int
	text       string
	img        *ebiten.Image
}

func NewButton(w, h int) *Button {
	b := &Button{
		color: colornames.White,
		w:     w,
		h:     h,
		text:  "出牌",
	}
	b.img, _ = ebiten.NewImage(b.w, b.h, ebiten.FilterDefault)
	return b
}

func (b *Button) isIn(x, y int) bool {
	return x >= b.x && x <= b.x+b.w && y >= b.y && y <= b.y+b.h
}

func (b *Button) MoveTo(x, y int) {
	b.x = x
	b.y = y
}

func (b *Button) Draw(img *ebiten.Image) {
	b.img.Fill(b.color)
	text.Draw(b.img, b.text, face, b.w/2-32, b.h/2+16, color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	img.DrawImage(chupaiBtn.img, op)
}

// 点击出牌按钮
func clickedChupai() {
	checked := getCheckedPokers()
	if len(checked) == 0 {
		printError("必须选择一张要出的牌")
		return
	}
	if len(checked) > 1 {
		printError("一次只能出一张牌")
		return
	}
	printError("")
}
var errText string
var errLock sync.Mutex

func printError(err string) {
	errLock.Lock()
	errText = err
	errLock.Unlock()
}


func getCheckedPokers() []uint32 {
	var ret []uint32
	for _, value := range myPokers {
		if value.checked {
			ret = append(ret, value.poker.Id())
		}
	}
	return ret
}
