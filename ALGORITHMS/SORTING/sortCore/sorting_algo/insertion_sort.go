// Take element from unsorted array and place it sorted array
package sorting

import (
	"fmt"
)

func Insertition_sort() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i])
	}
	for i := 1; i < n; i++ {
		curr := a[i]
		j := i - 1
		for j >= 0 && a[j] > curr {
			a[j+1] = a[j]
			fmt.Println(a[j])
			j--
		}
		a[j+1] = curr
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}
