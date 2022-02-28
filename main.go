package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
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
	functions()
	deferFunctions()
	functionPointerArguements()
	functionsPointerRefTypes()
	variadicFunctions()
	anonymousFunctions()
	closureFunctions()

	structs()
	methods()
	interfaces()
	typeAssertions()

	// embedding
	interfaceEmbedding()
	structEmbedding()

	// Concurrency
	goRoutines()
	channels()
	channelBuffering()
	channelClosing()
	channelLooping()

}

func channelLooping() {
	fmt.Println("-----------------------------")
	fmt.Println("channelLooping")

	ch := make(chan string, 10)
	ch <- "Maxine"
	ch <- "Angelo"
	ch <- "Alexis"
	ch <- "Kristin"

	close(ch) // close otherwise it will deadlock
	/*
		Had the channel not been closed, the loop would
		stop and wait for the next value to be sent
		along the channel, potentially indefinitely.
	*/

	for s := range ch {
		fmt.Println(s)
	}
}

func channelClosing() {
	fmt.Println("-----------------------------")
	fmt.Println("channelClosing")

	ch := make(chan string, 10)
	ch <- "Maxine"
	ch <- "Angelo"
	ch <- "Alexis"
	ch <- "Kristin"

	close(ch)

	msg, ok := <-ch
	fmt.Printf("%q, %v\n", msg, ok)

	msg, ok = <-ch
	fmt.Printf("%q, %v\n", msg, ok)

	msg, ok = <-ch
	fmt.Printf("%q, %v\n", msg, ok)

	msg, ok = <-ch
	fmt.Printf("%q, %v\n", msg, ok)

	msg, ok = <-ch
	fmt.Printf("%q, %v\n", msg, ok)

}

func channelBuffering() {
	fmt.Println("-----------------------------")
	fmt.Println("channelBuffering")
	/*
			Channel Buffering

		Go channels may be buffered, in which case they
		contain an internal value queue with a fixed
		capacity that’s specified when the buffer is
		initialized. Sends to a buffered channel only
		block when the buffer is full; receives from
		a channel only block when the buffer is empty.
		Any other time, send and receive operations
		write to or read from the buffer, respectively,
		and exit immediately.

		A buffered channel can be created by providing
		a second argument to the make function to
		indicate its capacity:

	*/

	ch := make(chan string, 3)

	ch <- "foo" // Two non-blocking sends
	ch <- "bar"
	ch <- "Tony"

	fmt.Println(<-ch) // Two non-blocking receives
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//	fmt.Println(<-ch) // will block
}

func channels() {
	fmt.Println("-----------------------------")
	fmt.Println("channels")

	fmt.Println("fmt.Println(val2)")

	/*
		Channels may be created using the make function.
		Each channel can transmit values of a specific
		type, called its element type. Channel types
		are written using the chan keyword followed by
		their element type. The following example
		declares and allocates an int channel:
	*/

	// create a channel - element type = int
	// var ch chan int = make(chan int)

	/*
		The two primary operations supported by channels
		are send and receive, both of which use
		the <- operator, where the arrow indicates the
		direction of the data flow as demonstrated
		in the following:
	*/

	//val := 10
	//ch <- val // Sending on a channel
	//
	//fmt.Println("ddd")
	//val2 := <-ch // recieving on a channel and assigning to val2
	//fmt.Println(val2)

	//fmt.Println(val2)
	//
	//<-ch // recieving on a channel and disvarding the result

	//fmt.Println("fmt.Println(val2)")

	/*
			Channel Blocking

		By default, a channel is unbuffered. Unbuffered channels
		have a very useful property: sends on them block until
		another goroutine receives on the channel, and receives
		block until another goroutine sends on the channel.
		This behavior can be exploited to synchronize two
		goroutines, as demonstrated in the following:
	*/

	ch2 := make(chan string) // Allocate a string channel

	go func() {
		message := <-ch2     // Blocking receive; assigns to message
		fmt.Println(message) // prints "ping"
		ch2 <- "pong"
	}()

	ch2 <- "ping"
	fmt.Println(<-ch2) // print "pong"

}

func foo() {
	fmt.Println("This is the foo function")

}
func goRoutines() {
	fmt.Println("-----------------------------")
	fmt.Println("goRoutines")

	foo()    //  call foo, ait for it to complete
	go foo() // Spawn a new goroutine that calls foo() concurrently
}

func Log(w io.Writer, message string) {
	go func() {
		fmt.Fprint(w, message)
	}() // Don't forget the trailing parentheses!
}

//type ReadWriterStruct struct {
//	*Reader
//	*Writer
//}

func structEmbedding() {
	fmt.Println("-----------------------------")
	fmt.Println("structEmbedding")

	fmt.Println("Error")

}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Embedded interfaces
type ReadWriter interface {
	Reader
	Writer
}

func interfaceEmbedding() {
	fmt.Println("-----------------------------")
	fmt.Println("interfaceEmbedding")

	st := `type ReadWriter interface {
	Reader
	Writer
}`

	fmt.Println(st)
}

func typeAssertions() {
	fmt.Println("-----------------------------")
	fmt.Println("typeAssertions")

	var s Shape
	s = Circle{}
	c := s.(Circle) // Assert that s is a Circle
	fmt.Printf("%T\n", c)

}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("%T's area is %0.2f\n", s, s.Area())
}

func interfaces() {
	fmt.Println("-----------------------------")
	fmt.Println("interfaces")
	/*
		An interface can thus be viewed as a contract that
		a type may satisfy, opening the door to powerful
		abstraction techniques.
	*/

	r := Rectangle{5, 10}
	PrintArea(r)

	c := Circle{5}
	PrintArea(c)
}

func methods() {
	fmt.Println("-----------------------------")
	fmt.Println("methods")
	/*
			Methods
		In Go, methods are functions that are attached to types,
		including but not limited to structs. The declaration
		syntax for a method is very similar to that of a function,
		except that it includes an extra receiver argument before
		the function name that specifies the type that the method
		is attached to. When the method is called, the instance
		is accessible by the name specified in the receiver.

	*/

	vert := Vertex{3, 4}
	fmt.Printf("fmt.Printf = %v\n", vert)

	vert.Square()
	fmt.Printf("fmt.Printf = %v\n", vert)

	/*
		Receivers are type specific: methods attached
		to a pointer type can only be called on a pointer
		to that type.
	*/

	mm := MyMap{"A": 1, "B": 2}

	fmt.Println(mm)                             // "map[A:1 B:2]"
	fmt.Printf("mm[A] = %v\n", mm["A"])         // "1"
	fmt.Printf("mm.Length() = %d", mm.Length()) // "2"

}

func (v *Vertex) Square() {
	v.x *= v.x
	v.y *= v.y
}

type Vertex struct {
	x, y float64
}

type MyMap map[string]int

func (m MyMap) Length() int {
	return len(m)
}

func structs() {
	fmt.Println("-----------------------------")
	fmt.Println("structs")

	var v Vertex // Structs are never nil
	fmt.Printf("fmt.Printf = %v\n", v)

	v = Vertex{}
	fmt.Printf("fmt.Printf = %v\n", v)

	v = Vertex{1.0, 2.0}
	fmt.Printf("fmt.Printf = %v\n", v)

	v = Vertex{y: 2.5}
	fmt.Printf("fmt.Printf = %v\n", v)

	v = Vertex{x: 1.0, y: 3.0}
	fmt.Printf("fmt.Printf = %v\n", v)

	v.x *= 1.5
	v.y *= 2.5
	fmt.Printf("fmt.Printf = %v\n", v)

	var ver *Vertex = &Vertex{1, 3}
	fmt.Println(ver) // pointer adress of Vertex{1,3}

	ver.x, ver.y = ver.y, ver.x
	fmt.Printf("fmt.Printf = %v\n", ver)

	// fmt.Println(v)

}

func incrementer() func() int {
	i := 0

	return func() int { // Return an anonymous function
		i++ // "Closes over" parent function's i
		return i
	}
}

func closureFunctions() {
	fmt.Println("-----------------------------")
	fmt.Println("closureFunctions")

	/*
		A closure is a nested function that has access to
		the variables of its parent function, even after
		the parent has executed.
	*/
	increment := incrementer()
	fmt.Println("increment := incrementer()")
	fmt.Printf("fmt.Println(increment()) = %d\n", increment())
	fmt.Printf("fmt.Println(increment()) = %d\n", increment())
	fmt.Printf("fmt.Println(increment()) = %d\n\n", increment())

	/*
		fmt.Println(increment()) // 1
		fmt.Println(increment()) // 2
		fmt.Println(increment()) // 3
	*/

	newIncrement := incrementer()
	fmt.Println("newIncrement := incrementer()")
	fmt.Printf("fmt.Println(newIncrement()) = %d\n", newIncrement())

	// fmt.Println(newIncrement()) // 1
}

func anonSum(x, y int) int     { return x + y }
func anonProduct(x, y int) int { return x * y }

func anonymousFunctions() {
	fmt.Println("-----------------------------")
	fmt.Println("anonymousFunctions")

	var f func(int, int) int

	fmt.Println("f = anonSum")
	fmt.Print("fmt.Println(f(3,5)) = ")

	f = anonSum
	fmt.Println(f(3, 5)) // 8

	fmt.Println("f = anonProduct")
	fmt.Print("fmt.Println(f(3, 5)) = ")

	f = anonProduct
	fmt.Println(f(3, 5)) // 15
}

func variadicFunctions() {
	fmt.Println("-----------------------------")
	fmt.Println("variadicFunctions")

	/*A variadic function is one that
	may be called with zero or more trailing
	arguments. The most familiar example is
	the members of the fmt.Printf family
	of functions, which accept a single
	format specifier string and an arbitrary
	number of additional arguments.

	*/

	const name, age = "kim", 22

	fmt.Printf("%s is %d years old.\n", name, age)

	// variadic functions #2
	fmt.Println("product(2, 2, 2) = ", product(2, 2, 2)) // 8

	m := []int{3, 3, 3} // a slice
	fmt.Println("m := []int{3, 3, 3} // a slice")
	fmt.Printf("fmt.Println(product(m...)) = ")
	fmt.Println(product(m...))

}

func product(factors ...int) int {
	p := 1

	for _, n := range factors {
		p *= n
	}

	return p
}

func functionsPointerRefTypes() {
	fmt.Println("-----------------------------")
	fmt.Println("functionsPointerRefTypes")

	/*
		slices, maps, functions, and channels. Changes made
		to such reference types in a function can affect
		the caller, without needing to explicitly
		dereference them:
	*/

	m := map[string]int{
		"a": 0,
		"b": 1,
	}

	fmt.Println(m)

	update(m)

	fmt.Println(m)

}

func update(m map[string]int) {

	m["c"] = 2
}

func functionPointerArguements() {
	fmt.Println("-----------------------------")
	fmt.Println("functionPointerArguements")

	x := 5

	zerobyValue(x)
	fmt.Println(x)

	zeroByReference(&x)
	fmt.Println(x)
}

func zeroByReference(i *int) {
	*i = 0
}

func zerobyValue(x int) {
	x = 0
}

func deferFunctions() {
	fmt.Println("-----------------------------")
	fmt.Println("deferFunctions")

	fmt.Println("")
	defer fmt.Println("fn: world")
	defer fmt.Println("fn: cruel")
	defer fmt.Println("fn: goodbye")
}

func add(x, y int) int {
	return x + y

}

/*
Functions multiple return types
*/
func swap(a, b string) (string, string) {
	return b, a
}

func functions() {
	fmt.Println("-----------------------------")
	fmt.Println("functions")

	sum := add(10, 5)
	fmt.Println(sum)

	a, b := "foo", "bar"

	c, d := swap(a, b)
	fmt.Printf("c, d = %s, %s\n", c, d)

	// recursion
	fmt.Println("recursion factorial(11)")
	fmt.Println(factorial(11))

	// defer function
	defer fmt.Println("cruel world")
	fmt.Printf("goodbye ")

	// second example of defer
	file, err := os.Create("/tmp/foo.txt")
	defer closeFile(file)

	_, err = fmt.Fprint(file, "Your mother was a hampster #2\n")
	if err != nil {
		return
	}

	fmt.Println("File written to successfully")

	// hostname
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname:", name)

}

func closeFile(f *os.File) {

	if err := f.Close(); err != nil {
		fmt.Println("Error closing file:", err.Error())
		return
	}
	fmt.Println("File closed successfully")
}

// recursion
func factorial(n int) int {

	if n < 1 {
		return 1
	}
	return n * factorial(n-1)
}

func controlStructures() {
	fmt.Println("-----------------------------")
	fmt.Println("controlStructures")

	forStatement()
	loopingOverArrays()
	loopingOverMaps()

	ifStatement()
	switchStatement()

	errorHandling()

}

type NestedError struct {
	Message string
	Err     error
}

func (e *NestedError) Error() string {
	return fmt.Sprintf("%s\n containse: %s\n",
		e.Message, e.Err.Error())
}

func errorHandling() error {

	//file, err := os.Open("somefile.ext")
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	//fmt.Println(file)

	e1 := errors.New("error 42")
	e2 := fmt.Errorf("error %d", 42)

	fmt.Println(e1)
	fmt.Println(e2)

	return e1
}

func switchStatement() {
	fmt.Println("-----------------------------")
	fmt.Println("switchStatement")
	i := 7

	switch i % 3 {
	case 0:
		fmt.Println("Zero")
		fallthrough
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("huh?")
	}

	hour := time.Now().Hour()
	fmt.Println("hour = %v\n", hour)

	tm := time.Now()
	hr := time.Now().Hour()
	fmt.Printf("tm = %v\n", tm)
	fmt.Printf("hr = %v\n", hr)

	switch {
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}

	switch hour := time.Now().Hour(); { // Empty expression means "true"
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}
}

func ifStatement() {
	fmt.Println("-----------------------------")
	fmt.Println("ifStatement")

	if 7%2 == 0 {
		fmt.Println("7 is even")
	}

	if 7%2 != 0 {
		fmt.Println("7 is odd")
	}

	fmt.Println("Looking fo a file")
	if _, err := os.Open("foo.ext"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is fine")
	}

	fmt.Println("Looking fo a file #2")
	if _, err := os.Open("/Users/mcclayac/SSH/AnthonyMcClayVAKeyPair.pem"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("/Users/mcclayac/SSH/AnthonyMcClayVAKeyPair.pem is found")
	}

	fmt.Println("Looking fo a file #3")
	if _, err := os.Open("main.go"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("main.go is found")
	}
}

func loopingOverMaps() {
	fmt.Println("-----------------------------")
	fmt.Println("loopingOverMaps")

	m := map[int]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
	}

	for k, v := range m {
		fmt.Printf("k[%d] -> v[%s]\n", k, v)
	}

	fmt.Println("Plain values")
	for k, v := range m {
		fmt.Printf("%d -> %s\n", k, v)
	}

}

func loopingOverArrays() {
	fmt.Println("-----------------------------")
	fmt.Println("loopingOverArrays")
	/*
		Go provides a useful keyword, range, that simplifies
		looping over a variety of data types.

		In the case of arrays and slices, range can be used
		with a for statement to retrieve the index and the
		value of each element as it iterates:
	*/

	s := []int{2, 4, 6, 8, 16, 32}

	fmt.Println("For Loop with Range")
	for i, v := range s {
		fmt.Printf("i[%d] -> v[%d]\n", i, v)
	}

	fmt.Println("the unneeded values can be discarded by \n",
		"using the “blank identifier,” signified by the underscore operator:")

	a := []int{0, 2, 4, 6, 8}
	sum := 0

	for _, v := range a {
		sum += v
	}

	fmt.Printf("sum of the loop = %d\n", sum) // 20

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
