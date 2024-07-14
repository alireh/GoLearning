// پکیج main به عنوان یک پکیج منحصر به فرد در نظر گرفته شده است
package main

//basicFunctions is alias
import basicFunctions "learning/basic"

//اگر بخواهیم کدی قبل از main اجرا شود
func init() {
	println("init")
}

//به طور پیش فرض این تابع فراخوانی میشود
//go run main.go : دستور اجرای برنامه
//go build main.go : دستور ساخت برنامه
//go build -o myExecutable main.go : دستور ساخت برنامه
// go fmt : دستور مرتب سازی کد

//{ باید بعد از
//main
//بیاید نه در خط بعدی
func main() {
	/*TestForLoop باید با حروف بزرگ شروع شود تا پابلیک باشد*/
	// basicFunctions.TestForLoop()
	// basicFunctions.TestRange()
	// basicFunctions.TestLabeledLoop()
	// basicFunctions.TestContinueInLoop()
	// basicFunctions.TestSwitchStatement()
	// basicFunctions.TestVariable()
	// basicFunctions.TestFunction()
	// basicFunctions.TestArray()
	// basicFunctions.TestSlice()
	// basicFunctions.TestMap() 
	// basicFunctions.TestStruct() 
	// basicFunctions.TestStructPointer() 
	// basicFunctions.TestError();
	// basicFunctions.TestAdvancedError();
	// basicFunctions.TestDefer()
	basicFunctions.TestPanic();
}
