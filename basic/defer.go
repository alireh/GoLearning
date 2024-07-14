package basic

import (
	"fmt"

)

func TestDefer() {
	fmt.Println("----------------------------------------Start Defer----------------------------------------")
	fmt.Println("                                    ***********             A            *********\n")
	
	defer fmt.Println("0")//delay in execute command
	fmt.Println("1")
	fmt.Println("                                    ***********             B            *********\n")
	
    i := 0
    i = 2
    defer fmt.Println(i)

    i = 3
    defer fmt.Println(i)

    i = 4
    defer fmt.Println(i)
	fmt.Println("                                    ***********             C            *********\n")
     // شدن آن انجام می‌شود call کردیم در همان لحظه defer مقداردهی پارامترهای ورودی، برای تابعی که آن را
	j:=1
	defer fmt.Println(j)
	j++
	fmt.Println(j)
	fmt.Println("end")

	fmt.Println("----------------------------------------End Defer  ----------------------------------------")
}
