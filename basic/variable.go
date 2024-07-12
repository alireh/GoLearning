package basic

import (
	"fmt"
)

// ایجاد auto increment بدون iota :
const (
	d1 = 0
	d2 = 1
	d3 = 2
)

// ایجاد auto increment با iota :
const (
	d4 = iota // 0
	d5        // 1
	d6        // 2
)

// شروع از ۱
const (
	_  = iota
	d7 // 1
	d8 // 2
)

// enum
const (
	small int8 = iota
	medium
	large
	extraLarge
)

// با رسیدن به اولین انتصاب تا iota بعدی روند ثابت میماند
const (
	A = iota  //0
	B         //1
	C         //2
	D = B + C //3
	E         //3
	F         //3
	G = iota  //6
	H         //7
	I = H     //7
	J         //7
	K         //7
)

// نام مستعار برای متغیر
type decimal float32
type str string

const cv string = "const variable"

func TestVariable() {
	fmt.Println("----------------------------------------Start Variable----------------------------------------")

	var var1 decimal = 9.0
	fmt.Println(var1)

	var s1 str = "test"
	//auto dtect type
	s2 := "test"
	fmt.Println(s1)
	fmt.Println(s2)

	a, b, c := "hello", 1, 1.5
	var d, e, f = "world", 13, 24
	fmt.Println(a, b, c, d, e, f)

	// Integers
	//Hex
	var hexVar = 0xC
	println(hexVar)
	//Octal
	var OctalVar = 06 //or 0o6 or 0O6
	println(OctalVar)
	//Binary
	var binaryVar = 0b1011 //or 0B1011
	println(binaryVar)
	//Decimal
	var decimalVar = 144
	println(decimalVar)

	println(15 == 017) // true
	println(15 == 0xF) // true

	fmt.Println("----------------------------------------End Variable  ----------------------------------------")
}
