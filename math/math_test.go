package math

import "testing"

func TestAdd(t *testing.T) {

	got := Add(14, 6)
	want := 110

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
