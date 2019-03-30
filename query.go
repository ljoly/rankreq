package rankreq

// Query represents a query
type Query struct {
	Str   string `json:"query"`
	Count int    `json:"count"`
}

// Queries is a map of queries
type Queries map[string]int

// Add adds a Query to a map of queries
func (queries *Queries) Add(value string) {

	if len(*queries) == 0 {
		*queries = make(Queries)
	}
	(*queries)[value] = 1
}

// GetAll builds a Queries
func (queries *Queries) GetAll(str string, count int) {

	if _, foundQuery := (*queries)[str]; foundQuery {
		(*queries)[str] += count
	} else {
		(*queries)[str] = count
	}
}
