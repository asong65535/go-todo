package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	task := "test"
	got := Add(task)

	var want error = nil

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
