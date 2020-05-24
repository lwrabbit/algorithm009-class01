package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(MaxArea(a))
}

func TestMaxArea1(t *testing.T) {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(MaxArea1(a))
}
