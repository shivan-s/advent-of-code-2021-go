package day05

import "testing"

func Test(t *testing.T) {
	got := 0
	want := 1

	if got != want {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}
