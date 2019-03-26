package rankreq

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

// Query represents a query
type Query struct {
	Str   string `json:"query"`
	Count int    `json:"count"`
}

// Moment represents a moment token and its children.
// It contains a Query if its value contains seconds.
type Moment struct {
	isSeconds bool
	Value     string
	Count     int
	Tree      Trie
	Query     Query
}

// Index parses and indexes queries in a prefix tree
func (root *Moment) Index() error {

	// Open and read file
	tsvFile, err := os.Open(os.Args[1])
	defer tsvFile.Close()

	if err != nil {
		return errors.New("Open error")
	}
	reader := csv.NewReader(bufio.NewReader(tsvFile))

	// Parse data
	// i := 0
	moment := root
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New("Read error")
		}
		log := strings.Fields(line[0])
		if len(log) != 3 {
			return errors.New("Wrong format")
		}

		// Get moment: there are 6 tokens to define a moment: yyyy/mm/dd/hh/mm/ss
		timeTokens := strings.Split(log[0]+"-"+strings.Replace(log[1], ":", "-", -1), "-")
		if len(timeTokens) != 6 {
			return errors.New("Wrong format")
		}
		// Create a new node in the trie for each token
		for i, token := range timeTokens {
			if found := moment.Tree.Find(string(token)); found != nil {
				found.Count++
			} else {
				// fmt.Println("NOT FOUND")
				isSeconds := false
				query := Query{}
				if i == 5 {
					isSeconds = true
					query.Str = log[2]
					query.Count++
				}
				new := &Moment{
					isSeconds: isSeconds,
					Value:     token,
					Tree:      make(Trie),
					Query:     query,
					Count:     1,
				}
				moment.Tree.Add(new)
			}
			i++
		}
		// if i == 1 {
		// 	os.Exit(0)

		// }
		// i++
	}
	return nil
}
