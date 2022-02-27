package main

import (
	"fmt"
	"math/rand"
)

func main() {

	languageFundamentals()
}

func languageFundamentals() {

	booleanLogic()
	complexNumbers()
	stringLiterals()
	variablesDeclarations()
	shortVariableDeclaration()
	zeroValues()

}

func zeroValues() {
	fmt.Println("-----------------------------")
	fmt.Println("zeroValues")
	/*
			Zero Values
		When a variable is declared without an
		explicit value, it’s assigned to the
		zero value for its type:
	*/

	var i int
	var f float64
	var b bool
	var st string

	fmt.Printf("%v, %v, %v %v\n", i, f, b, st)
	fmt.Printf("%T, %T, %T %T\n\n", i, f, b, st)

	fmt.Printf("integer: %d\n", i)
	fmt.Printf("float: %f\n", f)
	fmt.Printf("boolean %t\n", b)
	fmt.Printf("string %s\n", st)

}

func shortVariableDeclaration() {
	fmt.Println("-----------------------------")
	fmt.Println("shortVariableDeclaration")

	percent := rand.Float64() * 100.0
	x, y := 33, 2

	fmt.Printf("%.4f, %d, %d\n", percent, x, y)

}

func variablesDeclarations() {
	fmt.Println("-----------------------------")
	fmt.Println("variablesDeclarations")

	var foo int = 42
	var boo, bar int = 42, 1302 // multiple variables
	var coo = 42                // type inference
	var b, f, s = true, 2.3, "foobar"
	var st string

	fmt.Printf("%v, %v, %v\n", foo, boo, bar)
	fmt.Printf("%v, %v, %v, %v, %v\n", coo, b, f, s, st)

}

func stringLiterals() {
	fmt.Println("-----------------------------")
	fmt.Println("stringLiterals")

	// The interpreted form
	string1 := "Hello\nworld!\n"
	fmt.Printf("string1 : %s\n", string1)

	// The raw form
	string2 := `Hello
world!`
	fmt.Printf("string2: %s\n", string2)

	if string1 == string2 {
		fmt.Println("The Strings are the Same")
	}

	if string1 != string2 {
		fmt.Println("The strings are different")
	}

}

func complexNumbers() {
	fmt.Println("-----------------------------")
	fmt.Println("complexNumbers")

	var x complex64 = 3.1415i
	fmt.Println(x) // "(0+3.1415i)"

}

func booleanLogic() {
	fmt.Println("-----------------------------")
	fmt.Println("Boolean Logic")

	and := true && false
	fmt.Println(and)

	or := true || false
	fmt.Println(or)

	not := !true
	fmt.Println(not)
}
