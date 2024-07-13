package basic

import "fmt"

func TestStruct() {
	fmt.Println("----------------------------------------Start Struct----------------------------------------")
	type employee struct {
		name   string	
		age    int	
		salary int	
	}
	
	 //ساختار بصورت خالی
	type sample struct {}

	 //ایجاد متغیر ساختار و مقدار دهی فیلدها 
	// emp1 := employee{name: "Sam", age: 31, salary: 2000}

	// emp2 := employee{
	// 	name:   "Sam",	 
	// 	age:    31,	 
	// 	salary: 2000,	 
	// }
	 //انتصاب بر اساس ترتیب
	emp3 := employee{"Sam", 31, 2000}
	 //دسترسی به فیلدهای ساختار 
    fmt.Printf("Current name is: %s\n", emp3.name)
	 // تغییر فیلد
    emp3.name = "John"
    fmt.Printf("New name is: %s\n", emp3.name)

	fmt.Println("----------------------------------------End Struct  ----------------------------------------")
}
