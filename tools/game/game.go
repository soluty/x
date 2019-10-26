package main

import (
	"github.com/alexeyco/simpletable"
	"strconv"
)

type Game struct {
	w            int       // 宽
	h            int       // 高
	size         int       // 大小
	grids        [][]*Grid // 二维数组(数据源)
	lines        []*Line   // 两个相邻的点连成一条线 src.content = dst.content - 1
	holes        []Point   // 洞的坐标
	initContents []GridContent
}

func NewGame(w, h int, initContents []GridContent, holes ...Point) *Game {
	g := &Game{}
	g.w = w
	g.h = h
	g.holes = holes
	g.size = w*h - len(holes)
	g.initContents = initContents
	g.Reset()
	return g
}

func (g *Game) Reset() {
	w := g.w
	h := g.h
	holes := g.holes
	c := 0
	for i := 0; i < w; i++ {
		var lines []*Grid
		for j := 0; j < h; j++ {
			c++
			if holesContain(holes, i, j) {
				lines = append(lines, nil)
			} else {
				if contentIdx := gridContain(g.initContents, i, j); contentIdx >= 0 {
					lines = append(lines, &Grid{x: i, y: j, content: g.initContents[contentIdx].content, noEdit: true})
				} else {
					lines = append(lines, &Grid{x: i, y: j, content: NOVALUE})
				}
			}
		}
		g.grids = append(g.grids, lines)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if g.grids[i][j] == nil {
				continue
			}
			for di := -1; di < 2; di++ {
				for dj := -1; dj < 2; dj++ {
					if i+di >= 0 && i+di < w && j+dj >= 0 && j+dj < h {
						if di != 0 || dj != 0 {
							if g.grids[i+di][j+dj] != nil {
								g.grids[i][j].edges = append(g.grids[i][j].edges, g.grids[i+di][j+dj])
							}
						}
					}
				}
			}
		}
	}
	g.reGenerateLines(true)
}

func (g *Game) Clear() {
	for _, value1 := range g.grids {
		for _, value := range value1 {
			if value != nil {
				value.content = NOVALUE
				value.noEdit = false
			}
		}
	}
	g.reGenerateLines(true)
}

func (g *Game) String() string {
	table := simpletable.New()
	b := true

	for j := 0; j < g.h; j++ {
		var cells []*simpletable.Cell
		for i := 0; i < g.w; i++ {
			grid := g.grids[i][j]
			if grid == nil {
				cells = append(cells, &simpletable.Cell{Text: ""})
			} else if grid.content == NOVALUE {
				if b {
					cells = append(cells, &simpletable.Cell{Text: "-" + getEdges(grid.candidate)})
				} else {
					cells = append(cells, &simpletable.Cell{Text: "-"})
				}
			} else {
				if b {
					cells = append(cells, &simpletable.Cell{Text: strconv.Itoa(grid.content) + getEdges(grid.candidate)})
				} else {
					cells = append(cells, &simpletable.Cell{Text: strconv.Itoa(grid.content)})
				}
			}
		}
		table.Body.Cells = append(table.Body.Cells, cells)
	}
	table.SetStyle(simpletable.StyleRounded)
	return table.String()
}

func holesContain(holes []Point, x, y int) bool {
	for _, value := range holes {
		if value.x == x && value.y == y {
			return true
		}
	}
	return false
}

func gridContain(holes []GridContent, x, y int) int {
	for i, value := range holes {
		if value.x == x && value.y == y {
			return i
		}
	}
	return -1
}

// 利用已知条件求出一个解, 并打印出来, 深度优先搜索
func (g *Game) autoSolution2() []*Grid {
	startX, startY, ok := g.getStartCount()
	if !ok {
		return nil
	}
	return nil
}

func (g *Game) changeGrid(x, y, content int) {
	if g.grids[x][y].noEdit {
		panic("changeG1rid")
	}
	g.grids[x][y].content = content
	g.reGenerateLines(false)
}
