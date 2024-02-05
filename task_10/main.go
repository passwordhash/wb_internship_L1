package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int]map[float64]bool)

	for _, temp := range temps {
		key := int(math.Floor(temp/10.0)) * 10

		if subset := groups[key]; subset == nil {
			subset = make(map[float64]bool)
			groups[key] = subset
		}

		groups[key][temp] = true
	}

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, key := range keys {
		group := groups[key]

		fmt.Printf("%d:{", key)

		isFrist := true
		for temp := range group {
			if isFrist {
				fmt.Print(temp)
				isFrist = false
				continue
			}
			fmt.Printf(", %.1f", temp)
		}

		fmt.Print("}")
		if i != len(keys)-1 {
			fmt.Print(", ")
		}
	}
}
