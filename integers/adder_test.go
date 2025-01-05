package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("want %d but sum %d", expected, sum)
	}
	ExapmpleAdd()
}
func ExapmpleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
}
