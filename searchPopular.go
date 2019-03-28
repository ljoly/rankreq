package rankreq

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// compareQuery compares current Query to the
func compareQuery(momentQuery *Query, topQueries Queries) {

	for {
		i := 0
		for _, topQuery := range topQueries {
			if topQuery.Count < momentQuery.Count {
				topQuery = momentQuery
				break
			}
			i++
		}
		if i == len(topQueries) {
			fmt.Println(topQueries)

			break
		}
	}
}

// browsePopular recursively browses the tree from a particular Moment and returns top popular queries
func (moment *Moment) browsePopular(topQueries Queries, size int) {

	if len(moment.children) > 0 {
		for _, child := range moment.children {
			child.browsePopular(topQueries, size)
		}
	} else if moment.isSeconds {
		for str, momentQuery := range moment.queries {
			if len(topQueries) < size {
				topQueries[str] = momentQuery
			} else if found := topQueries.Find(momentQuery.Value); found != nil {
				topQueries[momentQuery.Value].Count++
			} else {
				fmt.Println("Need to compare")
				compareQuery(momentQuery, topQueries)
			}
		}
	}
}

type popularResponse struct {
	Queries Queries `json:"queries"`
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
	response := popularResponse{
		Queries: make(Queries, size),
	}
	// startIndex := time.Now()
	moment := root.BrowseTrie(timeTokens)
	if moment == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	moment.browsePopular(response.Queries, int(size))

	fmt.Println("response:")
	for k, v := range response.Queries {
		fmt.Println(k, ":", v)
	}

	// // Response
	// json, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(json)
}
