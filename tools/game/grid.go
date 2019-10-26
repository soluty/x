package main

import (
	"fmt"
)

const NOVALUE = -100



type (
	Point struct {
		x int
		y int
	}
	Line struct {
		src *Grid
		dst *Grid
	}
)

type GridContent struct {
	x         int     // x 坐标
	y         int     // y 坐标
	content  int
}

type Grid struct {
	x         int     // x 坐标
	y         int     // y 坐标
	noEdit    bool    // 不可编辑content不能修改
	edges     []*Grid // 所有相邻的格子  不变量
	candidate []*Grid // 相邻但是可编辑的格子

	content  int
	children []*Grid // 与它相邻并且内容为content+1的格子
}

func (this *Grid) String() string {
	return fmt.Sprintf("(%v,%v| %v %v %v)", this.x, this.y, this.content, this.noEdit, this.candidate)
}

func getEdges(s []*Grid) string {
	r := ""
	for _, g := range s {
		r += fmt.Sprintf("(%v,%v), ", g.x, g.y)
	}
	return r
}


func (g *Game) reGenerateLines(isInit bool) {
	g.lines = []*Line{}
	for _, value1 := range g.grids {
		for _, value := range value1 {
			if value == nil {
				continue
			}
			value.children = nil
			for _, grid := range value.edges {
				if grid.content-value.content == 1 {
					g.lines = append(g.lines, &Line{grid, value})
					value.children = append(value.children, grid)
				}
			}
		}
	}
	if isInit {
		for i := 0; i < g.w; i++ {
			for j := 0; j < g.h; j++ {
				if g.grids[i][j] == nil {
					continue
				}
				g.grids[i][j].regenerateCandidate()
			}
		}
	}
}



func (g *Game) initGrid(x, y, content int) {
	g.grids[x][y].content = content
	g.grids[x][y].noEdit = true
	g.reGenerateLines(true)
}

func containsArrIndex(arr []int, v int) int {
	for key, value := range arr {
		if value == v {
			return key
		}
	}
	return -1
}

// 找到startGrid， startGrid的child长度为1
// 继续寻找startGrid.children[0]的孩子，确定它们的长度均为1，最后size的孩子长度为0
func (g *Game) isOver() bool {
	startX, startY, ok := g.getStartCount()
	if !ok {
		//panic("isOver")
		return false
	}
	if !hasOneChild(g.grids[startX][startY], g.size) {
		return false
	}
	return true
}

func hasOneChild(grid *Grid, end int) bool {
	if grid.content == end {
		if len(grid.children) != 0 {
			return false
		}
		return true
	} else {
		if len(grid.children) != 1 {
			return false
		}
		return hasOneChild(grid.children[0], end)
	}
}

func (g *Game) getStartCount() (int, int, bool) {
	var startX, startY, startCount = -1, -1, 0
	for i, value := range g.grids {
		for j, grid := range value {
			if grid == nil {
				continue
			}
			if grid.content == 1 {
				startCount++
				startX = i
				startY = j
				if startCount > 1 {
					return startX, startY, false
				}
			}
		}
	}
	if startCount == 1 {
		return startX, startY, true
	} else {
		return startX, startY, false
	}
}

//deprecated: 时空复杂度太高，只适合求3*3的 利用已知条件求出一个解, 并打印出来, 全排列的算法，只能算3*3的情况
//func (g *Game) autoSolution() bool {
//	var validateContents []int
//	for i := 1; i <= g.size; i++ {
//		if containsArrIndex(g.initGridContents, i) < 0 {
//			validateContents = append(validateContents, i)
//		}
//	}
//	var validateGrids []Point
//	for y, value := range g.grids {
//		for x, v := range value {
//			if v != nil && v.content < 0 {
//				validateGrids = append(validateGrids, Point{x, y})
//			}
//		}
//	}
//	if len(validateContents) != len(validateGrids) {
//		panic("autoSolution")
//		return false
//	}
//	for _, value := range pailie(len(validateContents)) {
//		for idx, grid := range validateGrids {
//			g.grids[grid.y][grid.x].content = validateContents[value[idx]]
//		}
//		g.reGenerateLines(false)
//		if g.isOver() {
//			fmt.Println(g)
//			return true
//		} else {
//			//fmt.Println(g)
//		}
//	}
//	return false
//}



func (g *Game) xunhuan(grid *Grid) bool {
	if grid.content == g.size {
		return true
	}
	var nextIndex int
	var nextCount int
	var validGrids []*Grid
	for idx, value := range grid.edges {
		if value.content == NOVALUE {
			validGrids = append(validGrids, value)
		} else {
			if value.content-grid.content == 1 {
				nextCount++
				nextIndex = idx
			}
		}
	}
	if nextCount > 1 {
		return false
	}
	if nextCount == 1 {
		return g.xunhuan(grid.edges[nextIndex])
	}
	for _, testGrid := range validGrids {
		testGrid.content = grid.content + 1
		if g.xunhuan(testGrid) {
			return true
		} else {

		}
	}
	return false
}

//deprecated: 将只有一种可能的边，循环出来, 没啥用
func (g *Game) rangOnlyOneValidEdge() bool {
	var count = 0
	var mayInit []struct {
		x       int
		y       int
		content int
	}
	for i := 0; i < g.w; i++ {
		for j := 0; j < g.h; j++ {
			if g.grids[i][j] != nil && g.grids[i][j].content >= 0 {
				grid := g.grids[i][j]
				hole := 0
				var ha, hb bool
				var ei int
				for idx, v := range grid.edges {
					if v.noEdit == true {
						if v.content != NOVALUE {
							if v.content-grid.content == 1 {
								ha = true
							}
							if grid.content-v.content == 1 {
								hb = true
							}
						}
					} else {
						hole++
						ei = idx
					}
				}
				if hole == 1 {
					fmt.Println("grid ", grid)
					if !ha && !hb {
						//
					} else if ha {
						count++
						mayInit = append(mayInit, struct {
							x       int
							y       int
							content int
						}{x: grid.edges[ei].x, y: grid.edges[ei].y, content: grid.content - 1})
					} else if hb {
						count++
						mayInit = append(mayInit, struct {
							x       int
							y       int
							content int
						}{x: grid.edges[ei].x, y: grid.edges[ei].y, content: grid.content + 1})
					}
				}
			}
		}
	}
	fmt.Println(mayInit)
	for _, value := range mayInit {
		g.initGrid(value.x, value.y, value.content)
	}
	return count > 0
}

// 0 到 n-1 的全排列
func pailie(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}
	temp := pailie(n - 1)
	var ret [][]int
	for _, value := range temp {
		for i := 0; i < n; i++ {
			var newValue []int
			if i == 0 {
				newValue = append([]int{n - 1}, value...)
			} else if i == n-1 {
				newValue = append(value, n-1)
			} else {
				for k := 0; k < i; k++ {
					newValue = append(newValue, value[k])
				}
				newValue = append(newValue, n-1)
				for k := i; k < n-1; k++ {
					newValue = append(newValue, value[k])
				}
			}
			ret = append(ret, newValue)
		}
	}
	return ret
}

func (g *Game) findNext(this *Grid) bool {
	if this.content == g.w*g.h {
		return true
	}

	fixBro := this.tryFixBrotherLen()
	if fixBro > 1 {
		return false
	}
	if fixBro == 1 {
		return g.findNext(this.children[0])
	}

	return false
}

func (g *Grid) tryFixBrotherLen() int {
	r := 0
	for _, value := range g.children {
		if value.noEdit == true {
			r++
		}
	}
	return r
}

func (this *Grid) regenerateCandidate() {
	this.candidate = nil
	for _, value := range this.edges {
		if value.noEdit == false {
			this.candidate = append(this.candidate, value)
		}
	}
}
