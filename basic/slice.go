package basic

import (
	"fmt"
	"sort"
)

// آرایه دارای برخی از محدودیت‌ها علی الخصوص، اندازه ثابت می‌باشد اما در slice شما این محدودیت‌ها را نخواهید داشت و خیلی ساده می‌توانید المنت‌ها را افزایش، حذف و حتی کپی کنید.
func TestSlice() {
	fmt.Println("----------------------------------------Start Slice----------------------------------------")
	fmt.Println("                                    ***********             A            *********\n")

	slice := make([]int, 5) //5 is len
	fmt.Println(slice)
	fmt.Println(len(slice)) // Print 5
	fmt.Println(cap(slice)) // Print 5

	// مقدار ظرفیت نباید کمتر از مقدار اندازه باشد
	//test := make([]int, 5, 4)//5 is len & 4 is cap

	fmt.Println("                                    ***********             B            *********\n")
	slice1 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(len(slice1)) //Print 5
	fmt.Println(cap(slice1)) //Print 5

	fmt.Println("                                    ***********             C            *********\n")
	test := []int{99: 88}
	fmt.Println(len(test), cap(test))
	//len is 100 & cap is 100
	fmt.Println(test)
	//۹۹ تا صفر
	//0 0 0 .... 0 88

	fmt.Println("                                    ***********             D            *********\n")
	//تعریف یک slice خالی
	sliceOne := make([]int, 0)
	sliceTwo := []int{}
	fmt.Println(sliceOne == nil) // print false
	fmt.Println(len(sliceOne))   // print 0
	fmt.Println(cap(sliceOne))   // print 0
	fmt.Println(sliceTwo == nil) // print false
	fmt.Println(len(sliceTwo))   // print 0
	fmt.Println(cap(sliceTwo))   // print 0

	fmt.Println("                                    ***********             E            *********\n")
	slice2 := []int{10, 20, 30, 40}
	fmt.Println(slice2) //print [10 20 30 40]

	slice2[1] = 25
	fmt.Println(slice2) // print [10 25 30 40]

	fmt.Println("                                    ***********             F            *********\n")
	// ایجاد یک slice جدید بر اساس یک slice از پیش تعریف شده
	x := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(x)      // Print [10 20 30 40 50 60 70]
	fmt.Println(len(x)) // Print 7
	fmt.Println(cap(x)) // Print 7
	y := x[1:3]
	//Len: 3 - 1 = 2 Cap: 5 - 1 = 4
	fmt.Println(y)      //Print [20 30]
	fmt.Println(len(y)) //Print 2
	fmt.Println(cap(y)) //Print 6

	fmt.Println("                                    ***********             G            *********\n")
	slice4 := []int{10, 20, 30, 40, 50}
	newSlice := slice4[1:3]
	fmt.Println(len(newSlice)) // Print 2
	fmt.Println(cap(newSlice)) // Print 4
	newSlice = append(newSlice, 60)
	fmt.Println(len(newSlice)) // Print 3
	fmt.Println(cap(newSlice)) // Print 4

	fmt.Println("                                    ***********             G            *********\n")
	slice5 := []int{10, 20, 30, 40, 50}
	slice6 := []int{60, 70, 80}
	newSlice1 := []int{}
	newSlice2 := []int{}
	newSlice1 = append(newSlice1, slice5...)
	newSlice2 = append(newSlice2, slice5...)
	newSlice2 = append(newSlice2, slice6...)
	fmt.Println(newSlice1) // Print 10, 20, 30, 40, 50
	fmt.Println(newSlice2) // Print 10, 20, 30, 40, 50, 60, 70, 80

	fmt.Println("                                    ***********             H            *********\n")
	//حذف عنصر در ایندکس یک ولی خروجی نامنظم است
	slice7 := []int{10, 20, 30, 40, 50}
	slice7[1] = slice7[len(slice7)-1]
	slice7 = slice7[:len(slice7)-1]
	fmt.Println(slice7) //[10 50 30 40]

	fmt.Println("                                    ***********             I            *********\n")
	slice8 := []int{1, 2, 3, 4, 5}
	index := 2 // ایندکس المنتی که میخاییم حذفش کنیم
	//slice8[:index] تا قبل ایندکس ۲
	//slice8[index+1:] از ایندکس ۳ به بعد
	slice8 = append(slice8[:index], slice8[index+1:]...)
	fmt.Println(slice8) // خروجی: [1 2 4 5]

	fmt.Println("                                    ***********             J            *********\n")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)
	count := copy(dst, src)
	fmt.Println(count, dst)

	fmt.Println("                                    ***********             K            *********\n")
	src0 := []int{1, 2, 3, 4, 5}
	src1 := []int{10, 0, 0, 0, 0, 6, 7}
	count1 := copy(src1, src0)
	fmt.Println(count1, src1)

	fmt.Println("                                    ***********             L            *********\n")
	src3 := []int{1, 2, 3, 4, 5}
	src4 := []int{10, 0, 0, 0, 0, 6, 7}
	count2 := copy(src3, src4)
	fmt.Println(count2, src4)

	fmt.Println("                                    ***********             M            *********\n")

	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s) //1 2 3 4

	// فرق بین آرایه و slice
	// array_ := [3]int{10, 20, 30}
	// slice_ := []int{10, 20, 30}
	fmt.Println("                                    ***********             N            *********\n")
	var _slice []int32
	var _array [2]int32

	fmt.Println(_slice == nil) // print true
	fmt.Println(len(_slice))   // print 0
	fmt.Println(cap(_slice))   // print 0
	fmt.Println(len(_array))   // print 2
	fmt.Println(cap(_array))   // print 2
	fmt.Println(_array)        // print [0 0]

	fmt.Println("----------------------------------------End Slice  ----------------------------------------")
}
