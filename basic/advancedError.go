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

	//ایجاد خطا با محتوای دلخواه
	sampleErr := errors.New("error occured")

	fmt.Println(sampleErr)

	msg := "database connection issue"
	sampleErr1 := fmt.Errorf("Err is: %s", msg)

	fmt.Println(sampleErr1)
	fmt.Println("                                    ***********             B            *********\n")

	err := validate("", "")

	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}

	fmt.Println("----------------------------------------End Advanced Error  ----------------------------------------")
}
