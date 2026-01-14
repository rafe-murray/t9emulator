package util

import (
	"bufio"
	"fmt"
	"io/fs"
	"strings"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children   map[byte]*TrieNode
	validWords []string
}

var numberMappings = map[byte]byte{
	'a': '2',
	'b': '2',
	'c': '2',
	'd': '3',
	'e': '3',
	'f': '3',
	'g': '4',
	'h': '4',
	'i': '4',
	'j': '5',
	'k': '5',
	'l': '5',
	'm': '6',
	'n': '6',
	'o': '6',
	'p': '7',
	'q': '7',
	'r': '7',
	's': '7',
	't': '8',
	'u': '8',
	'v': '8',
	'x': '9',
	'y': '9',
	'z': '9',
}

func NewTrie(file fs.File) (*Trie, error) {
	trie := &Trie{
		newTrieNode(),
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trie.insert(strings.TrimSpace(line))
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	} else {
		return trie, nil
	}
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		map[byte]*TrieNode{},
		[]string{},
	}
}

// func (trie *Trie) Save(file fs.File) error

// TODO: allow users to add their own words
func (trie *Trie) insert(s string) error {
	t := trie.root
	for i := 0; i < len(s); i++ {
		b := numberMappings[s[i]]
		// fmt.Println("b: ", b)
		// fmt.Println("root.children: ", t.children)
		u, ok := t.children[b]
		if !ok {
			t.children[b] = newTrieNode()
			t = t.children[b]
		} else {
			t = u
		}
	}
	t.validWords = append(t.validWords, s)
	return nil
}

// func (trie *Trie) Remove(s string) error
func (trie *Trie) Lookup(s []byte, root *TrieNode) ([]string, *TrieNode, error) {
	var t *TrieNode
	if root == nil {
		t = trie.root
	} else {
		t = root
	}

	for i := range s {
		var ok bool
		t, ok = t.children[s[i]]
		if !ok {
			return nil, nil, fmt.Errorf("Could not find any words for given string: %s", s)
		}
	}
	return t.getSuggestions([]string{}), t, nil
}

func (node *TrieNode) getSuggestions(suggestions []string) []string {
	for _, child := range node.children {
		suggestions = child.getSuggestions(suggestions)
	}
	return append(suggestions, node.validWords...)
}
