package main

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	_ = tea.NewProgram(InitialModel())

	// init the game
	var g Game
	g.Create(40, 20)
	g.Render()
}
