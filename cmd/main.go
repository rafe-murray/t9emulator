package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	t9 "github.com/rafe-murray/t9emulator/pkg/t9emulator"
)

func main() {
	p := tea.NewProgram(t9.initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
