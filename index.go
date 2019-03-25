package rankreq

import (
	"errors"
	"os"
)

// Leaf represents a query and its childs
type Leaf struct {
	query   Query
	queries []Query
}

// Trie is a prefix tree of queries
type Trie []Leaf

// Index parses and indexes queries
func (tree Trie) Index() error {

	if len(os.Args) < 2 {
		return errors.New("No file provided")
	}

	return nil
}
