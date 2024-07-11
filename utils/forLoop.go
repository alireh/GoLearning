package utils

/**هر پکیج باید حتما در فولدر جداگانه باشد */
import (
	"fmt"
)

func TestForLoop() {
	fmt.Println("----------------------------------------Start Loop----------------------------------------")
	for i := 1; i <= 3; i++ {
		message := "loop A " + fmt.Sprintf("%d", i)
		println(message)
	}

	fmt.Println("                                    ***********             A            *********\n")

	j := 1

	for j <= 3 {

		fmt.Println("loop B " + fmt.Sprintf("%d", j))
		j++

	}

	fmt.Println("                                    ***********             B            *********\n")

	sum := 0

	//این حلقه بی نهایت است که به
	//break
	//جلوی ادامه آن گرفتهخ شده است
	for {

		sum++
		if sum == 4 {
			break
		}
	}
	fmt.Println("Sum is " + fmt.Sprintf("%d", sum))
	fmt.Println("----------------------------------------End Loop----------------------------------------")

}
