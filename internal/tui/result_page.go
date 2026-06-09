package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type resultScreenModel struct {
}

func NewResultScreen() tea.Model {
	return &resultScreenModel{}
}
func (s *resultScreenModel) Init() tea.Cmd {
	return nil
}
func (m *resultScreenModel) View() tea.View {
	header := `
	 /$$$$$$$                                /$$   /$$    
| $$__  $$                              | $$  | $$    
| $$  \ $$  /$$$$$$   /$$$$$$$ /$$   /$$| $$ /$$$$$$  
| $$$$$$$/ /$$__  $$ /$$_____/| $$  | $$| $$|_  $$_/  
| $$__  $$| $$$$$$$$|  $$$$$$ | $$  | $$| $$  | $$    
| $$  \ $$| $$_____/ \____  $$| $$  | $$| $$  | $$ /$$
| $$  | $$|  $$$$$$$ /$$$$$$$/|  $$$$$$/| $$  |  $$$$/
|__/  |__/ \_______/|_______/  \______/ |__/   \___/  
                                                      
                                                      

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

	s.WriteString("\n\n\n\n\nPress q to return to menu.")

	return tea.NewView(s.String())
}

func (s *resultScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit
		case "esc", "q":
			menuScreen := NewMenuScreen(1)
			return s, func() tea.Msg { return ChangeScreenMsg{NewModel: menuScreen} }
		}
	}
	return s, nil
}
