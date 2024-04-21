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
	grid := make([][]int, width)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, height)
	}
	g.grid = grid

	// Create the player
	g.p = &Player{
		positions: make([]Vector2, width*height),
		pieces:    1,
	}
	g.p.positions[0].x = rand.Intn(width)
	g.p.positions[0].y = rand.Intn(height)

	// Create the Apple
	g.a = &Apple{
		pos: Vector2{rand.Intn(width), rand.Intn(height)},
	}

	for g.a.pos.x == g.p.positions[0].x && g.a.pos.y == g.p.positions[0].y {
		g.a.pos.x = rand.Intn(width)
		g.a.pos.y = rand.Intn(height)
	}
}

func (g *Game) Render() {
	// Draw the top line border
	for i := 0; i < g.width+2; i++ {
		fmt.Print("#")
	}
	fmt.Println("")

	for i := 0; i < g.height; i++ {
		fmt.Print("#")
		for j := 0; j < g.width; j++ {
			fmt.Print(g.grid[j][i]) // idk why j and i are flipped
		}
		fmt.Println("#")
	}

	// Draw the bottom line border
	for i := 0; i < g.width+2; i++ {
		fmt.Print("#")
	}
	fmt.Println("")
}

type Player struct {
	positions []Vector2
	pieces    int
}

type Apple struct {
	pos Vector2
}

func main() {
	// init the game
	var g Game
	g.Create(40, 20)
	g.Render()
}
