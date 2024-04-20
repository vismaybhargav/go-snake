package main

import "fmt"

type Game struct {
	width, height int
	grid          [][]int
	a             *Apple
	p             *Player
}

func (g Game) Create(width int, height int) {
	g.width = width
	g.height = height

	p := &Player{
		x:      make([]int, width*height),
		y:      make([]int, width*height),
		pieces: 1,
	}
	g.p = p

	a := &Apple{
		x: 1,
		y: 1,
	}
	g.a = a

	var grid [][]int = make([][]int, width*height)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, height)
	}
	g.grid = grid
}

func (g Game) Render() {
	for i := 0; i < g.width+2; i++ {
		if i == 0 {
			fmt.Print("*")
		}
	}
}

type Player struct {
	x, y   []int
	pieces int
}

type Apple struct {
	x, y int
}

func main() {
	var g Game
	g.Create(20, 20)
}
