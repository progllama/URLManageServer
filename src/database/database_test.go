package database

import "testing"

func TestOpen(t *testing.T) {
	t.Setenv("MODE", "dev")
	Open("", "")
}
