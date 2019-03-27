package rankreq

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
)

// FileDescribe opens a file and returns a reader
func FileDescribe(path string) (*os.File, *csv.Reader, error) {

	tsvFile, err := os.Open(path)

	if err != nil {
		return nil, nil, errors.New("Open error")
	}
	return tsvFile, csv.NewReader(bufio.NewReader(tsvFile)), nil

}
