package day02

import "testing"

func TestProductDepthHorizontalPosition(t *testing.T) {
	got, _ := ProductDepthHorizontalPosition("sample_input_01.txt")
	want := 150

	if got != want {
		t.Error("Got: %v \n Want: %v", got, want)
	}
}
