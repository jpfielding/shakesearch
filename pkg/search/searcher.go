package search

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"net/http"
	"strings"
)

// NewHandler ...
func NewHandler(works string, width int) (http.HandlerFunc, error) {
	searcher := searcher{Width: 250}
	err := searcher.load(works)

	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results := searcher.search(query[0])
		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}, err
}

type searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
	Titles        []string
	Width         int
}

func (s *searcher) load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.SuffixArray = suffixarray.New(dat)
	s.Titles = readTitles(dat)
	return nil
}

func (s *searcher) search(query string) []string {
	results := []string{}
	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	for _, idx := range idxs {
		results = append(results, s.CompleteWorks[idx-s.Width:idx+s.Width])
	}
	return results
}

func readTitles(dat []byte) []string {
	var titles []string
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		trimmed := strings.TrimSpace(scanner.Text())
		if strings.ToLower(trimmed) != "contents" {
			continue
		}
		for scanner.Scan() {
			trimmed := strings.TrimSpace(scanner.Text())
			if trimmed == "" {
				continue
			}
			if len(titles) > 0 && trimmed == titles[0] {
				return titles
			}
			titles = append(titles, trimmed)
			fmt.Printf("found work: %s \n", trimmed)
		}
	}
	return titles
}
