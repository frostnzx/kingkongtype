package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type settingScreenModel struct {
	cursor, list  int
	width, height int
}

var choices = [3][3]string{
	{"Easy", "Medium", "Hard"},
	{"15s", "30s", "60s"},
	{"Time", "Word", "Quote"},
}

func NewSettingScreen() tea.Model {
	return &settingScreenModel{cursor: 0, list: 0}
}
func (s *settingScreenModel) Init() tea.Cmd {
	return nil
}
func (m *settingScreenModel) View() tea.View {
	header := `
	  /$$$$$$              /$$     /$$     /$$                              
	 /$$__  $$            | $$    | $$    |__/                              
	| $$  \__/  /$$$$$$  /$$$$$$ /$$$$$$   /$$ /$$$$$$$   /$$$$$$   /$$$$$$$
	|  $$$$$$  /$$__  $$|_  $$_/|_  $$_/  | $$| $$__  $$ /$$__  $$ /$$_____/
	 \____  $$| $$$$$$$$  | $$    | $$    | $$| $$  \ $$| $$  \ $$|  $$$$$$ 
	 /$$  \ $$| $$_____/  | $$ /$$| $$ /$$| $$| $$  | $$| $$  | $$ \____  $$
	|  $$$$$$/|  $$$$$$$  |  $$$$/|  $$$$/| $$| $$  | $$|  $$$$$$$ /$$$$$$$/
	 \______/  \_______/   \___/   \___/  |__/|__/  |__/ \____  $$|_______/ 
														 /$$  \ $$          
														|  $$$$$$/          
											 			\______/
	`
	// centeredHeader := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(header)
	s := strings.Builder{}

	headerStyle := lipgloss.NewStyle().
		Foreground(kongYellow).
		Bold(true).
		MarginBottom(1)

	s.WriteString("\n\n")
	s.WriteString(headerStyle.Render(header))
	s.WriteString("\n\n")

	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			s.WriteString("Choose your diffuculty : \n\n")
		case 1:
			s.WriteString("Choose your time duration : \n\n")
		case 2:
			s.WriteString("Choose your game mode : \n\n")
		}
		for j := 0; j < 3; j++ {
			prefix := "( ) "
			if (i == 0 && gset.Difficulty == j) || (i == 1 && gset.Duration == j) || (i == 2 && gset.Mode == j) {
				prefix = "(X) "
			} else if m.cursor == j && m.list == i {
				prefix = "(•) "
			}
			s.WriteString(prefix)
			s.WriteString(fmt.Sprintf("		%6s", choices[i][j]))
			s.WriteString("\n")
		}
		s.WriteString("\n\n")
	}

	s.WriteString("\n\n\n\n\nPress q to return to menu.")

	return tea.NewView(s.String())
}

func (s *settingScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit
		case "esc", "q":
			menuScreen := NewMenuScreen(1)
			return s, func() tea.Msg { return ChangeScreenMsg{NewModel: menuScreen} }
		case "enter":
			switch s.list {
			case 0:
				gset.Difficulty = s.cursor
			case 1:
				gset.Duration = s.cursor
			case 2:
				gset.Mode = s.cursor
			}
		case "down", "j":
			s.cursor++
			if s.cursor >= len(choices[s.list]) {
				s.cursor = 0
				s.list++
			}
			if s.list >= 3 {
				s.list = 0
			}
		case "up", "k":
			s.cursor--
			if s.cursor < 0 {
				s.list--
				if s.list < 0 {
					s.list = 2
				}
				s.cursor = len(choices[s.list]) - 1
			}
		}
	}

	return s, nil
}
