// name package 'main'
package main

// import package 'format'
import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

// need to run 'go mod init example.com/hello-world' to create a module - TODO: look into it

func main() {
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ MODULE 1 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ //
	///////////////////////////////////// PRINTING /////////////////////////////////////
	fmt.Printf("Hello, world!") // print a string
	a := "world"
	b := 3
	// print a var in a string
	fmt.Printf("Hello, %s", a)
	fmt.Printf("%d", b)

	// declare a type (alias)
	type Celsius float64
	var temp Celsius = 14.32
	fmt.Print(temp)

	///////////////////////////////////// USER INPUT /////////////////////////////////////
	// Scan reads user input (part of fmt package)
	// takes a pointer as an argument; typed data is written to pointer
	// returns 2 things: the number of scanned items + error ('nil' if none)
	fmt.Printf("Number of apples?")
	var appleNum int
	num, err := fmt.Scan(&appleNum)
	if err != nil {
		fmt.Print(appleNum, num)
	}

	///////////////////////////////////// POINTERS /////////////////////////////////////
	var x int = 100 // declare variable x (integer) and initialize it too (value of 100)
	var y int       // unitialized value is 0
	var ip *int     // ip is a pointer towards an integer (isn't an int)
	fmt.Println(x, y, ip)

	ip = &x // get the address of the variable x through '&'
	y = *ip // get the value at the address ip through '*'
	fmt.Println(x, y, ip)

	///////////////////////////////////// TYPE CONVERSION /////////////////////////////////////
	// convert between types using T()
	var m int32 = 2
	var n int16 = 1
	m = int32(n)

	// express floats as decimals or scientific notation
	var q float64 = 123.45
	var w float64 = 1.2345e2
	fmt.Println(x, y, q, w)

	z := new(int) // create a new variable; 'new()' returns a pointer to the variable
	*z = 3        // value at that pointer's address is 3
	fmt.Println(z, &z, *z)

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ MODULE 2 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ //

	///////////////////////////////////// STRINGS /////////////////////////////////////
	// ASCII = 8-bit character (max 2^8 = 256 options, works for english alphabet)
	// unicode = 32-bit character code (ideal for chinese because can go up to 2^32 "code points" aka unicode characters)
	// "code point" is called 'Rune' in Golang
	// UTF-8 is of variable length <- default in Golang

	// Unicode package
	r := '$'
	unicode.IsDigit(r)
	unicode.IsLower(r)
	unicode.IsUpper(r)
	unicode.IsSpace(r)
	unicode.IsLetter(r)
	unicode.ToUpper(r)
	unicode.ToLower(r)

	// String search (using 'strings' package)
	s := "abcdefggggg"
	strings.Compare("3", "4")  // Return 0 if a==b, -1 if a<b and +1 if a>b
	strings.HasPrefix(s, "hi") // return True if s starts with prefix
	strings.Contains(s, "abc") // return True if s contains substring
	strings.Index(s, "d")      // return the index of the first instance of substring

	// String manipulation
	strings.Replace(s, "g", "z", 3) // Replace 3 instances of 'g'->'z' in variable s
	strings.ToUpper(s)              // set to uppercase
	strings.ToLower(s)              // set to lowercase
	strings.TrimSpace(s)            // trim leading/trailing whitespaces

	// Strconv package for str manipulation
	strconv.Atoi(s)                            // converts str -> int
	strconv.Itoa(s)                            // converts int -> str
	strconv.FormatFloat(f, fmt, prec, bitSize) // convert float -> str
	strconv.ParseFloat(s, bitSize)             // convert str -> float

	///////////////////////////////////// CONSTANTS /////////////////////////////////////
	// Define constants (type is inferred from value)
	const p = 1.3
	const ( // can define many at once
		t = 4
		l = "Hi"
	)

	// iota generates a set of related but distinct constants ("one-hot", ex: days of the week)
	// constants are different (starts at 1 & increments) but the values aren't important
	type Grades int
	const (
		A Grades = iota // need to define for the first one; will infer 'iota' for all others
		B
		C
		D
		F
	)

	///////////////////////////////////// CONTROL FLOW /////////////////////////////////////
	// if else statements ("if <condition> {<consequence>}"")
	if x > 5 {
		fmt.Printf("x is greater than 5")
	}

	// for loops ("for <init>; <cond>; <update> {<stmts>}")
	for i := 0; i < 10; i++ {
		fmt.Printf("Value of i is %d", i)
	}
	j := 0
	for j < 10 {
		fmt.Printf("Value of j is %d", j)
		j++
	}

	// Switch (multi-case if statement)
	// may contain a tag (a variable to be compared to a constant in each case: 'x')
	// the case that matches the tag is executed
	switch x {
	case x < 3:
		fmt.Println("Case 1")
	case x > 5:
		fmt.Println("Case 2")
	default:
		fmt.Println("No case")
	}

	// Tagless switch (no variable to compare)
	// we will execute the first case that is true
	switch {
	case x < 3:
		fmt.Println("Case 1")
	case x > 5:
		fmt.Println("Case 2")
	default:
		fmt.Println("No case")
	}

	// Break exits the containing loop
	for i < 10 {
		i++
		if i == 5 {
			break
		}
	}

	// Continue skips the rest of the current iteration
	for i < 10 {
		i++
		if i == 5 {
			continue
		}
	}

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ MODULE 3 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ //
	///////////////////////////////////// ARRAYS /////////////////////////////////////
	// array: fixed length, get index by [], start at 0
	// elements are initialized to zero value (according to type)
	var arr [5]int // array of 5 integers
	arr[0] = 2     // change value at index 0

	// array literal: array predefined with values
	var lit_arr [5]int = [5]{1, 2, 3, 4, 5} // if declared of size 5, you need to provide 5 values
	// "..." infers size from number of initializers
	var o = [...]int {1, 2, 3, 4} // infers size 4

	// use a for loop to iterate over an array (i=index, v=element)
	for i, v := range o {
		fmt.Println("ind %d, val %d", i, v)
	}

	///////////////////////////////////// SLICES /////////////////////////////////////
	// slice: window on an (possibly) larger array
	// variable size, can increase up to the size of the whole array
	// can have many slices & they can overlap
		// pointer = start of the slice
		// length = # elements in the slice -> "len()"
		// capacity = max # elements -> "cap()"
	whole_array := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := whole_array[1:3] // b & c
	s2 := whole_array[2:5] // c, d, e, f
	fmt.Println(len(s1), cap(s1)) // will print 2, 7

	// writing to slices changes the underlying array
	// overlapping slices refer to the same elements
	fmt.Println(s1[1] == s2[0]) // True, since both refer to 'c'

	// slice literals (like array literals) are predefined with values
	// creates the underlying array at the same time (of same length as slice so length==capacity)
	sli := []int{1,2,3} // slices have no '...'

	// create a slice/array with 'make()' (can be called with 2 or 3 arguments)
	// 2 arguments: type & length/capacity (array == slice)
	make_2_slice := make([]int, 10)
	// 3 arguments: type, length & capacity (underlying array > slice)
	make_3_slice := make([]int, 10, 15)

	// 'append()' to add elements to the end of a slice
	// inserts into the array & increases its size if necessary
	my_slice := make([]int, 0, 3) // slice of size 0 but of capacity 3
	my_slice = append(my_slice, 100) // add one entry (100) to the end => slice is now of length 1 

	///////////////////////////////////// HASH TABLES /////////////////////////////////////
	// contain key:value pairs (unique keys) like a dictionary
	// 'hash function' computes the slot for a key 
	// advantages: faster lookup than lists + arbitrary keys
	// disadvantages: may have collisions (2 keys for the same spot)

	///////////////////////////////////// MAPS /////////////////////////////////////
	// implementation of a hash table
	// use 'make()' to make a map
	var idMap map[string]int // key type = string, value type = int
	idMap = make(map[string]int)
	name_map := map[string]int {"joe": 123}

	// access values through the keys
	fmt.Println(name_map["joe"])

	// add key:value pair through this (or change existing one)
	name_map["jane"] = 456

	// delete key:value pair
	delete(name_map, "jane")

	// find number of entries through len()
	fmt.Println(len(name_map))

	// two value assignment for the existence of key
	key_value, presence := name_map["joe"] // key_value = value, presence = boolean (key present or not)

	// iterating through a map
	for key, val := range name_map {
		fmt.Println(key, val)
	}

	///////////////////////////////////// STRUCTS /////////////////////////////////////
	// groups together objects of an arbitrary type (1 struct contains many variables)
	// ex: person (name, address, phone) where each property is a 'field'
	type struct Person {
		name string
		address string
		phone string
	}
	var p1 Person

	// access a field through dot notation
	p1.name = "Joe"
	x = p1.address

	// Initializing option 1: new() initializes all fields to 0
	p2 := new(Person)

	// Initializing option 2: struct literal
	p3 := Person(name: "Sandrine", address: "10 Blabla", phone: "12345")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ MODULE 4 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ //
	///////////////////////////////////// RFCs (Request for Comment) /////////////////////////////////////
	// Definitions of protocols (ex: HTML, URI, HTTP) and formats
	// Packages for immportant RFCs
	// ex: "net/http" is a web communication protocol: http.Get(www.uci.edu)
	// ex: "net" is a TCP/IP & socket programming: net.Dial("tcp", "uci.edu:80")

	// ex: Go object equivalent to JSON {"name": "Sandy", "address": "724 Thicket", "phone": "123"}
	p4 := Person(name: "Sandy", address: "724 Thicket", phone: "123")

	// "JSON Marshalling" = Go object to JSON []byte
	byte_array, error := json.Marshal(p4)

	// "JSON Unmarshalling" = JSON []byte to Go (object must fit JSON format)
	var p5 Person
	err := json.Unmarshal(byte_array, &p5) // pointer to p5 object is passed as input

	///////////////////////////////////// File access /////////////////////////////////////
	/* basic operations: 
	open (get handle for access), read (read bytes into []byte), 
	write (write []byte into file), close (close handle), seek (move read/write head) */

	// 'ioutil' package has basic functions
	data, e := ioutil.ReadFile("test.txt") // returns data []byte & error
	// explicit open/close not needed
	// large files can be a problemm
	dat = "Hello World"
	ioutil.WriteFile("test_out.txt", dat, 0777) // 3rd arg is permission

	// 'os' package
	// os.Open() opens a file (returns a file descriptor)
	// os.Close() closes a file
	// os.Read() reads from a file to fill a []byte (can control the amount read)
	f, err := os.Open("test.txt")
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	f.Close()
	// this reads & fills barr, reads # of bytes read (may be less than []byte length)

	// File Create/Write
	f, err := os.Create("outfile.txt")
	my_slice := []byte{1, 2, 3}
	nb, err := f.Write(my_slice)	 // writes a byte (any unicode sequence)
	nb, err := f.WriteString("hi") // writes a string

	///////////////////////////////////// FUNCTIONS /////////////////////////////////////
	// return type declared in function header
	def foo(x int) int {
		return x + 1
	}
	y := foo(1)

	// multiple return values
	def foo(x int) (int, int) {
		return x, x + 1
	}
	a, b := foo(1)

	///////////////////////////////////// CALL BY VALUE/REFERENCE /////////////////////////////////////
	// Call by value:
	// parameters are COPIED to a function (so modifications in function don't affect original variables)

	// Call by reference:
	// want to modify variables sent as parameters (use pointers instead)
	// advantage: saves times (no need to copy, especially large variables)
	func foo(y *int) {
		*y = *y + 1
	}
	x := 2
	foo(&x)
	fmt.Println(x) // will print 3

	// slices contain pointers to arrays, so you can use 'call by value' for those instead (better!)
	func foo(slice int) {
		sli[0] = sli[0] + 1
	}	

	///////////////////////////////////// FUNCTIONAL PROGRAMMING /////////////////////////////////////
	// VARIABLES AS FUNCTIONS
	var funcVar func(int) int // declare variable as a function
	func incFn(x int) int { // define function
		return x + 1
	}
	pp = incFn // set function into variable (no parentheses since not calling it, only assigning) 
	fmt.Print(pp(1))

	// VARIABLES AS ARGUMENTS
	// takes 'afunct' (function that takes an int & returns an int) & a integer as input
	func applyIt(afunct func (int) int), val int) int {
		return afunct(val)
	}
	// define 2 random functions
	func incFn(x int) int {return x+1}
	func decFn(x int) int {return x-1}
	// set functions as parameters of 'applyIt'
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	// ANONYMOUS FUNCTIONS (don't need a name)
	// aka lambda functions, defined on the spot
	v := applyIt(func (x int) int {return x+1}, 2)
	fmt.Println(v)

	///////////////////////////////////// VARIABLE NUMBER OF ARGUMENTS /////////////////////////////////////
	// can take 2,3,4... arguments to a function using ellipsis '...'
	// treated as a slice in a function
	func getMax(vals ...int) int {
		maxV := -1
		for _, v in range vals {
			if v > maxV {
				maxV = V
			}
		}
		return maxV
	}
	// can call like this:
	getMax(1, 2, 3, 4, 5) // list can be as long as you want
	vslice := []int{1, 2, 3, 4, 5}
	getMax(vslice...) // can pass a slice as argument

	// deferred variables (call deferred until function completes)
	// typically used for clean-up activity
	func my_main() {
		defer fmt.Println("Bye!")
		fmt.Println("Hello!") 
		// will execute Hello first & Bye at the very end of function
	}

	// arguments will be evaluated immediately, regardless of defer
	// ex: by the time the deferred line is called, the value of i is 2
	func my_main2() {
		i := 1
		defer fmt.Println(i)
		i++
		fmt.Println("Hello!")
	}

	///////////////////////////////////// CLASSES & ENCAPSULATION /////////////////////////////////////
	// Encapsulation: 
	// daa can be hidden from the programmer; can only be accessed through methods
	// maybe we don't trust the programmer to keep things consistent
	// ex: DoubleDistance (to double in both x & y directions)

	// no 'class' keyword in Go unlike most object-oriented languages

	// Associating methods with data
	// method has a 'receiver' type that it's associated with
	// use dot notation to call the method
	type MyInt int 
	func (mi MyInt) Double () int {
		return int(mi*2)
	}
	v := MyInt(3)
	fmt.PrintLn(v.Double()) 	// v is an implicit argument (since not passed as explicit input but still passed to function)

	// Use structs to manage many variables
	type Point struct {
		x float64
		y float64
	}
	func (p Point) DistToOrig () float64 {
		t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
		return math.Sqrt(t)
	}
	p1 := Point(3,4)
	fmt.Println(p1.DistToOrig())

	// Controlling access
	// define a public function to allow access to hidden data
	package data
	var x int = 1
	func PrintX() { fmt.Println(x) }
	
	// other packages can see value of x without directly accessing x
	// can't modify it
	package main
	import "data"
	func main() {
		fmt.Println(data.PrintX())
	}

	// Controlling access to structs
	// hide fields of structs by starting with lowercase
	package data
	type Point struct {
		x float64
		y float64
	}
	func (p *Point) InitMe(xn, yn float64) {
		p.x = xn
		p.y = yn
	}
	func (p *Point) Scale(v float64) {
		p.x = p.x * v
		p.y = p.y * v
	}
	func (p *Point) PrintMe() { fmt.Println(p.x, p.y) }

	// main gets access to hidden fields through public methods
	package main
	import "data"
	func main() {
		var p data.Point
		p.InitMe(3,4)
		p.Scale(2)
		p.PrintMe()
	}

	// these create copies of 'p', since theyre call by value

	///////////////////////////////////// POINT RECEIVERS /////////////////////////////////////
	// receiver can be a pointer, to change variable in place (call by reference)
	func (p *Point) OffsetX(v float64) {
		p.x = p.x + v
	}

	///////////////////////////////////// POLYMORPHISM /////////////////////////////////////
	// ability of an object to change 'forms' according to context
	// ex: area_rectangle = base x height, area_triangle = 1/2 base x height
	// identical at high-level abstraction (what they generally do), different at low-level abstraction (how they do it)
	// same signature (name, params & return type)
	// usually managed through Inheritance (not supported in Golang) where subclass inherits from the superclass

	// Interface = set of method signatures (implementation isn't defined)
	// used to express conceptual similarity between types
	// a type satisfies an interface if it defines all methods specified in the interface (same method signatures)

	// ex: Shape2D interface (represents 2-dimensional shapes) where all must have Area() & Perimeter()
	// Rectangle + Triangle types satisfy the Shape2D interface if they both have Area() & Perimeter()
	// additional methods are OK
	type Shape2D interface {
		Area() float64
		Perimeter() float64
	}
	type Triangle {...}
	func (t Triangle) Area() float64 {...}
	func (t Triangle) Perimeter() float64 {...}

	// Concrete vs interface types
	// Concrete types: 
		// Specify the exact representation of the data & methods
		// Complete method implementation is included

	// Interface types: 
		// Specifies some method signatures
		// Implementations are abstracted
	
	// Interface values can be treated like other values (assigned to variables, passed, returned)
	// have 2 components (come in pair):
		// dynamic type: concrete type to which it's assigned
		// dynamic value: value of the dynamic type
	type Speaker interface {Speak()}
	type Dog struct {name string}
	func (d Dog) Speak() {
		fmt.Println(d.name)
	}
	func main() {
		var s1 Speaker // s1 is an interface value
		var d1 Dog("Brian")
		s1 = d1        // dog type satisfies the s1 interface so s1'd interface type is d1
		s1.Speak()
		// dynamic type = Dog, dynamic value = d1
	}

	// Interface can have a 'nil' (empty) dynamic value
	var s2 Speaker
	var d2 *Dog
	s2 = d2
	// d2 isn't a concrete object (type points toward Dog but doesn't have any data)
	// d2 has no concrete value
	// s2 has a dynamic type but no dynamic value

	// can still call the Speak() function on s2 (doesn't need a dynamic value)
	// need to check inside the method
	func (d *Dog) Speak() {
		if d == nil {
			fmt.Println("<noise>")
		} else {
			fmt.Println(d.name)
		}
	}
	s2.Speak()

	// Nil interface value = interface with nil dynamic type
	// different from interface w/ nil dynamic value
	var s3 Speaker // no dynamic type = can't call Speak() on it (runtime error)

	// Ways to use an interface
	// 1- need a function that takes in multiple types of parameter
	// function foo() that wantss to take in type X or Y as parameter
	// define interface Z; foo() parameter is interface Z
	// types X and Y satisfy Z
	// ex: any shape (triangle, rectangle, square) can fit in yard as long as perimeter/area under threshold
	func FitInYard(s Shape2D) bool {
		if (s.Perimeter() <= 100 && s.Area() <= 100) {
			return true
		} else {return false}
	}

	// 2- empty interface defines no methods
	// all types satisfy the empty interface
	func PrintMe(val interface{}) {
		fmt.Println(val)
	}

	// Exposing type differences
	// interfaces allow us to hide differences between types
	// sometimes we need to treat different types differently
	// ex: Graphics program: DrawShape() will draw any shape
	func DrawShape (s Shape2D) {...}
	// underlying API have functions for each shape
	func DrawRectangle (r Rectangle) {...}
	func DrawTriangle (t Triangle) {...}
	// DrawShape() needs to call the correct one (identify concrete type of s) --> type assertions
	func DrawShape (s Shape2D) {
		rect, ok := s.(Rectangle) // concrete type in parentheses
		if ok { DrawRectangle(rect) }
		tri, ok := s.(Triangle)
		if ok { DrawTriangle(tri) }
	}
	// if recognized, rect=concrete type, ok=true
	// if not recognized, rect=zero, ok=false

	// use 'type switch' instead of if/else
	func DrawShape (s Shape2D) {
		switch := shape := s.(type) {
		case Rectangle:
			DrawRectangle(shape)
		case Triangle:
			DrawTriangle(shape)
		}
	}

	// many Go programs return an Error() interface to indicate errors
	type error interface {
		Error() string
	}
	// correct operation: error == nil
	// incorrect operation: Error() prints error message

	///////////////////////////////////// PARALLEL EXECUTION /////////////////////////////////////
	// parallel: two programs are executing at the same time
	// at time t, an instruction is being performed for both P1 & P2
	// ex: taskA -------------->
		// taskB -------------> 

	// concurrent: start & end times overlap
	// ex: taskA --->            taskA ---->
		//            taskB --->

	// parallel must be executed on diff hardware, while concurrent can be on the same

	///////////////////////////////////// GO ROUTINES /////////////////////////////////////
	// 1 goroutine is automatically created to run the main script
	// others are created using 'go' keyword

	// ex: this has 1 goroutine: main
	// main blocks on foo()
	b = 1
	foo()
	b = 2

	// ex: this has 2 goroutines: 1 for the main, 1 for foo()
	// main goroutine doesn't block
	a = 1
	go foo()
	a = 2

	// Early exit: goroutines exit when their code is completed
	// when the main goroutine is completed, all other goroutines are forced to exit
		// so your goroutine may not complete because the main finishes first
	// Ex: can't tell the order before running but you're likely to only get the 'main' printed
	// because the runtime scheduler prefers the main
	func main() {
		go fmt.Println("New goroutine")
		fmt.Println("Main goroutine")
	}
	// how to fix? delayed exit: add time.Sleep(100*time.Milliseconds) to make the main sleep (hack, not good because you assume the timing)
	
	///////////////////////////////////// SYNCHRONIZATION /////////////////////////////////////
	// what if there are multiple interleavings possible but you want a certain one?
	// ex: x=1, x=x+1 & print(x) vs x=1, print(x), x=x+1
	// want option 1 to run => need a global event whose execution is viewed by all threads simultaneously
	// may slow down the efficiency but can be necessary

	// sync waitgroup = package to synchronize goroutines
	sync.WaitGroup // forces a goroutine to wait for others
	// contains an internal counter:
		// increment for each goroutine you want to wait for: Add()
		// decrement when each goroutine completes: Done()
		// waits until counter=0: Wait()
	var wg sync.WaitGroup 
	wg.Add(1)
	go foo(&wg) //foo will run concurrently, outputting wg.Done() when completed
	wg.Wait() // blocks until counter=0

	// concrete example:
	func foo(wg *sync.WaitGroup) {
		fmt.Printf("New routine")
		wg.Done()
	}
	func main() {
		var wg sync.WaitGroup
		wg.Add(1)
		wg.Wait()
		go foo(&wg)
		fmt.Printf("Done main routine")
	}

	///////////////////////////////////// GOROUTINE COMMUNICATION /////////////////////////////////////
	// goroutines often need to send data to collaborate
	// ex: find the product of 4 ints: 2 goroutines, each multiplies a pair. main goroutine multiplies the 2 results
	// need to:
		// send initial values from main -> 2 subroutines
		// send results from subroutines -> main
	// use channels (typed) to transfer data
	// create using make(), send data using <-
	c := make(chan int)
	c <- 3 // send data on a channel
	x := <- c // receive data from a channel

	// concrete example
	func prod(v1 int, v2 int, c chan int) {c <- v1 * v2}
	func main() {
		c := make(chan int)
		go prod(1, 2, c)
		go prod(3, 4, c)
		a := <- c
		b := <- c
		fmt.Println(a * b)
	}

	// unbuffered channels can't hold data in transit (default)
	// sending blocks until data is received
	// receiving blocks until data is sent
	// ex: Task1          Task2
		c <- 3
				// 1h later
						x := <- c
	// Task1 can now continue

	// Blocking & synchronization
	// channel communication is synchronous
	// blocking is the same as waiting for communication
	// ex: Task1        Task2
		c <- 3           <- c
	// receiving & ignoring the result is the same as Wait()

	///////////////////////////////////// BUFFERED CHANNELS /////////////////////////////////////
	// channels can contain a limited # of objects
	// default size is 0 (unbuffered)
	// capacity = # objects it can hold in transit (optional value when calling make())
	c := make(chan int, 3) // capacity=3 --> can send up to 3x
	// sending only blocks if the buffer is full
	// receiving only blocks if the buffer is empty

	// Ex: channel (capacity=1) with 2 threads
	// Thread1  --> channel --> Thread2
	c <- 3						a := <- c
								b := <- c
	// first receive blocks until send occurs
	// once T1 executes, first receive of T2 can work
	// second receive in T2 blocks forever => error

	///////////////////////////////////// SYNCHRONIZED COMMUNICATION /////////////////////////////////////
	// Iterating over channels
	for i := range c { // continually iterates over channel
		fmt.Println(i)
	}
	// finishes when sender closes channel by close(c)

	// Receiving from multiple goroutines
	// multiple channels can be used to receive info from diff sources
	// data from diff sources may be needed
	a := <- c1
	b := <- c2
	fmt.Println(a*b)

	// "Select" statement used to read from certain channels (not necessarily all)
	// may have a choice of which data to use (first come, first serve)
	select {
		case a = <- c1:
			fmt.Println(a)
		case b = <- c2:
			fmt.Println(b)
	}

	// may either send or receive data (same first come, first serve logic)
	select {
	case a = <- in_channel:
		fmt.Println("Received a")
	case out_channel <- b:
		fmt.Println("Sent b")
	}

	// use Select with a special abort channel
	// for loop to keep receiving data until an abort signal is received
	for {
		select {
		case a <- c:
			fmt.Println(a)
		case <-abort:
			return
		}
	}

	// may want default case to avoid blocking
	select {
	case a = <- c1:
		fmt.Println("Received a")
	case b = <= c2:
		fmt.Println("Sent b")
	default:
		fmt.Println("Nope")
	}
	
	///////////////////////////////////// CONCURRENCY SAFE /////////////////////////////////////
	// goroutines sharing variables can be dangerous -> can interfere with one another
	// concurrency-safe code = will not do this

	///////////////////////////////////// MUTUAL EXCLUSION /////////////////////////////////////
	// don't let 2 goroutines write to a shared variable at the same time
	// need to restrict possible interleavings
	// access to shared variables can't be interleaved

	// mutual exclusion: code segments in diff goroutines that can't execute concurrently
	// writing to a shared variable must be mutually exclusive

	// Sync.Mutex package ensures mutual exclusion
	// uses a binary semaphore (flag up vs down):
		// flag up = shared variable in use
		// flag down = shared variable available
	
	// Mutex functions
	Lock() // puts the flag up
	// if lock already taken by another goroutine, Lock() blocks until the flag is taken down
	Unlock() // puts the flag down
	// when unlock is called, a blocked Lock() can proceed

	// Ex: mutually exclusive increment operation
	var i int = 0
	var mut sync.Mutex
	func inc() {
		mut.Lock()
		i = i + 1
		mut.Unlock()
	}

	///////////////////////////////////// ONCE SYNCHRONIZATION /////////////////////////////////////
	// synchronous initialization
	// initialization must happen once + before everything else
	// Sync.Once has one method: once.Do(f)
		// f will only be executed 1x, even if called in many goroutines
		// all calls to once.Do() block until the first returns => ensures that initialization runs first
	// Ex: make 2 goroutines, 1 initialization
	var on sync.Once
	func setup() {
		fmt.Println("Init")
	}
	func dostuff() {
		on.Do(setup)
		fmt.Println("Hello") // should not print until setup() returns
		wg.Done()
	}
	func main() {
		var wg sync.WaitGroup
		wg.Add(2)
		go do_stuff()
		go do_stuff()
		wg.Wait()
	}
	// Result: prints Init, Hello, Hello (in that order)
	// initialization (setup) only takes place once. two goroutines = two "Hellos"

	///////////////////////////////////// DEADLOCK /////////////////////////////////////
	// synchronization = execution of goroutines depend on one another
	// problem with synchronization => deadlock
	// Ex: G2 depends on G1 (can't assign variable to x until channel received 1 + can't lock until G1 unlocks)
	// G1                G2
	ch <- 1				x := <- ch
	mut.Unlock()		mut.Lock()
	
	// circular dependencies = all goroutines block (deadlock)
	// Ex: G1 waits for G2 & G2 waits for G1
	// can be caused by waiting on channels

	// Ex:
	func dostuff(c1 chan int, c2 chan int) {
		<- c1
		c2 <- 1
		wg.Done()
	}
	// read from 1st channel (wait to write onto 1st channel)
	// write to 2nd channel (wait for read from 2nd channel)
	func main() {
		ch1 := make(chan int)
		ch2 := make(chan int)
		wg.Add(2)
		go dostuff(ch1, ch2)
		go dostuff(ch2, ch1)
		wg.Wait()
	}
	// each goroutine blocked on channel read

	// Golang runtime automatically detects deadlocks
	// "fatal error: all goroutines are asleep - deadlock!"
	// Cannot detect when only a SUBSET of goroutines are deadlocked

	///////////////////////////////////// DINING PHILOSOPHER PROBLEM /////////////////////////////////////
	// classic problem with concurrency & synchronization
















}