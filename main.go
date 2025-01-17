// پکیج main به عنوان یک پکیج منحصر به فرد در نظر گرفته شده است
package main

//basicFunctions is alias
// import basicFunctions "learning/basic"
// import deeperFunctions "learning/deeper"
// import cliFunctions "learning/cli"
// import ormFunctions "learning/orm"
// import webFrameworkFunctions "learning/webFramework"
// import loggingFunctions "learning/logging"
// import restFunctions "learning/ApiClients/rest"
// import realtimeCommunicationFunctions "learning/realtimeCommunication"
// import microservicesToolsGoKitFunctions "learning/microservicesTools/goKit"
// import microservicesToolsFunctions "learning/microservicesTools"
import graphQLFunctions "learning/ApiClients/graphQL"

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
	// basicFunctions.TestForLoop();
	// basicFunctions.TestRange();
	// basicFunctions.TestLabeledLoop();
	// basicFunctions.TestContinueInLoop();
	// basicFunctions.TestSwitchStatement();
	// basicFunctions.TestVariable();
	// basicFunctions.TestFunction();
	// basicFunctions.TestArray();
	// basicFunctions.TestSlice();
	// basicFunctions.TestMap();
	// basicFunctions.TestStruct();
	// basicFunctions.TestStructPointer();
	// basicFunctions.TestError();
	// basicFunctions.TestAdvancedError();
	// basicFunctions.TestDefer();
	// basicFunctions.TestPanic();
	// basicFunctions.TestTypeAssertion();
	// basicFunctions.TestContext();
	// basicFunctions.TestChannel();
	// basicFunctions.TestSelect();
	// basicFunctions.TestBuffer();
	// basicFunctions.TestScheduler();
	// basicFunctions.TestMutex();
	// basicFunctions.TestGenerics();
	// basicFunctions.TestGoroutine();
	// basicFunctions.TestPointer();
	// deeperFunctions.TestMarshaling();
	// cliFunctions.TestCobra();
	// ormFunctions.TestGorm();
	// webFrameworkFunctions.TestGin();
	// webFrameworkFunctions.TestGinPostgres();
	// webFrameworkFunctions.TestGofiber();
	// webFrameworkFunctions.TestGorilla();
	// loggingFunctions.TestZerolog();
	// loggingFunctions.TestZap();
	// loggingFunctions.TestLog();
	// restFunctions.TestHeimdall();
	// realtimeCommunicationFunctions.TestMelody();
	// microservicesToolsFunctions.TestRpcx();
	// microservicesToolsGoKitFunctions.TestGoKit()
	// microservicesToolsFunctions.TestMicro();
	graphQLFunctions.TestGraphqlGo();
}
