package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	width, height int
	grid          [][]int
	a             *Apple
	p             *Player
}

// Initializes the game given a width and height
func (g *Game) Create(width int, height int) {
	g.width = width
	g.height = height

	// Create the grid for the game
	var grid [][]int = make([][]int, width)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, height)
	}
	g.grid = grid

	// Create the player
	g.p = &Player{
		x:      make([]int, width*height),
		y:      make([]int, width*height),
		pieces: 1,
	}
	g.p.x[0] = rand.Intn(width)
	g.p.y[0] = rand.Intn(height)

	// Create the Apple
	g.a = &Apple{
		x: rand.Intn(width),
		y: rand.Intn(height),
	}

	for g.a.x == g.p.x[0] && g.a.y == g.p.y[0] {
		g.a.x = rand.Intn(width)
		g.a.y = rand.Intn(height)
	}
}

func (g *Game) Render() {
	for i := 0; i < g.width+2; i++ {
		fmt.Print("#")
	}
	fmt.Println("")

	for i := 0; i < g.height; i++ {
		fmt.Print("#")
		for j := 0; j < g.width; j++ {
			fmt.Print(" ")
		}
		fmt.Println("#")
	}

	for i := 0; i < g.width+2; i++ {
		fmt.Print("#")
	}
	fmt.Println("")
}

type Player struct {
	x, y   []int
	pieces int
}

type Apple struct {
	x, y int
}

func main() {
	// init the game
	var g Game
	g.Create(40, 20)
	g.Render()
}
