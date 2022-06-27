package main

import "fmt"

func main() {
	s1 := []int{1,2,3,4,5}
	s2 := []int{1,2,3,4,6}
	fmt.Println(equal(s1, s2))
}

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
