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
	// basic.TestForLoop()
	// basic.TestRange()
	// basic.TestLabeledLoop()
	// basic.TestContinueInLoop()
	// basic.TestSwitchStatement()
	// basic.TestVariable()
	// basicFunctions.TestFunction()
	basicFunctions.TestArray()
}
