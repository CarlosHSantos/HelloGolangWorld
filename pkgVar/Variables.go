package Variables

import (
	"fmt"
	"math"
)

func VariableFormats() {
	fmt.Println("Entrou no package Variables")

	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = 2
	fmt.Println("firstNumber -> ", firstNumber)

	// another way, declare type and name and assign value
	var secondNumber = 5
	fmt.Println("secondNumber -> ", secondNumber)

	// one step variable: declare name, assign value, and let Go figure out the type
	subtraction := 7
	fmt.Println("subtraction -> ", subtraction)

	var b, c int = 1, 2
    fmt.Println(b, c)
}
	
const s string = "constant"

func ConstantFormats() {
	fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}