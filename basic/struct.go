package basic

import "fmt"
	// تعریف ساختار تو در تو
type user struct {
	name    string	
	age     int	
	salary  int	
	address address	
}	

type address struct {	
	city    string	
	country string	
}

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
	fmt.Println("                                    ***********             A            *********\n")
	emp3 := employee{"Sam", 31, 2000}
	 //دسترسی به فیلدهای ساختار 
    fmt.Printf("Current name is: %s\n", emp3.name)
	 // تغییر فیلد
    emp3.name = "John"
    fmt.Printf("New name is: %s\n", emp3.name)


	// تعریف فیلد ناشناس در ساختار
	type person struct {
		string	
		age    int	
		salary int	
	}
	
	fmt.Println("                                    ***********             B            *********\n")
    emp := person{string: "Sam", age: 31, salary: 2000}
    //Accessing a struct field
    n := emp.string
    fmt.Printf("Current name is: %s\n", n) //Current name is: Sam
    emp.string = "John"
    fmt.Printf("New name is: %s\n", emp.string)//New name is: John
	
	//زمانی که از فیلد های ناشناس استفاده می کنید، از هر دیتاتایپ فقط یکبار می توانید استفاده کنید
	/*type employee struct {
		string // name	
		int    // age	
		int    // salary	//error : int redeclared
	}*/
	fmt.Println("                                    ***********             C            *********\n")
	
    address := address{city: "London", country: "UK"}
    x := user{name: "Sam", age: 31, salary: 2000, address: address}
    fmt.Printf("City: %s\n", x.address.city)
    fmt.Printf("Country: %s\n", x.address.country)
	fmt.Println("                                    ***********             D            *********\n")
	var mobile Mobile = Mobile{}
	mobile.HasCamera = true
	mobile.Product.Name = "Iphone 11"
    fmt.Printf("Has Camera : %t\n", mobile.HasCamera)
    fmt.Printf("Name : %s\n", mobile.Product.Name)

	fmt.Println("----------------------------------------End Struct  ----------------------------------------")
}

type Product struct {
	Name  string	
	Price int	
}	
type Feature struct {
	HasCamera  bool
}	

type Mobile struct {	
	Feature //mobile.HasCamera
	Product  Product // but here mobile.Product.Name	
	Ram      int	
	SimCount int	
}
