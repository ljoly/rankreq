package rankreq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// BrowseForMoment returns the Moment corresponding to a time range
func (root *Moment) BrowseForMoment(timeTokens []int64) *Moment {

	currentMoment := root
	i := 0
	for _, timeToken := range timeTokens {
		if found := currentMoment.children.Find(timeToken); found != nil {
			currentMoment = found
			i++
		}
	}
	if i < len(timeTokens) {
		return nil
	}
	return currentMoment
}

type countResponse struct {
	Count int `json:"count"`
}

// CountQueries returns the number of distinct queries done during a specific time range
func (root *Moment) CountQueries(w http.ResponseWriter, r *http.Request) {

	// Get time tokens
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	rawMoment := r.URL.Path[len("/1/queries/count/"):]
	if len(rawMoment) > 19 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	replacer := strings.NewReplacer(" ", "-", ":", "-")
	timeTokens, _ := ParseTime(rawMoment, replacer, false)

	// Search
	response := countResponse{}
	startIndex := time.Now()
	moment := root.BrowseForMoment(timeTokens)
	if moment == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	response.Count = moment.count
	fmt.Printf("%-30s%s\n", "Search count:", time.Since(startIndex))

	// Response
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
