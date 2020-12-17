package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHamlets(t *testing.T) {
	searcher := Searcher{Width: 10}
	err := searcher.Load("../../completeworks.txt")
	assert.Nil(t, err)

	hamlets := searcher.Search("Hamlet")
	assert.Equal(t, 111, len(hamlets))
}

func TestWhereForArtThous(t *testing.T) {
	searcher := Searcher{Width: 10}
	err := searcher.Load("../../completeworks.txt")
	assert.Nil(t, err)

	hamlets := searcher.Search("herefore art thou")
	assert.Equal(t, 1, len(hamlets))
}
