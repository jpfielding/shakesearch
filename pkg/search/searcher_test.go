package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearcher(t *testing.T) {
	searcher := Searcher{}
	err := searcher.Load("../../completeworks.txt")
	assert.Nil(t, err)

	hamlets := searcher.Search("Hamlet")
	assert.Equal(t, 111, len(hamlets))
}
