package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	db := NewDatabase()
	assert.NotNil(t, db)
}

func TestDBOpen(t *testing.T) {
	db := NewDatabase()
	assert.Nil(t, db.Open())
}

func TestDBClose(t *testing.T) {
	db := NewDatabase()
	db.Open()
	assert.Nil(t, db.Close())
}
