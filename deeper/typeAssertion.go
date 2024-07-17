package deeper

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Define a rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define a circle type
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func TestTypeAssertion() {
	fmt.Println("----------------------------------------Start Type Assertion----------------------------------------")
	// Define a slice of shapes (interfaces)
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 1.5},
		Rectangle{Width: 2, Height: 2},
	}

	// Iterate through the shapes
	for _, shape := range shapes {
		// Perform a type assertion to access specific fields or methods
		if rect, ok := shape.(Rectangle); ok {
			fmt.Printf("Rectangle: Width=%.2f, Height=%.2f, Area=%.2f\n", rect.Width, rect.Height, rect.Area())
		} else if circle, ok := shape.(Circle); ok {
			fmt.Printf("Circle: Radius=%.2f, Area=%.2f\n", circle.Radius, circle.Area())
		} else {
			fmt.Println("Unknown shape")
		}
	}
	fmt.Println("----------------------------------------End Type Assertion  ----------------------------------------")
}
