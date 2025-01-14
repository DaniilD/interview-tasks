package exp

import "fmt"

func ExpSlice() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	nums2 := nums[:0]
	fmt.Println(nums2)
	nums2 = append(nums2, 5)
	fmt.Println("returned")
	fmt.Println(nums)
	fmt.Println(nums2)

}
