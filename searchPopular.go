package rankreq

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// browsePopular recursively browses the tree from a particular Moment and returns top popular queries
func (moment *Moment) browsePopular(queries []Query) {

	if len(moment.children) > 0 {
		moment.browsePopular(queries)
	}
	if moment.isSeconds {

	}
}

type popularResponse struct {
	Queries []Query `json:"queries"`
}

// PopularQueries returns the top popular queries done during a specific time range
func (root *Moment) PopularQueries(w http.ResponseWriter, r *http.Request) {

	// Get time tokens and size
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	rawMoment := r.URL.Path[len("/1/queries/popular/"):]
	if len(rawMoment) > 19 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	replacer := strings.NewReplacer(" ", "-", ":", "-")
	timeTokens, _ := ParseTime(rawMoment, replacer, false)

	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	size, err := strconv.ParseInt(q["size"][0], 10, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Search
	response := popularResponse{
		Queries: make([]Query, size),
	}
	// startIndex := time.Now()
	moment := root.BrowseTrie(timeTokens)
	if moment == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	moment.browsePopular(response.Queries)
	fmt.Println(response.Queries)
	// // Response
	// json, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(json)
}
