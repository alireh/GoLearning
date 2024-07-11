package utils

import (
	"fmt"
)

func TestLabeledLoop() {
	fmt.Println("----------------------------------------Start LabeledLoop----------------------------------------")
	letters := []string{"a", "b", "c"}

	fmt.Println("***************************************1")
	for j := 1; j < 3; j++ {

		fmt.Println("" + fmt.Sprintf("%d", j))

		for i := 2; i < 5; i++ {
			fmt.Println("out ----" + fmt.Sprintf("%d", i))

		first:
			for _, l := range letters {

				if l == "c" {
					break first
				}
				fmt.Println("in ---------" + fmt.Sprintf("%s", l))

			}
		}
	}

	fmt.Println("***************************************2")
	for j := 1; j < 3; j++ {

		fmt.Println("" + fmt.Sprintf("%d", j))

	second:
		for i := 2; i < 5; i++ {

			fmt.Println("out ----" + fmt.Sprintf("%d", i))

			for _, l := range letters {

				if l == "c" {

					// break the loop with second lable

					break second

				}
				fmt.Println("in ---------" + fmt.Sprintf("%s", l))

			}

		}

	}
	fmt.Println("----------------------------------------End LabeledLoop  ----------------------------------------")
}
