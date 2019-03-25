package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ljoly/rankreq"
)

// CreateRoutes adds routes to the multiplexer
func CreateRoutes(mux *http.ServeMux, tree rankreq.Trie) {

	mux.HandleFunc("/1/queries/count/", tree.Count)
	mux.HandleFunc("/1/queries/popular/", tree.Popular)
}

func main() {

	tree := make(rankreq.Trie, 0)

	err := tree.Index()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()
	CreateRoutes(mux, tree)
	fmt.Println("listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
