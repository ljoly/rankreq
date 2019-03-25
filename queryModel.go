package rankreq

// Query represents a query
type Query struct {
	time  string
	Query string `json:"query"`
	Count int    `json:"count"`
}
