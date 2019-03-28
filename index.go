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
		result := r.Replace(log[0])

		// timeTokens stores the 6 tokens of a Moment
		timeTokens := strings.Split(result, "-")
		if len(timeTokens) != 6 {
			return errors.New("Wrong format")
		}

		// Create a new Moment in the MomentTrie for each new time token
		currentMoment := moment
		for i, timeToken := range timeTokens {
			// Check if the time token already exists in the current MomentTrie
			if foundMoment := currentMoment.children.Find(string(timeToken)); foundMoment != nil {
				// fmt.Print("Update Moment ", foundMoment.value)
				foundMoment.Update(log[1])
				currentMoment = foundMoment
			} else {
				var isSeconds bool
				if i == 5 {
					isSeconds = true
				}
				newMoment := currentMoment.children.Add(timeToken, log[1], isSeconds)
				// fmt.Println("Moment", newMoment.value, "created")
				// fmt.Print("currentMoment ", currentMoment.value)
				currentMoment = newMoment
				// fmt.Println(" becomes newMoment", newMoment.value, "so currentMoment becomes", currentMoment.value)
				// if len(currentMoment.queries) > 0 {
				// fmt.Println("LOL")
				// }
			}
		}
	}
	tsvFile.Close()
	return nil
}
