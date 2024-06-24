package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	width, height int
	grid          [][]int
	apple         *Apple
	player        *Player
}

// Create Initializes the game given a width and height
func (game *Game) Create(width int, height int) {
	game.width = width
	game.height = height

	// Create the grid for the game
	grid := make([][]int, width)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, height)
	}
	game.grid = grid

	// Create the player
	game.player = &Player{
		positions: make([]Vector2, width*height),
		pieces:    1,
	}
	game.player.positions[0].x = rand.Intn(width)
	game.player.positions[0].y = rand.Intn(height)

	// Create the Apple
	game.apple = &Apple{
		pos: Vector2{rand.Intn(width), rand.Intn(height)},
	}

	for game.apple.pos.x == game.player.positions[0].x && game.apple.pos.y == game.player.positions[0].y {
		game.apple.pos.x = rand.Intn(width)
		game.apple.pos.y = rand.Intn(height)
	}
}

// Render Draws the game to the console
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

func (g *Game) RenderGame() {
	for i := 0; i < g.width; i++ {
		fmt.Print("#")
	}
}

func PrintTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {

		}
	}
}

type Player struct {
	positions []Vector2
	pieces    int
	direction int
}

type Apple struct {
	pos Vector2
}
