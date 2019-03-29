package rankreq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// BrowseQueries recursively browses the tree to get all the queries of a particular time range
func (rangeQueries Queries) BrowseQueries(moment Moment) {

	if len(moment.children) > 0 {
		for _, child := range moment.children {
			rangeQueries.BrowseQueries(*child)
		}
	} else if moment.isSeconds {
		for str, count := range moment.queries {
			if _, foundQuery := rangeQueries[str]; foundQuery {
				rangeQueries[str] += count
			} else {
				rangeQueries[str] = count
			}
		}
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
	if q["size"] == nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	size, err := strconv.ParseInt(q["size"][0], 10, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Search
	startIndex := time.Now()
	moment := root.BrowseForMoment(timeTokens)
	if moment == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	rangeQueries := make(Queries)
	rangeQueries.BrowseQueries(*moment)
	response := popularResponse{}

	for str, count := range rangeQueries {
		response.Queries = append(response.Queries, Query{Str: str, Count: count})
	}
	// sort.Slice is faster than sort.SliceStable
	sort.Slice(response.Queries, func(i, j int) bool {
		return response.Queries[i].Count > response.Queries[j].Count
	})
	response.Queries = response.Queries[:size]
	fmt.Printf("%-30s%s\n\n", ">>> Search:", time.Since(startIndex))

	// Response
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}
