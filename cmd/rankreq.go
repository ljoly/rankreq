package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ljoly/rankreq"
)

func createRoutes(mux *http.ServeMux, root rankreq.Moment) {

	mux.HandleFunc("/1/queries/count/", root.CountQueries)
	mux.HandleFunc("/1/queries/popular/", root.PopularQueries)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No file provided\n")
		os.Exit(1)
	}

	// Open and get file descriptor
	tsvFile, reader, err := rankreq.FileDescribe(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	// Create the root of the prefix tree
	root := rankreq.Moment{}
	startIndex := time.Now()
	err = root.Index(tsvFile, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	// Create server
	mux := http.NewServeMux()
	createRoutes(mux, root)
	fmt.Printf("%-30s%s\n", ">>> Listening on port:", "8080")
	fmt.Printf("%-30s%s\n\n", ">>> Indexing:", time.Since(startIndex))
	// Expose api
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
