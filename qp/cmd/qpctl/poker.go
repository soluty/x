package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/soluty/x/qp/entity"
	"golang.org/x/image/colornames"
	"image/color"
)

//var win *pixelgl.Window
//var canvas *pixelgl.Canvas
//
type Poker struct {
	poker   entity.Poker
	x       int
	y       int
	color   color.Color
	checked bool
	str     string
}

//
func NewPoker(poker entity.Poker) *Poker {
	p := &Poker{
		poker: poker,
		color: colornames.White,
	}
	p.str = p.poker.Display()
	//+ "\n" + strconv.Itoa(int(p.poker.Id()))
	return p
}

func (this *Poker) Draw(screen *ebiten.Image) {
	text.Draw(screen, this.str, face, this.x, this.y, this.color)
}

//
//func (this *Poker) Draw() {
//	//mat := pixel.IM
//	//mat = mat.Moved(win.Bounds().Center())
//	//mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(1, 1))
//	//this.basicTxt.Clear()
//	this.basicTxt.Color = this.color
//	//this.basicTxt.WriteString(this.str)
//	this.basicTxt.Draw(canvas, pixel.IM.Moved(pixel.V(this.x, this.y)))
//
//}
//
//func (this *Poker) SetText() {
//	this.basicTxt.Clear()
//	this.basicTxt.WriteString(this.str)
//}
//
func (this *Poker) Toggle() {
	if !this.checked {
		this.color = colornames.Red
	} else {
		this.color = colornames.White
	}
	this.checked = !this.checked
}
