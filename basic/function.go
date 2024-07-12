package basic

import (
	"fmt"
	"time"

)

func plus(a int, b int) int {
	return a + b
}

// or func plus(a, b int)
// func name(a string, b,c int) is equal to func name(a string, b int,c int)
//

/* or
plus := func (a int, b int) int {
	return a + b
}
*/

// تابع چند بازگشتی
func vals() (int, int) {
	return 13, 17
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
} //return x, y even tought return is empty

// تابع با ورودیهای متنوع
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// پاس دادن تابع به تابع
func add10AndSum(num1 int, num2 int, sum func(n1, n2 int) int) {
	result := sum(num1+10, num2+10)
	fmt.Println("Sum by adding 10 is:", result)
}

func TestFunction() {
	fmt.Println("----------------------------------------Start Function----------------------------------------")
	fmt.Println("                                    ***********             A            *********\n")
	fmt.Println(plus(3, 4))

	fmt.Println("                                    ***********             B            *********\n")

	_, c := vals()
	fmt.Println(c)

	d, _ := vals()
	fmt.Println(d)

	vals()

	e, f := vals()
	fmt.Println(e, f)

	fmt.Println("                                    ***********             C            *********\n")
	sum(1, 2)    //3
	sum(1, 2, 3) //6
	nums := []int{1, 2, 3, 4}
	sum(nums...) //10

	fmt.Println("                                    ***********             D            *********\n")
	add10AndSum(5, 3, func(n1, n2 int) int {
		sum := n1 + n2
		return sum
	})

	fmt.Println("                                    ***********             E            *********\n")

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 10)
	//output : 10 10 10 ......

	fmt.Println("                                    ***********             F            *********\n")

	for i := 0; i < 10; i++ {

		go func(num int) {

			fmt.Println(num)

		}(i)

	}
	time.Sleep(time.Second * 1)
	fmt.Println("----------------------------------------End Function  ----------------------------------------")
}
