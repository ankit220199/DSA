// select pivot element from array and place it to right place of it not necessary that pivot left and right is sorted
package sorting

import "fmt"

func partition(arr []int, l int, r int) int {
	pi := arr[r]
	i := 0
	pos := 0
	for i = 0; i <= r; i++ {
		if arr[i] <= pi {
			swapElement(&arr[i], &arr[pos])
			pos++
		}
	}
	return pos - 1
}
func quickSort(arr []int, l int, r int) {
	if l < r {
		pi := partition(arr, l, r)
		quickSort(arr, l, pi-1)
		quickSort(arr, pi+1, r)
	}
}
func QuickInput() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	quickSort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
