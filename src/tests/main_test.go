/*
	Gumball API in Go (Version 1)
	Basic Version with Wercker
*/

package tests

import (
	"server"
	"testing"
)

func TestGumball(t *testing.T) {
	m := server.Machine()
	cnt := m.CountGumballs
	if cnt != 989 {
		t.Errorf("Count of Gumballs Incorrect, got: %d, want: %d.", cnt, 989)
	}
}
