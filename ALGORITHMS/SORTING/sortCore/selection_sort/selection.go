// select min number from unsorted array and place it to right place in sorted array .
package selection_st

import "fmt"

func swapElement(a *int, b *int) {

	temp := *a
	*a = *b
	*b = temp
}
func Selection_sort() {
	//for taking input for testcases
	var t int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		var n int
		var minIdx int
		fmt.Scanln(&n)
		//declaring slice because if you create array and try to allocate size to array in this case at compile time go throw error so
		//this is not allow in golang .
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		for i := 0; i < n-1; i++ {
			minIdx = i
			for j := i + 1; j < n; j++ {
				if arr[j] < arr[minIdx] {
					minIdx = j
				}
			}
			if minIdx != i {
				swapElement(&arr[minIdx], &arr[i])
			}
		}
		for i := 0; i < n; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println()
	}
}
