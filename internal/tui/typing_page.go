package tui

import (
	"kingkongtype/internal/data"
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"kingkongtype/internal/writer"
)

type typingScreenModel struct {
	buffer writer.Buffer
	author string
	width  int
	height int
}

func (m *typingScreenModel) Init() tea.Cmd {
	quote, err := data.FetchQuote(0)
	if err != nil {
		log.Fatal("Error: Can't fetch quote")
	}
	m.author = quote.Author
	m.buffer = writer.NewBuffer(quote.Text)
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
		case "esc":
			menuScreen := NewMenuScreen(0)
			return m, func() tea.Msg { return ChangeScreenMsg{NewModel: menuScreen} }
		case "tab":
			newQuote, err := data.FetchQuote(0)
			if err != nil {
				log.Fatal("Error: Can't fetch quote")
			}
			m.author = newQuote.Author
			m.buffer = writer.NewBuffer(newQuote.Text)
		case "backspace":
			m.buffer.Pop()
		case "space":
			m.buffer.InsertNextChar(" ")
		default:
			m.buffer.InsertNextChar(msg.String())
		}
	}
	return m, nil
}

func (m *typingScreenModel) View() tea.View {

	text := "\nESC to main menu\n\n\n\n\n\n"
	for i := 0; i < len(m.buffer.Text); i++ {
		if i >= len(m.buffer.InputText) {
			text += greyText(string(m.buffer.Text[i]))
			continue
		}
		if m.buffer.CheckPos(i) {
			text += yellowText(string(m.buffer.Text[i]))
		} else {
			for j := i; j < len(m.buffer.InputText); j++ {
				text += redText(string(m.buffer.InputText[j]))
			}
			text += yellowText("|")
			for j := i; j < len(m.buffer.Text); j++ {
				text += greyText(string(m.buffer.Text[j]))
			}
			break
		}
	}
	centeredTitle := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(text)
	content := centeredTitle
	return tea.NewView(content)
}

func NewTypingScreen() tea.Model {
	return &typingScreenModel{}
}

func redText(s string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF746C")).
		Render(s)
}
func yellowText(s string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Render(s)
}
func greyText(s string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#C2BDB9")).
		Render(s)
}
