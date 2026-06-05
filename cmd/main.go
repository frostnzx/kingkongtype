package main

import (
	"fmt"
	"os"

	"kingkongtype/internal/tui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(tui.NewRootScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
