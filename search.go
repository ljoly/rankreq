package rankreq

import (
	"fmt"
	"net/http"
)

type countRequest struct {
	Count int `json:"count"`
}

// Count returns the number of distinct queries done during a specific time range
func (tree Trie) Count(w http.ResponseWriter, r *http.Request) {

	date := r.URL.Path[len("/1/queries/count/"):]
	fmt.Println(date)
	// Parse date
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}

type popularRequest struct {
	Queries []Query `json:"queries"`
}

// Popular returns the top popular queries done during a specific time range
func (tree Trie) Popular(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}
