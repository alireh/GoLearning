package basic

import (
	"fmt"

)

func TestContinueInLoop() {
	fmt.Println("----------------------------------------Start Continue In Loop----------------------------------------")
	for i := 1; i < 10; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
	fmt.Println("----------------------------------------End Continue In Loop  ----------------------------------------")
}
