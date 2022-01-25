package Strings

import (
	"fmt"
	"strings"
)

func RunStringsPackage() {
	courses := []string {
		"Learn Go",
		"Learn Java",
		"Learn Python",
		"Learn C",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
		}
	}

	// 						1         2         3         4
	// 			  0123456789012345678901234567890123456789012345
	newString := "Go is a great programming language. Go for it!"

	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Fish"))
	fmt.Println(strings.Index(newString, "Python"))
	fmt.Println(strings.LastIndex(newString, "Go"))
}

func RunStringsManipulation() {
	myString := "This is a clear EXAMPLE of why we search in one case only."

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Did not find it!")
	}

	// other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))
}
