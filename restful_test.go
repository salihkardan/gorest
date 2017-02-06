package main

import "testing"

func TestSum(t *testing.T) {
	if 8 != Sum(3, 5) {
		t.Error("should be 8")
	}
}

func TestSum2(t *testing.T) {
	if 8 != Sum(3, 5) {
		t.Error("should be 8")
	}
}
