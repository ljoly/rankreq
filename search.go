package rankreq

import (
	"fmt"
	"net/http"
)

type countRequest struct {
	Count int `json:"count"`
}

// CountQueries returns the number of distinct queries done during a specific time range
func (root *Moment) CountQueries(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	reqMoment := r.URL.Path[len("/1/queries/count/"):]
	fmt.Println(reqMoment)
	// Parse reqMoment

}

type popularRequest struct {
	Queries []Query `json:"queries"`
}

// PopularQueries returns the top popular queries done during a specific time range
func (root *Moment) PopularQueries(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}
