package basic

import (
	"errors"
	"fmt"

)

type inputError struct {
	message      string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

func validate(name, gender string) error {
	if name == "" {
		return &inputError{message: "Name is mandatory", missingField: "name"}
	}

	if gender == "" {
		return &inputError{message: "Gender is mandatory", missingField: "gender"}
	}
	return nil
}

func TestAdvancedError() {
	fmt.Println("----------------------------------------Start Advanced Error----------------------------------------")

	fmt.Println("                                    ***********             A            *********\n")
	err := validate("", "")

	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}
	fmt.Println("                                    ***********             B            *********\n")
	//wrapping error
    err = doSomething()
    if err != nil {
        fmt.Println(err) //first error is : errro 1
        fmt.Println(errors.Unwrap(err)) //errro 1
    }

	fmt.Println("----------------------------------------End Advanced Error  ----------------------------------------")
}


func doSomething() error {
    err1 := doFirstThing()
    if err1 != nil {
        return fmt.Errorf("first error is : %w", err1)
    }

    return nil
}

func doFirstThing() error {
    return errors.New("errro 1")
}
