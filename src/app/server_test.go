package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Setenv("GIN_MODE", "test")
	stopGinLogging()

	server := NewServer()
	assert.NotNil(t, server)
}

func TestServerOpen(t *testing.T) {
	t.Setenv("GIN_MODE", "test")
	stopGinLogging()

	server := NewServer()
	var err error
	go func() {
		err = server.Open()
	}()
	time.Sleep(1 * time.Second)
	assert.NoError(t, err)
}

func TestGetEngine(t *testing.T) {
	stopGinLogging()
	server := NewServer()
	engine := server.Engine()
	assert.NotNil(t, engine)
}
