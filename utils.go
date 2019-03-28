package rankreq

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// FileDescribe opens a file and returns a reader
func FileDescribe(path string) (*os.File, *csv.Reader, error) {

	tsvFile, err := os.Open(path)

	if err != nil {
		return nil, nil, errors.New("Open error")
	}
	return tsvFile, csv.NewReader(bufio.NewReader(tsvFile)), nil

}

// ParseTime returns a slice of time tokens
func ParseTime(time string, r *strings.Replacer, checkLen bool) ([]string, error) {

	result := r.Replace(time)
	timeTokens := strings.Split(result, "-")
	if checkLen && len(timeTokens) != 6 {
		return nil, errors.New("Wrong format")
	}
	return timeTokens, nil
}
