package structs

import "math"

/*
	Perimeter takes a width and a height of a rectangle and returns its perimeter (perimeter = 2 * (width + height))
*/
func Perimeter (width float64, height float64) float64 {
	return 2 * (width + height)
}

/*
	Area takes a wifht and a height of a rectangle and returns its area (area = width * height )
*/
func Area (width, height float64) float64 {
	return width * height
}


// Structs
/*
	Rectangle - decalring our own type, rectangle struct 
	You can access the fields of a struct with the syntax of myStruct.field.
*/
type Rectangle struct {
	Width  float64
	Height float64
}	

/*
	Circle struct
*/
type Circle struct {
	Radius float64
}

/*
	Triangle struct
*/
type Triangle struct {
	Base  float64
	Height float64
}


/* 
	StructPerimeter takes Rectangle struct that has 2 fields Width and Height. It returns the rectangle's perimeter.
*/	

func StructPerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

/* 
	StructArea takes Rectangle struct that has 2 fields Width and Height. It returns the rectangle's area.
*/
func StructArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// Methods
/*
	Area method for Rectangle
	The syntax for declaring methods is almost the same as functions and that's because they're so similar. 
	The only difference is the syntax of the method receiver func (receiverName ReceiverType) MethodName(args).
	When your method is called on a variable of that type, you get your reference to its data via the receiverName variable. 
	In many other programming languages this is done implicitly and you access the receiver via this.
	It is a convention in Go to have the receiver variable be the first letter of the type.
	For example: r Rectangle
*/
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
/*
	Area method for Circle
*/
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

/*
	Area method for Triangle
*/
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Interfaces
/*
	Interface - decalring our own type as a Shape interface 
	You can define functions that can be used by different types (parametric polymorphism)
	https://en.wikipedia.org/wiki/Parametric_polymorphism
*/
type Shape interface {
	Area() float64
}
