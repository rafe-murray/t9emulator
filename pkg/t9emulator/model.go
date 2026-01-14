package t9emulator

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rafe-murray/t9emulator/pkg/util"
)

type model struct {
	dictionary        *util.Trie
	lookupStart       *util.TrieNode
	currentWord       []byte
	completionOptions []string
	selected          int
	message           string
}

func NewModel() (*model, error) {
	// TODO: make this a better file location
	file, error := os.Open("dictionary.txt")
	if error != nil {
		return nil, error
	}
	dictionary, error := util.NewTrie(file)
	if error != nil {
		return nil, error
	}

	return &model{
		dictionary: dictionary,
	}, nil
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) resetCompletion() {
	m.completionOptions = []string{}
	m.selected = 0
	m.lookupStart = nil
	m.currentWord = []byte{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		s := msg.String()
		// Cool, what was the actual key pressed?
		switch s {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.currentWord = append(m.currentWord, s[0])
			var err error
			m.completionOptions, m.lookupStart, err = m.dictionary.Lookup(m.currentWord, nil)
			if err != nil {
				fmt.Printf("Error: %v", err)
				return m, tea.Quit
			}

		case "0":
			m.message = m.message + m.completionOptions[m.selected] + " "
			m.resetCompletion()
		case "*":
		case "#":
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	var s strings.Builder
	s.WriteString(m.message + string(m.currentWord))

	// Iterate over our choices
	for i, choice := range m.completionOptions {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.selected == i {
			cursor = ">" // cursor!
		}

		// Render the row
		fmt.Fprintf(&s, "%s %s\n", cursor, choice)
	}

	// Send the UI for rendering
	return s.String()
}
