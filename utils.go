package rankreq

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
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
func ParseTime(time string, r *strings.Replacer, checkLen bool) ([]int64, error) {

	result := r.Replace(time)
	timeTokens := strings.Split(result, "-")
	if checkLen && len(timeTokens) != 6 {
		return nil, errors.New("Wrong format")
	}

	var ret []int64
	for _, timeToken := range timeTokens {
		time, err := strconv.ParseInt(timeToken, 10, 64)
		if err != nil {
			return nil, errors.New("Wrong format")
		}
		ret = append(ret, time)
	}
	return ret, nil
}
