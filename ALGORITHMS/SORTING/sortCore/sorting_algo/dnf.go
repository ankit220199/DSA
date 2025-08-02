// 0,1,2 sort
package sorting

import "fmt"

func dnfSort(arr []int, n int) {
	l := 0
	m := 0
	h := arr[n-1]
	for m < h {
		if arr[m] == 0 {
			swapElement(&arr[m], &arr[l])
			m++
			l++
		} else if arr[m] == 1 {
			m++
		} else {
			swapElement(&arr[m], &arr[h])
			m++
			h--
		}
	}
}
func DnfInput() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	dnfSort(arr, n)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
