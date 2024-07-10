package main

import "testing"

func TestMain(t *testing.T) {
	want := 1
	if want != 1 {
		t.Error("Error")
	}
}
