package rankreq

// Moment represents one of the 6 Moment tokens (yyyy-mm-dd-hh-mm-ss) and its children.
// It contains a Query(ies) if its value contains seconds.
type Moment struct {
	isSeconds bool
	value     string
	count     int
	children  MomentTrie
	queries   Queries
}

// MomentTrie is a map of Moments
type MomentTrie map[string]*Moment

// Add adds a Moment to a MomentTrie and returns it
func (tree *MomentTrie) Add(value string, query string, isSeconds bool) *Moment {

	new := &Moment{
		isSeconds: isSeconds,
		value:     value,
		count:     1,
	}
	new.queries.Add(query)
	if len(*tree) == 0 {
		*tree = make(MomentTrie)
	}
	(*tree)[value] = new

	return (*tree)[value]
}

// Find returns a Moment from a MomentTrie
func (tree MomentTrie) Find(key string) *Moment {

	if _, found := tree[key]; found {
		return tree[key]
	}
	return nil
}

// Update updates a Moment and its Queries
func (moment *Moment) Update(query string) {
	// fmt.Println(" with count++")
	moment.count++
	if moment.isSeconds {
		// fmt.Println(" and with", query)
		if foundQuery := moment.queries.Find(query); foundQuery != nil {
			foundQuery.Count++
		} else {
			moment.queries.Add(query)
		}
	}
}
