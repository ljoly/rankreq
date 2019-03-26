package rankreq

// Trie is a prefix tree of moments
type Trie map[string]*Moment

// Add adds a Moment to a Trie
func (tree Trie) Add(moment *Moment) {

	tree[moment.Value] = moment
	// fmt.Println("Added:", moment)
	// fmt.Println("Map:")
	// for k, v := range tree {
	// 	fmt.Println(k, v, moment.Tree)
	// }
}

// Find searches a moment in a Trie and returns it
func (tree Trie) Find(timeToken string) *Moment {

	// fmt.Println("To find:", timeToken)
	if _, ok := tree[timeToken]; ok {
		// fmt.Println("Found")
		return tree[timeToken]
	}
	return nil
}
