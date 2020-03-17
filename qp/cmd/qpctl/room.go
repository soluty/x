package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/soluty/x/qp/entity"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"io/ioutil"
	"os"
)

//
//import (
//	"fmt"
//	"github.com/faiface/pixel"
//	"github.com/faiface/pixel/pixelgl"
//	"github.com/faiface/pixel/text"
//	"github.com/golang/freetype/truetype"
//	"github.com/soluty/x/qp/entity"
//	"golang.org/x/image/colornames"
//	"golang.org/x/image/font"
//	"golang.org/x/image/font/gofont/goregular"
//	"io/ioutil"
//	"os"
//	"time"
//	"unicode"
//)
//
//// 房间 大厅 玩家 账户 游戏 房卡
//// Room Hall Player Account game Card
//var atlas *text.Atlas
//
var face font.Face
func init() {
	ff, err := loadTTF("/Library/Fonts/Arial Unicode.ttf", pokerH)
	if err != nil {
		panic(err)
	}
	face = ff
	//atlas = text.NewAtlas(face, text.ASCII, text.RangeTable(unicode.Han), text.RangeTable(unicode.P), []rune{'♥', '♠', '♣', '♦'})
	//errText = text.New(pixel.V(0, 0), atlas)
}
//
var myPokers []*Poker
var startX = 30
var startY = 500
var pokerW = 70
const pokerH = 32

//var errText *text.Text
//
func changeShoupai(ps ...entity.Poker) {
	myPokers = nil
	for idx, value := range ps {
		p := NewPoker(value)
		p.x = startX + idx * pokerW
		p.y = pokerH + startY
		myPokers = append(myPokers, p)
	}
}
//
func getClickIndex(x, y int) int {
	if y > startY+pokerH || y < startY {
		return -1
	}
	if x < startX {
		return -1
	}
	i := (x - startX) / pokerW
	if i >= len(myPokers) {
		return -1
	}
	return i
}
//
//func setError(s string) {
//	errText.Clear()
//	errText.WriteString(s)
//}
//
//func main() {
//
//	var (
//		frames = 0
//		second = time.Tick(time.Second)
//	)
//
//	cfg := pixelgl.WindowConfig{
//		Title:  "棋牌客户端",
//		Bounds: pixel.R(0, 0, 1024, 768),
//	}
//
//	changeShoupai(entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_3, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2)
//
//	setError("我错了aaa,!。")
//
//	pixelgl.Run(func() {
//		wind, err := pixelgl.NewWindow(cfg)
//		if err != nil {
//			panic(err)
//		}
//		win = wind
//		canvas = pixelgl.NewCanvas(pixel.R(0,0,1024,768))
//		fps := time.Tick(time.Second / 60)
//
//		for !win.Closed() {
//			win.Clear(colornames.Black)
//			canvas.Clear(colornames.Black)
//
//			if win.JustPressed(pixelgl.MouseButtonLeft) {
//				idx := getClickIndex(win.MousePosition())
//				if idx != -1 {
//					myPokers[idx].Toggle()
//				}
//
//			}
//
//			for _, value := range myPokers {
//				value.Draw()
//			}
//
//			canvas.Draw(win, pixel.IM)
//
//			errText.Draw(win, pixel.IM.Moved(pixel.Vec{100, 100}))
//
//			win.Update()
//			//if delta := time.Now().Sub(t); delta > 20000000 {
//			//	fmt.Println("耗时", delta)
//			//}
//			frames++
//			select {
//			case <-second:
//				win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
//				frames = 0
//			default:
//				<-fps
//			}
//		}
//	})
//}
//
func loadTTF(path string, size float64) (font.Face, error) {
	// /Users/soluty/Library/Fonts
	var bs []byte
	if path != "" {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		bs, err = ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
	} else {
		bs = goregular.TTF
	}
	f, err := truetype.Parse(bs)
	if err != nil {
		return nil, err
	}
	return truetype.NewFace(f, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
