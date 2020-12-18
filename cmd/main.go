package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"pulley.com/shakesearch/pkg/search"
)

func main() {
	works := flag.String("file", "completeworks.txt", "Complete works of Shakespeare file location")
	port := flag.String("port", os.Getenv("PORT"), "http port")
	flag.Parse()

	if *port == "" {
		*port = "3001"
	}

	handler, err := search.NewHandler(*works, 250)
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", search.HandleSearch(searcher))

	fmt.Printf("Listening on port %s...", *port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
