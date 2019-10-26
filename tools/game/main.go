package main

import "fmt"

func main() {
	g := NewGame(3, 3,nil)
	//fmt.Println(g)
	//g.changeGrid(0, 0, 8)
	//g.changeGrid(0, 1, 1)
	//g.changeGrid(0, 2, 2)
	//g.changeGrid(1, 0, 9)
	//g.changeGrid(1, 1, 7)
	//g.changeGrid(1, 2, 3)
	//g.changeGrid(2, 0, 6)
	//g.changeGrid(2, 1, 5)
	//g.changeGrid(2, 2, 4)
	//fmt.Println(g, g.isOver())
	//
	//g.Clear()
	//g.initGrid(0, 0, 8)
	//g.initGrid(0, 2, 1)
	//g.initGrid(2, 0, 3)
	//g.autoSolution()
	//
	//g.Clear()
	//
	//g.initGrid(1, 1, 8)
	//g.initGrid(1, 2, 1)
	//g.initGrid(2, 1, 2)
	//
	//fmt.Println(g.autoSolution())

	g = NewGame(7, 6, nil, []Point{
		{0, 0}, {3, 0}, {6, 0},
		{0, 3}, {6, 3},
		{0, 4}, {1, 4}, {5, 4}, {6, 4},
		{0, 5}, {1, 5}, {2, 5}, {4, 5}, {5, 5}, {6, 5},
	}...)
	//g.reGenerateLines()
	//g.changeGrid(0, 1, 8)
	//g.changeGrid(0, 2, 7)
	//
	//g.changeGrid(1, 0, 9)
	//g.changeGrid(1, 1, 6)
	//g.changeGrid(1, 2, 5)
	//g.changeGrid(1, 3, 2)
	//
	//g.changeGrid(2, 0, 10)
	//g.changeGrid(2, 1, 4)
	//g.changeGrid(2, 2, 3)
	//g.changeGrid(2, 3, 25)
	//g.changeGrid(2, 4, 1)
	//
	//g.changeGrid(3, 1, 11)
	//g.changeGrid(3, 2, 26)
	//g.changeGrid(3, 3, 27)
	//g.changeGrid(3, 4, 24)
	//g.changeGrid(3, 5, 23)
	//
	//g.changeGrid(4, 0, 13)
	//g.changeGrid(4, 1, 12)
	//g.changeGrid(4, 2, 20)
	//g.changeGrid(4, 3, 21)
	//g.changeGrid(4, 4, 22)
	//
	//g.changeGrid(5, 0, 14)
	//g.changeGrid(5, 1, 15)
	//g.changeGrid(5, 2, 16)
	//g.changeGrid(5, 3, 19)
	//
	//g.changeGrid(6, 1, 17)
	//g.changeGrid(6, 2, 18)
	//
	//fmt.Println(g, g.isOver())

	g.Clear()

	//g.changeGrid(0, 1, 12)
	//g.changeGrid(0, 2, 13)

	//g.changeGrid(1, 0, 11)
	g.initGrid(1, 1, 9)
	g.initGrid(1, 2, 8)
	//g.changeGrid(1, 3, 14)

	//g.changeGrid(2, 0, 10)
	//g.changeGrid(2, 1, 5)
	g.initGrid(2, 2, 6)
	//g.changeGrid(2, 3, 7)
	//g.changeGrid(2, 4, 15)

	//g.changeGrid(3, 1, 4)
	g.initGrid(3, 2, 27)
	//g.changeGrid(3, 3, 26)
	g.initGrid(3, 4, 16)
	//g.changeGrid(3, 5, 17)

	//g.changeGrid(4, 0, 3)
	//g.initGrid(4, 1, 2)
	g.initGrid(4, 2, 1)
	//g.initGrid(4, 3, 25)
	//g.changeGrid(4, 4, 18)

	g.initGrid(5, 0, 22)
	g.initGrid(5, 1, 23)
	g.initGrid(5, 2, 24)
	//g.changeGrid(5, 3, 19)

	g.initGrid(6, 1, 21)
	//g.changeGrid(6, 2, 20)
	fmt.Println(g)

	fmt.Println(g.autoSolution2(), "\n", g)
}