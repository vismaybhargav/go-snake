package main

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	game      *Game
	gameState int
}

const (
	RUNNING = 0
	WON     = 1
	LOST    = 2
	PAUSED  = 3
)

func InitialModel() Model {
	var g Game
	g.Create(20, 20)

	return Model{
		game:      &g,
		gameState: PAUSED,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q": // Quit
			return m, tea.Quit
		case "up", "k":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	var s string
	// Draw the top line border
	for i := 0; i < m.game.height; i++ {
		s += "#"
		for j := 0; j < m.game.width; j++ {
			if m.game.grid[j][i] == 0 {
				s += " "
			}
		}
		s += "#\n"
	}

	// Draw the bottom line border
	for i := 0; i < m.game.width+2; i++ {
		s += "#"
	}
	s += "\n"

	return s
}
