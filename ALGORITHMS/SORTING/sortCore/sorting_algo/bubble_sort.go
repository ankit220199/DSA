// sort adjacent element of the unsorted array
package sorting

import "fmt"

func swapElement(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func Bubble_sort() {
	var t int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanln(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanln(&a[i])
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if a[j] > a[j+1] {
					swapElement(&a[j], &a[j+1])
				}
			}
		}
		for i := 0; i < n; i++ {
			fmt.Print(a[i], " ")
		}
		fmt.Println()
	}
}
