package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type menuScreenModel struct {
	choice int
	width  int
	height int
}

func (m *menuScreenModel) Init() tea.Cmd {
	return nil
}

func (m *menuScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.choice--
			if m.choice < 0 {
				m.choice = 2
			}
		case "down", "j":
			m.choice++
			if m.choice > 2 {
				m.choice = 0
			}
		case "enter":
			switch m.choice {
			case 0:
				typingScreen := NewTypingScreen()
				return m, func() tea.Msg { return ChangeScreenMsg{NewModel: typingScreen} }
			case 1:
				settingScreen := NewSettingScreen()
				return m, func() tea.Msg { return ChangeScreenMsg{NewModel: settingScreen} }
			case 2:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m *menuScreenModel) View() tea.View {
	centeredTitle := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(title())

	menu := ""
	choices := []string{"Start", "Settings", "Quit"}
	for i, c := range choices {
		sColored := c
		if i == m.choice {
			sColored = lipgloss.NewStyle().Foreground(kongYellow).Bold(true).Render(c)
		}
		menu += fmt.Sprintf("%s\n", sColored)
	}
	menu += "\nPress q to quit.\n"
	signature := "Created by frostnzx\ngithub.com/frostnzx"
	centeredSig := lipgloss.NewStyle().Align(lipgloss.Center).Width(m.width).Render(signature)
	content := centeredTitle + "\n" + menu + "\n\n\n\n\n\n" + centeredSig
	return tea.NewView(content)
}

func NewMenuScreen(c int) tea.Model {
	return &menuScreenModel{choice: c}
}

func title() string {
	titleStyle := lipgloss.NewStyle().
		Foreground(kongYellow).
		Bold(true).
		MarginBottom(1)
	header1 := `
	$$\   $$\ $$$$$$\ $$\   $$\  $$$$$$\  $$\   $$\  $$$$$$\  $$\   $$\  $$$$$$\  
	$$ | $$  |\_$$  _|$$$\  $$ |$$  __$$\ $$ | $$  |$$  __$$\ $$$\  $$ |$$  __$$\ 
	$$ |$$  /   $$ |  $$$$\ $$ |$$ /  \__|$$ |$$  / $$ /  $$ |$$$$\ $$ |$$ /  \__|
	$$$$$  /    $$ |  $$ $$\$$ |$$ |$$$$\ $$$$$  /  $$ |  $$ |$$ $$\$$ |$$ |$$$$\ 
	$$  $$<     $$ |  $$ \$$$$ |$$ |\_$$ |$$  $$<   $$ |  $$ |$$ \$$$$ |$$ |\_$$ |
	$$ |\$$\    $$ |  $$ |\$$$ |$$ |  $$ |$$ |\$$\  $$ |  $$ |$$ |\$$$ |$$ |  $$ |
	$$ | \$$\ $$$$$$\ $$ | \$$ |\$$$$$$  |$$ | \$$\  $$$$$$  |$$ | \$$ |\$$$$$$  |
	\__|  \__|\______|\__|  \__| \______/ \__|  \__| \______/ \__|  \__| \______/ 

	`
	header2 := `
	$$$$$$$$\ $$\     $$\ $$$$$$$\  $$$$$$$$\ 
	\__$$  __|\$$\   $$  |$$  __$$\ $$  _____|
	   $$ |    \$$\ $$  / $$ |  $$ |$$ |      
	   $$ |     \$$$$  /  $$$$$$$  |$$$$$\    
	   $$ |      \$$  /   $$  ____/ $$  __|   
	   $$ |       $$ |    $$ |      $$ |      
	   $$ |       $$ |    $$ |      $$$$$$$$\ 
	   \__|       \__|    \__|      \________|

	`
	return lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.MarginTop(5).Render(header1),
		titleStyle.Render(header2),
	)
}
