package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Traingle struct {
	height float64
	base   float64
}

// methods
// both functions and are almost same but only difference in syntax method syntax will be like
// func (recieverName RecieverType) methodName(args){}
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * c.radius
}

func (t Traingle) Area() float64 {
	return t.base * t.height * 0.5
}

// func Perimeter(rectangle Rectangle) float64 {
// 	perimeter := 2 * (rectangle.width + rectangle.height)
// 	return perimeter
// }

// func Area(rectangle Rectangle) float64 {
// 	area := rectangle.width * rectangle.height
// 	return area
// }

// func main() {
// 	rect := Rectangle{
// 		width:  2.0,
// 		height: 4.0,
// 	}
// 	primeter := Perimeter(rect)
// 	fmt.Println(primeter)

// 	area := Area(rect)
// 	fmt.Println(area)
// }
