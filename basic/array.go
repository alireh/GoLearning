package basic

import (
	"fmt"

)

func TestArray() {
	fmt.Println("----------------------------------------Start Array----------------------------------------")

	fmt.Println("                                    ***********             A            *********\n")
	arrayInts := [5]int{1, 25, 12354, 654, 32}
	fmt.Println(arrayInts)

	fmt.Println("                                    ***********             B            *********\n")
	nums := [5]int{}
	fmt.Printf("%v, len : %d, cap : %d\n", nums, len(nums), cap(nums))

	nums[0] = 1
	nums[1] = 2
	nums[2] = 7
	//num[3] = 4
	nums[4] = 9

	fmt.Printf("%v, len : %d, cap : %d\n", nums, len(nums), cap(nums))

	// آرایه با اندازه تعیین شده توسط کامپایلر
	fmt.Println("                                    ***********             C            *********\n")
	nums1 := [...]int{1, 25, 45, 8797, 78, 879, 541, 11}
	fmt.Printf("%v, len %d, cap %d\n", nums1, len(nums1), cap(nums1))

	fmt.Println("                                    ***********             D            *********\n")
    //تعریف آرایه دو بعدی یا چند بعدی
	nums3 := [2][2][2]int{{{1, 2}, {2, 3}}, {{4, 5}, {6, 7}}}
	fmt.Printf("%v, len %d, cap %d\n", nums3, len(nums3), cap(nums3))


	fmt.Println("----------------------------------------End Array  ----------------------------------------")
}
