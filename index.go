package rankreq

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

// Index parses and indexes moments and queries in a prefix tree
func (moment *Moment) Index(tsvFile *os.File, reader *csv.Reader) error {

	moment.children = make(MomentTrie)
	// Using a replacer is faster than using string.Replace()
	r := strings.NewReplacer(" ", "-", ":", "-")
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New("Read error")
		}
		// log stores the 2 parts of a query: time and query
		log := strings.Split(line[0], "\t")
		if len(log) != 2 {
			return errors.New("Wrong format")
		}

		// timeTokens stores the 6 tokens of a Moment
		timeTokens, err := ParseTime(log[0], r, true)
		if err != nil {
			return err
		}

		// Create a new Moment in the MomentTrie for each new time token
		currentMoment := moment
		for i, timeToken := range timeTokens {
			// Check if the time token already exists in the current MomentTrie
			if foundMoment := currentMoment.children.FindChild(timeToken); foundMoment != nil {
				foundMoment.Update(log[1])
				currentMoment = foundMoment
			} else {
				var isSeconds bool
				if i == 5 {
					isSeconds = true
				}
				newMoment := currentMoment.children.Add(timeToken, log[1], isSeconds)
				currentMoment = newMoment
			}
		}
	}
	tsvFile.Close()
	return nil
}
