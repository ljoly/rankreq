package rankreq

// Moment represents one of the 6 Moment tokens (yyyy-mm-dd-hh-mm-ss) and its children.
// It contains a Query(ies) if its value contains seconds.
type Moment struct {
	isSeconds bool
	value     int64
	count     int
	children  MomentTrie
	queries   Queries
}

// MomentTrie is a map of Moments
type MomentTrie map[int64]*Moment

// Add adds a Moment to a MomentTrie and returns it
func (tree *MomentTrie) Add(value int64, query string, isSeconds bool) *Moment {

	new := &Moment{
		isSeconds: isSeconds,
		value:     value,
		count:     1,
	}
	if isSeconds {
		new.queries.Add(query)
	}
	if *tree == nil {
		*tree = make(MomentTrie)
	}
	(*tree)[value] = new

	return (*tree)[value]
}

// FindChild returns a Moment from a MomentTrie
func (tree MomentTrie) FindChild(key int64) *Moment {

	if _, found := tree[key]; found {
		return tree[key]
	}
	return nil
}

// Update updates a Moment and its Queries
func (moment *Moment) Update(query string) {
	moment.count++
	if moment.isSeconds {
		if _, foundQuery := moment.queries[query]; foundQuery {
			moment.queries[query]++
		} else {
			moment.queries.Add(query)
		}
	}
}

// FindMoment returns the Moment corresponding to a time range
func (moment *Moment) FindMoment(timeTokens []int64) *Moment {

	currentMoment := moment
	i := 0
	for _, timeToken := range timeTokens {
		if found := currentMoment.children.FindChild(timeToken); found != nil {
			currentMoment = found
			i++
		}
	}
	if i < len(timeTokens) {
		return nil
	}
	return currentMoment
}

// Browse recursively browses the tree to get all the queries of a particular time range
func (moment Moment) Browse(rangeQueries *Queries) {

	if len(moment.children) > 0 {
		for _, child := range moment.children {
			child.Browse(rangeQueries)
		}
	} else if moment.isSeconds {
		for str, count := range moment.queries {
			rangeQueries.GetAll(str, count)
		}
	}
}
