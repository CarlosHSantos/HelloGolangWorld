package Operators

import (
	"fmt"
	"math"
	"errors"
)

func RunBasicOperators() {
	answer := 7 + 3 * 4
	fmt.Println("Answer:", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer is now", answer)
}

func RunPrimaryOperators() {
	// multiplication
	var radius = 12.0

	area := math.Pi * radius * radius
	fmt.Println("area", area)

	// integer division
	half := 1 / 2
	fmt.Println("Half", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("half float", halfFloat)

	// squaring, and raising to power
	badThreeSquared := 3 ^ 2 // this is the bitwise XOR operator
	fmt.Println("Three squared is not", badThreeSquared)
	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("Three squared is", goodThreeSquared)

	// modulus operator
	remainder := 50 % 3
	fmt.Println("remaider is", remainder)

	// unary operators
	x := 3
	// can't do this
	//y := x++
	// can do this
	var y = x
	y++
	fmt.Println("x is", x, "and y is", y)
}

func RunShortCircuitEvaluation() {
	a := 12
	b := 6

	// if b != 0 {
	// 	c := divideTwoNumbers(a, b)

	// 	if c == 2 {
	// 		fmt.Println("We found a two")
	// 	}
	// }

	// if b != 0 && divideTwoNumbers(a, b) == 2 {
	// 	fmt.Println("Found 2")
	// }

	// if divideTwoNumbers(a, b) == 2 && b != 0 {
	// 	fmt.Println("Found 2")
	// }

	c, err := divideTwoNumbers(a, b)
	if err != nil {
		fmt.Print(err)
	} else {
		if c == 2 {
			fmt.Println("We found 2")
		}
	}

}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func RunAssigmentOperators() {
	x := 12
	x--

	fmt.Println(x)

	y := 10
	y /= 2
	fmt.Println(y)
}