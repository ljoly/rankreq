package rankreq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type countResponse struct {
	Count int `json:"count"`
}

// Browse returns the Count of a Moment corresponding to a time range
func (root *Moment) Browse(timeTokens []string) int {

	currentMoment := root
	i := 0
	for _, timeToken := range timeTokens {
		if found := currentMoment.children.Find(timeToken); found != nil {
			currentMoment = found
			i++
		}
	}
	if i < len(timeTokens) {
		return 0
	}
	return currentMoment.count
}

// CountQueries returns the number of distinct queries done during a specific time range
func (root *Moment) CountQueries(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	moment := r.URL.Path[len("/1/queries/count/"):]
	len := len(moment)
	if len > 19 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	response := countResponse{}
	replacer := strings.NewReplacer(" ", "-", ":", "-")
	timeTokens, _ := ParseTime(moment, replacer, false)
	var err error

	// Search
	startIndex := time.Now()
	response.Count = root.Browse(timeTokens)
	fmt.Printf("%-30s%s\n", moment, time.Since(startIndex))
	if response.Count == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

type popularResponse struct {
	Queries []Query `json:"queries"`
}

// PopularQueries returns the top popular queries done during a specific time range
func (root *Moment) PopularQueries(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}
