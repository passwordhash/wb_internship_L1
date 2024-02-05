package main

import "fmt"

func main() {
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}
	intersec := intersection(set1, set2)

	fmt.Println(intersec)

	set3 := map[string]bool{"hello": true, ",": true, "world": true}
	set4 := map[string]bool{",": true, "!": true, "world": true}

	intersec1 := intersection(set3, set4)

	fmt.Println(intersec1)
}

func intersection[K comparable, V bool](set1 map[K]V, set2 map[K]V) map[K]V {
	set := make(map[K]V)
	for key := range set1 {
		if _, ex := set2[key]; ex {
			set[key] = true
		}
	}

	return set
}
