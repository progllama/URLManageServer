package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkRepo(t *testing.T) {
	repo := NewLinkRepository()
	assert.NotNil(t, repo)
}

func TestLinkRepoAdd(t *testing.T) {
	repo := NewLinkRepository()
	link := &Link{}
	repo.Add(link)
	assert.Equal(t, repo.links[0], *link)
}
