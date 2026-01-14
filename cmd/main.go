package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	t9 "github.com/rafe-murray/t9emulator/pkg/t9emulator"
)

func main() {
	model, err := t9.NewModel()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
