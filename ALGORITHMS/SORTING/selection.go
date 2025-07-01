package main

import "fmt"

func swap(a *int, b *int) {
	var temp int // *a has a type of int so we dont need to use pointer for temp
	temp = *a
	*a = *b
	*b = temp
}
func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if arr[i] > arr[j] {
					swap(&arr[i], &arr[j])
				}
			}
		}
		for i := 0; i < n; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println()
		t--
	}
}
