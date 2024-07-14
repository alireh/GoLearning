package basic

import (
	"fmt"

)

func TestPanic() {
	fmt.Println("----------------------------------------Start Panic----------------------------------------")
	
	a := []string{"a", "b"}
	 // در زمان اجرا panic خطای
	println(a[2])
	
	 //از قبل تعیین شده panic خطای
	//checkAndPrint(a, 2)
	
	fmt.Println("----------------------------------------End Panic  ----------------------------------------")
}

func checkAndPrint(a []string, index int) {

	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}
