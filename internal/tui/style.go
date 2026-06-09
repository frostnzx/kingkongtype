package tui

import "github.com/charmbracelet/lipgloss"

const (
	kongYellow = lipgloss.Color("#FFD700")
)

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
