package ExpressionsAndCompositions

import (
	"fmt"
	"log"
	"sort"

	"github.com/eiannone/keyboard"
)

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// interface type

func RunBasicTypes() {
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Trevor"

	log.Println(myString)

	myString = "John"

	var myBool = true
	myBool = false

	log.Println(myBool)
}

// aggregate types (array, struct)

// Car is the type for Cars
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func RunAggregateTypes() {
	// declare an array of strings
	var myStrings [3]string

	// add values to the array; remember that arrays start counting at 0
	myStrings[0] = "dog"
	myStrings[1] = "fish"
	myStrings[2] = "cat"

	// print out the first element of the array
	fmt.Println("The first item in the array is", myStrings[0])

	// create a variable of type Car and populate it in one step
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
	fmt.Println()
}

// reference types (pointers, slices, maps, functions, channels)

func RunPointer() {

	x := 10

	myFirstPointer := &x

	fmt.Println("x is", x)

	*myFirstPointer = 15

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)
}


func changeValueOfPointer(num *int) {
	*num = 25
}

func RunSlice() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = deleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // first, copy the last element in the slice to index i
	a[len(a)-1] = ""   // next, erase the last element by setting to an empty string
	a = a[:len(a)-1]   // finally, truncate the slice by getting all elements except for the last one
	return a           // return the truncated slice
}

func RunMap() {
	// maps are reference types, so they are always passed by reference. You don't need a pointer.
	// maps are keys and values, and you must use the make function when creating
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// range through the map and print out key and value
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete a value from a map by key
	delete(intMap, "four")

	// check to see if an item exists in a map by key
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is in not map")
	}

	// update a value in the map
	intMap["two"] = 4
}

func RunFunction() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.says()
	cat.howManyLegs()
}

// Animal is a type, with three fields
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says is a receiver function, tied to the type Animal
// Because it is tied to the type, you can only call it
// when using a variable of type Animal
func (a *Animal) says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

// HowManyLegs is a function tied to the Animal type
func (a *Animal) howManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

// addTwoNumbers returns the sum of 2 ints
func addTwoNumbers(x, y int) int {
	return x + y
}

// naked return - rarely used, but possible, and only
// intended for short functions.
func add(x, y int) (sum int) {
	sum = x + y
	return
}

// variadic function - takes variable number of parameters
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}

var keyPressChan chan rune

func RunChannel() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

// listenForKeyPress is called using the "go" keyword, so it runs as a goroutine while the
// calling function (main) continues to execute
func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}





// interface type
type AnimalInterface interface {
	saysInterface() string
	howManyLegsInterface() int
}

// Dog is the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says is required by the Animal interface
func (d *Dog) saysInterface() string {
	return d.Sound
}

// HowManyLegs is required by the Animal interface
func (d *Dog) howManyLegsInterface() int {
	return d.NumberOfLegs
}

// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

// Says is required by the Animal interface
func (c *Cat) saysInterface() string {
	return c.Sound
}

// HowManyLegs is required by the Animal interface
func (c *Cat) howManyLegsInterface() int {
	return c.NumberOfLegs
}

func RunInterface() {
	// create a variable of type Dog
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	// Pass dog to riddle. Although dog is of type Dog, it satisifies the
	// interface requirements for Animal because it implements all of Animal's required functions.
	riddle(&dog)

	// Create  variable of type Cat
	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	// Pass cat to riddle. Although dog is of type Cat, it satisifies the
	// interface requirements for Animal because it implements all of Animal's required functions.
	riddle(&cat)
}

// Riddle takes a parameter of type Animal, but will accept both Dog and Cat, since both of those types
// satisfy the interface requirements for Animal, because they both have the correct functions.
func riddle(a AnimalInterface) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.saysInterface(), a.howManyLegsInterface())
	fmt.Println(riddle)
}

func RunExpressions() {
	age := 10           // age is an integer literal
	name := "Jack"      // name is an integer literal
	rightHanded := true // rightHanded is a boolean literal

	// in the next line, name, age, and righthanded are expressions, evaulated at run time
	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded)
	fmt.Println()

	// age + 10 is an expression, because it can be evaluated to a single value
	ageInTenYears := age + 10

	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears)
	fmt.Println()

	// age >= 13 is an expression, because it can be evaluated to a signle value
	isATeenager := age >= 13
	fmt.Println(name, "is a teenager:", isATeenager)

}

