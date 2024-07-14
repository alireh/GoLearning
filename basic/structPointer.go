package basic

import "fmt"
import "encoding/json"
import "log"

type employee struct {
	name   string	
	age    int	
	salary int	
}
type person struct {
	Name   string	
	Age    int	
	salary int	
}

func TestStructPointer() {
	fmt.Println("----------------------------------------Start Struct Pointer----------------------------------------")
	emp := employee{name: "Sam", age: 31, salary: 2000}

	empP := &emp
	fmt.Println(empP)//&{Sam 31 2000}

	 //تعریف مستقیم اشاره گر به ساختار
	// empP := &employee{name: "Sam", age: 31, salary: 2000}

	empP1 := new(employee)
	fmt.Println(empP1)
	 // آدرس خانه حافظه ساختار
	fmt.Printf("%p\n", empP1) // 0xc000064420
	 //مقدار کلی فیلدها
	//%v - مقدار value هر کدام از فیلدهای ساختار را چاپ می‌کند.
	fmt.Printf("%v\n", *empP) //{name:Sam age:31 salary:2000}
	 //%+v - مقدار هرکدام از فیلدها به همراه اسم فیلد key-value را چاپ می‌کند
	fmt.Printf("%+v\n", *empP) //{name:Sam age:31 salary:2000}


	// چاپ ساختار با استفاده از پکیج JSON
	prs := person{Name: "Sam", Age: 31, salary: 12}

    prsJSON, err := json.Marshal(prs)// ۲ خروجی از نوع بایت و خطا

    if err != nil {
        log.Fatalf(err.Error())
    }
	 // بود private چاپ نشد چون //salary
    fmt.Printf("Marshal funnction output %s\n", string(prsJSON))


    //MarshalIndent
    prsJSON, err = json.MarshalIndent(prs, "", "  ")//در این تابع ما ۳ تا ورودی به تابع می‌فرستیم, به ترتیب ساختار، پیشوند و indent
    if err != nil {
        log.Fatalf(err.Error())
    }
	 // بود private چاپ نشد چون //salary
    fmt.Printf("MarshalIndent funnction output %s\n", string(prsJSON))


	//اگر اسم فیلدها را با کوچک شروع کنیم در خروجی مارشال چاپ نمیشود از طرفی به صورت بزرگ چاپ میشود برای رفع این مشکل از تگ استفاده میکنیم
	type user struct {
		Name   string `json:"name"`	
		Age    int    `json:"age"`	
	    Email    string    `json:"-"`// //برای اینکه درjson نیاید //
	    Phone    string    `json:"age,omitempty"`//در صورتیکه مقدار نداشته باشد در خروجی نمی‌آید
		Salary int    `json:"salary"`	
	}

	user1 := user{Name: "Sam", Age: 31, Salary: 2000}
	//Converting to jsonn
	userJSON, err := json.MarshalIndent(user1, "", "   ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("----user----")
	fmt.Println(string(userJSON))

	fmt.Println("----------------------------------------End Struct Pointer  ----------------------------------------")
}
