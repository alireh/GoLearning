package basic

import (	
    "errors"
	"fmt"
	"os"

)

func TestError() {
	fmt.Println("----------------------------------------Start Error----------------------------------------")
	fmt.Println("                                    ***********             A            *********\n")
	file1, err := os.Open("main.go") //os.open return two value

	//for ignore handing error
	filex, _ := os.Open("go.mod")
	fmt.Println(filex.Name() + " opened succesfully")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file1.Name() + " opened succesfully")
	}
	fmt.Println("                                    ***********             B            *********\n")
	file2, err := os.Open("not-existed-file.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file2.Name() + " opened succesfully")
	}
	fmt.Println("                                    ***********             C            *********\n")
	 // ایجاد یک خطا	 
	sampleErr1 := errors.New("error occured")
	fmt.Println(sampleErr1)

 	// fmt	 ایجاد یک خطا با //	
	msg := "database connection issue"
	sampleErr2 := fmt.Errorf("Err is: %s", msg) 
	fmt.Println(sampleErr2)
	fmt.Println("----------------------------------------End Error  ----------------------------------------")
}
