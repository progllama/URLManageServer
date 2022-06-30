package session

import (
	"testing"

	"github.com/gin-contrib/sessions/cookie"
)

func TestMakeStore(t *testing.T) {
	t.Setenv("MODE", "dev")

	store := makeStore()
	switch store.(type) {
	case cookie.Store:
	default:
		t.Fatal("MODE is dev but store is not cookie.")
	}
}
