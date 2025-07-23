// divide the array into subarray so untill its not easy to merge
// divide problem into sub problems until It is divisible .
package sorting

import "fmt"

func merge(arr []int, l int, mid int, r int) {
	n1 := mid + 1 - l
	n2 := r - mid
	temp1 := make([]int, n1)
	temp2 := make([]int, n2)
	for i := 0; i < n1; i++ {
		temp1[i] = arr[l+i]
	}
	for i := 0; i < n2; i++ {
		temp2[i] = arr[mid+i+1]
	}
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if temp1[i] >= temp2[j] {
			arr[k] = temp2[j]
			j++
			k++
		} else {
			arr[k] = temp1[i]
			i++
			k++
		}
	}
	for i < n1 {
		arr[k] = temp1[i]
		k++
		i++
	}
	for j < n2 {
		arr[k] = temp2[j]
		k++
		j++
	}
}
func mergeSort(arr []int, l int, r int) {

	if l < r {
		var mid int
		mid = (l + r) / 2
		mergeSort(arr, l, mid)
		mergeSort(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}
func MergeInput() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	mergeSort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
