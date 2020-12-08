package calcit

import "github.com/motiso/calcit/strings"

func GetStringLength(str string) int {
	return strings.StringLength(str)
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Mul(x, y int) int {
	return x * y
}

func Div(x, y int) int {
	return x / y
}

func Square(x int) int {
	return x * x
}
