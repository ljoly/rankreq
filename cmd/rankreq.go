package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ljoly/rankreq"
)

// CreateRoutes adds routes to the multiplexer
func CreateRoutes(mux *http.ServeMux, root rankreq.Moment) {

	mux.HandleFunc("/1/queries/count/", root.CountQueries)
	mux.HandleFunc("/1/queries/popular/", root.PopularQueries)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No file provided\n")
		os.Exit(1)
	}

	root := rankreq.Moment{
		Tree: make(rankreq.Trie),
	}
	err := root.Index()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()
	CreateRoutes(mux, root)
	fmt.Println("listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
