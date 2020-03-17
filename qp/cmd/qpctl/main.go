package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/soluty/x/qp/entity"
	"golang.org/x/image/colornames"
	"log"
	"net/http"
)

var chupaiBtn = NewButton(150, 80)

func update(screen *ebiten.Image) error {
	//if ebiten.IsDrawingSkipped() {
	//	return nil
	//}
	//for range make([]int, 17) {
	//	text.Draw(screen, "Hello,我 W♠orld!我曹 ", face, 100,100, colornames.Red)
	//}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		idx := getClickIndex(x, y)
		if idx != -1 {
			myPokers[idx].Toggle()
		}
		if chupaiBtn.isIn(x, y) {
			clickedChupai()
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	http.Serve()

	chupaiBtn.Draw(screen)

	text.Draw(screen, errText, face, 100, 100, colornames.Red)

	for _, value := range myPokers {
		value.Draw(screen)
	}

	return nil
}

func main() {
	chupaiBtn.x = 590
	chupaiBtn.y = 600
	changeShoupai(entity.G1_Xiaogui, entity.G1_Dagui, entity.G1_Blank, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_3, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2, entity.G1_Fangpian_2)
	if err := ebiten.Run(update, 1280, 768, .5, "棋牌游戏"); err != nil {
		log.Fatal(err)
	}
}
