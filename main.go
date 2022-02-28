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
	blankIdentifier()
	constantsVariables()
	containerTypesArraysSlicesMaps()
	pointerLogic()

	controlStructures()

}

func controlStructures() {
	fmt.Println("-----------------------------")
	fmt.Println("controlStructures")

	forStatement()

}

func forStatement() {
	fmt.Println("-----------------------------")
	fmt.Println("forStatement")

	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1
		fmt.Printf("sum - [i]  = %v - [%v]\n", sum, i)
	}

	fmt.Println("sum result = ", sum) // 10

	sum2, i2 := 0, 0

	for i2 < 10 {
		sum2 += i2
		i2++
	}

	fmt.Println(i2, sum2)
}

func pointerLogic() {
	fmt.Println("-----------------------------")
	fmt.Println("pointerLogic")
	/*
		Retrieving the address of a variable
		The address of a named variable can be retrieved by using
		the & operator. For example, the expression p := &a will
		obtain the address of a and assign it to p.

		Pointer types
		The variable p, which you can say “points to” a, has a
		type of *int, where the * indicates that it’s a pointer
		type that points to an int.

		Dereferencing a pointer
		To retrieve the value of the value a from p, you can
		dereference it using a * before the pointer variable
		name, allowing us to indirectly read or update a.
	*/

	a := 10

	var p *int = &a
	fmt.Printf("fmt.Println(p) = %v\n", p)
	fmt.Printf("fmt.Println(*p) = %v\n", *p)

	*p = 20
	fmt.Println("*p = 20")
	fmt.Printf("fmt.Println(a) = %v\n", a)

	var n *int
	var x, y int

	fmt.Println(n)        // "<nil>"
	fmt.Println(n == nil) // "true" (n is nil)

	fmt.Println(x == y)    // "true" (x and y are both zero)
	fmt.Println(&x == &x)  // "true" (*x is equal to itself)
	fmt.Println(&x == &y)  // "false" (different vars)
	fmt.Println(&x == nil) // "false" (*x is not nil)

}

func containerTypesArraysSlicesMaps() {
	fmt.Println("-----------------------------")
	fmt.Println("containerTypesArraysSlicesMaps")

	/*
		ArrayArray
		A fixed-length sequence of zero or more elements of a
		particular type.

		Slice
		An abstraction around an array that can be resized
		at runtime.

		Map
		An associative data structure that allows distinct
		keys to be arbitrarily paired with, or “mapped
		to,” values.
	*/

	containerArrays()
	containerSlice()
	containerMap()

}

func containerMap() {
	fmt.Println("-----------------------------")
	fmt.Println("containerMap")
	/*
		Go’s map data type references a hash table:
		an incredibly useful associative data structure
		that allows distinct keys to be arbitrarily
		“mapped” to values as key-value pairs. This
		data structure is common among today’s
		mainstream languages: if you’re coming to
		Go from one of these then you probably
		already use them, perhaps in the form of
		Python’s dict, Ruby’s Hash, or Java’s HashMap.

		Map types in Go are written map[K]V, where
		K and V are the types of its keys and values,
		respectively.
	*/

	// initialization #1
	freezing := make(map[string]float32)

	freezing["celsius"] = 0.0
	freezing["fahrenheit"] = 32.0
	freezing["kelvin"] = 273.2

	fmt.Println(freezing)

	fmt.Println(freezing["kelvin"])
	fmt.Println(len(freezing))

	delete(freezing, "kelvin")
	fmt.Println(len(freezing))

	// initialization #2
	freezing2 := map[string]float32{
		"celsius":    0.0,
		"fahrenheit": 32.0,
		"kelvin":     273.2,
	}
	fmt.Println(freezing2)

	noSuchKey := freezing["NoSuchKey"]
	fmt.Printf("no such key = %v", noSuchKey)

	newton, ok := freezing["newton"] // What about the Newton scale?
	fmt.Println(newton)              //  0 value
	fmt.Println(ok)                  // false = not found

}

func containerSlice() {
	fmt.Println("-----------------------------")
	fmt.Println("containerSlice")

	/*
			Slices

		Slices are a data type in Go that provide a powerful
		abstraction around a traditional array, such that
		working with slices looks and feels to the programmer
		very much like working with arrays. Like arrays,
		slices provide access to a sequence of elements of a
		particular type via the familiar bracket notation,
		indexed from 0 to N-1. However, where arrays are
		fixed-length, slices can be resized at runtime.
	*/

	// slice initialization
	n := make([]int, 3)

	fmt.Println(n)
	fmt.Println(len(n))

	n[0] = 8
	n[1] = 16
	n[2] = 32

	fmt.Println(n)

	// slice initialization #2
	m := []int{1}
	fmt.Println(m)

	/*
		Slices can be extended using the append built-in,
		which returns an extended slice containing one or
		more new values appended to the original one:
	*/
	m = append(m, 2)
	fmt.Println(m)

	m = append(m, 3)
	fmt.Println(m)

	m = append(m, 4, 5)
	fmt.Println(m)

	m = append(m, 6, 7, 8)
	fmt.Println(m)

	fmt.Println("--- Slice Operators ---")
	s := []int{0, 1, 2, 3, 4, 5, 6} // a slice literal
	fmt.Println(s)                  // "[0 1 2 3 4 5 6]"

	s1 := s[:4]
	fmt.Println(s1) //  [0 1 2 3]

	s2 := s[3:]
	fmt.Println(s2) //  [3 4 5 6]

	s[3] = 42
	fmt.Printf("Changed value s[3] = 42\n")
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("--- Strings as Slices ---")
	sValue := "foö"
	r := []rune(sValue)
	b := []byte(sValue)

	fmt.Printf("%7T %v\n", sValue, sValue)
	fmt.Printf("%7T %v\n", r, r)
	fmt.Printf("%7T %v\n", b, b)

}

func containerArrays() {
	fmt.Println("-----------------------------")
	fmt.Println("containerArrays")

	// array initialization
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])

	a[1] = 42
	fmt.Println(a)
	fmt.Println(a[1])

	i := a[1]
	fmt.Println(i)

	// array initialization #2
	b := [3]int{2, 4, 6}
	fmt.Println(b)

	c := [...]int{2, 4, 6}
	fmt.Println(c)

	/*
		As with all container types, the len built-in
		function can be used to discover the length of an array:

	*/

	fmt.Println(len(b))      // "3"
	fmt.Println(b[len(b)-1]) // "6"

}

func constantsVariables() {
	fmt.Println("-----------------------------")
	fmt.Println("constantsVariables")

	/*
		Constants are very similar to variables, using the
		const keyword to associate an identifier with some
		typed value. However, constants differ from variables
		in some important ways. First, and most obviously,
		attempting to modify a constant will generate an error
		at compile time. Second, constants must be assigned a
		value at declaration: they have no zero value.
	*/

	const language string = "Go"

	var favorite bool = true

	const text = "Does %s rule? %t!"
	var output = fmt.Sprintf(text, language, favorite)

	fmt.Println(output)
	fmt.Printf(text, language, favorite)

}

func blankIdentifier() {
	fmt.Println("-----------------------------")
	fmt.Println("blankIdentifier")
	/*
		The Blank Identifier
		The blank identifier, represented by the _ (underscore)
		operator, acts as an anonymous placeholder.
		It may be used like any other identifier in a declaration,
		except it doesn’t introduce a binding.

	*/

	str := "world"

	_, err := fmt.Printf("Hello %s\n", str)
	if err != nil {
		fmt.Println(err)
	}

	/*

		The blank identifier can also be used to import a package
		solely for its side effects:

		import _ "github.com/lib/pq"
		Packages imported in this way are loaded and initialized
		as normal, including triggering any of its init functions,
		but are otherwise ignored and need not be referenced or
		otherwise directly used.
	*/

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
