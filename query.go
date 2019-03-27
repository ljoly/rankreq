package rankreq

// Query represents a query
type Query struct {
	Value string `json:"query"`
	Count int    `json:"count"`
}

// Queries is a map of queries
type Queries map[string]*Query

// Add adds a Query to a QueryTrie
func (tree *Queries) Add(value string) {

	new := &Query{
		Value: value,
		Count: 1,
	}

	if len(*tree) == 0 {
		*tree = make(Queries)
	}
	(*tree)[value] = new
}

// Find returns a query in a QueryTrie
func (tree Queries) Find(key string) *Query {

	if _, found := tree[key]; found {
		return tree[key]
	}
	return nil
}
