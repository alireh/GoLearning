package main

import (
	"learning/utils"
)

func f() bool {
	return true
}

func main() {
	switch f() {

	case true:

		println(1)

	case false:

		println(0)

	}
	//utils.TestForLoop()
	//utils.TestRange()
	// utils.TestLabeledLoop()
	// utils.TestContinueInLoop()
	utils.TestSwitchStatement()

}
