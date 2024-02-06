package main

import "fmt"

// Худшее время O(n2)
// Лучшее время	O(n log n)
// Среднее время O(n log n)
func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var (
		left  []int
		right []int
		eq    []int
	)

	for _, x := range arr {
		if x == pivot {
			eq = append(eq, x)
		} else if x > pivot {
			right = append(right, x)
		} else if x < pivot {
			left = append(left, x)
		}
	}

	return append(quicksort(left), append(eq, quicksort(right)...)...)
}

func main() {
	arr := []int{5, 4, 3, 8, -1, 5}

	sorted := quicksort(arr)

	fmt.Println(sorted)
}
