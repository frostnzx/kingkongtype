package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"kingkongtype/internal/domain"
)

// msg
type ChangeScreenMsg struct {
	NewModel tea.Model
}

var gset domain.Settings // global settings

type rootScreenModel struct {
	width  int
	height int
	model  tea.Model
}

func NewRootScreen() tea.Model {
	var rootModel tea.Model

	mainMenu := NewMenuScreen(0)
	rootModel = mainMenu

	gset = domain.Settings{
		Difficulty: 0,
		Duration:   0,
		Mode:       0,
	}

	return &rootScreenModel{
		model: rootModel,
	}
}

func (m *rootScreenModel) Init() tea.Cmd {

	return m.model.Init()
}
func (m *rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case ChangeScreenMsg:
		return m.SwitchScreen(msg.NewModel)
	}
	var cmd tea.Cmd
	m.model, cmd = m.model.Update(msg)
	return m, cmd
}
func (m *rootScreenModel) View() tea.View {
	v := m.model.View()
	centeredContent := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Height(m.height).Render(v.Content)

	v = tea.NewView(centeredContent)
	v.AltScreen = true
	return v
}
func (m *rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m, m.model.Init()
}
