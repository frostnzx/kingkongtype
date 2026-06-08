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
	header := lipgloss.NewStyle().
		Width(m.width).
		MarginTop(4).
		Align(lipgloss.Left).
		Render("ESC to main menu")

	body := ""
	for i := 0; i < len(m.buffer.Text); i++ {
		if i >= len(m.buffer.InputText) {
			if i == len(m.buffer.InputText) {
				body += yellowText("|")
			}
			body += greyText(string(m.buffer.Text[i]))
			continue
		}

		if m.buffer.CheckPos(i) {
			body += yellowText(string(m.buffer.Text[i]))
		} else {
			for j := i; j < len(m.buffer.InputText); j++ {
				body += redText(string(m.buffer.InputText[j]))
			}
			body += yellowText("|")
			for j := i; j < len(m.buffer.Text); j++ {
				body += greyText(string(m.buffer.Text[j]))
			}
			break
		}
	}

	centeredBody := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		MarginTop(6).
		Render(body)

	return tea.NewView(header + "\n" + centeredBody)
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
