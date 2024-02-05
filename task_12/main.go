package main

import "fmt"

type Set struct {
	s map[string]bool
}

func NewSetOfSlice(slice []string) *Set {
	set := make(map[string]bool)

	for _, w := range slice {
		if _, ex := set[w]; !ex {
			set[w] = true
		}
	}

	return &Set{set}
}

func (s Set) String() string {
	str := "{ "

	for k := range s.s {
		str += k + " "
	}

	str += "}"

	return str
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree", "dog"}
	set := NewSetOfSlice(words)

	fmt.Println(set)
}
