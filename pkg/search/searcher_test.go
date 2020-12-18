package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHamlets(t *testing.T) {
	searcher := searcher{Width: 10}
	err := searcher.load("../../completeworks.txt")
	assert.Nil(t, err)

	hamlets := searcher.search("Hamlet")
	assert.Equal(t, 111, len(hamlets))
}

func TestWhereForArtThous(t *testing.T) {
	searcher := searcher{Width: 10}
	err := searcher.load("../../completeworks.txt")
	assert.Nil(t, err)

	hamlets := searcher.search("herefore art thou")
	assert.Equal(t, 1, len(hamlets))
}
