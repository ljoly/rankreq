package rankreq

// Query represents a query
type Query struct {
	Str   string `json:"query"`
	Count int    `json:"count"`
}

// Queries is a map of queries
type Queries map[string]int

// Add adds a Query to a map of queries
func (tree *Queries) Add(value string) {

	if len(*tree) == 0 {
		*tree = make(Queries)
	}
	(*tree)[value] = 1
}
