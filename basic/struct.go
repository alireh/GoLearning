package basic

import fmt

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
	emp := employee{name: "Sam", age: 31, salary: 2000}

	emp := employee{
		name:   "Sam",	 
		age:    31,	 
		salary: 2000,	 
	 }
	 //انتصاب بر اساس ترتیب
	emp := employee{"Sam", 31, 2000}
	 //دسترسی به فیلدهای ساختار 
    fmt.Printf("Current name is: %s\n", emp.name)
	 // تغییر فیلد
    emp.name = "John"
    fmt.Printf("New name is: %s\n", emp.name)

	fmt.Println("----------------------------------------End Struct  ----------------------------------------")
}
