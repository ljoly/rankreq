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
	moment := root.FindMoment(timeTokens)
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
