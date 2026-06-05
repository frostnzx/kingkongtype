package tui

import (
	"fmt"
	"kingkongtype/internal/data"
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type typingScreenModel struct {
	text   string
	author string
	width  int
	height int
}

func (m *typingScreenModel) Init() tea.Cmd {
	quote, err := data.FetchQuote(0)
	if err != nil {
		log.Fatal("Error: Can't fetch a quote")
	}
	m.text = quote.Text
	m.author = quote.Author
	return nil
}

func (m *typingScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			menuScreen := NewMenuScreen(0)
			return m, func() tea.Msg { return ChangeScreenMsg{NewModel: menuScreen} }
		}
	}
	return m, nil
}

func (m *typingScreenModel) View() tea.View {
	centeredTitle := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(m.GameInit())

	// final
	content := centeredTitle
	return tea.NewView(content)
}

func NewTypingScreen() tea.Model {
	return &typingScreenModel{}
}

func (m *typingScreenModel) GameInit() string {
	result := fmt.Sprintf("\n\n\n%s\n\n\t\t%s", m.text, m.author)
	return result
}
