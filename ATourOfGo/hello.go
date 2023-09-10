package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

/*
   Go's basic types are
   bool

   string

   int  int8  int16  int32  int64
   uint uint8 uint16 uint32 uint64 uintptr

   byte // alias for uint8

   rune // alias for int32
       // represents a Unicode code point

   float32 float64

   complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// A var statement can be at package or function level.
var c, python, java bool

// A var declaration can include initializers, one per variable.
var ii, jj int = 1, 2

// Constants are declared like variables, but with the const keyword.
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {

	// A var statement can be at package or function level.
	fmt.Println("My favorite number is", rand.Intn(10))

	// The name of the function begins with a capital letter, it is exported from the package.
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// Pi is exported from math package.
	fmt.Println(math.Pi)

	// A name is exported if it begins with a capital letter.
	fmt.Println(add(42, 13))

	// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	fmt.Println(addV2(42, 13))

	// A var statement can be at package or function level.
	a, b := swap("hello", "world")

	// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
	fmt.Println(a, b)

	// The := short assignment statement can be used in place of a var declaration with implicit type.
	fmt.Println(split(17))

	// Variables declared without an explicit initial value are given their zero value.
	var i int
	fmt.Println(i, c, python, java)

	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var cc, python, java = true, false, "no!"
	fmt.Println(ii, jj, cc, python, java)

	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	var iii, jjj int = 1, 2
	kkk := 3
	ccc, pythonn, javaa := true, false, "no!"
	fmt.Println(iii, jjj, kkk, ccc, pythonn, javaa)

	// The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
	// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*
	   Variables declared without an explicit initial value are given their zero value.

	   The zero value is:

	   0 for numeric types,
	   false for the boolean type, and
	   "" (the empty string) for strings.
	*/
	var iiii int
	var ffff float64
	var bbbb bool
	var ssss string
	fmt.Printf("%v %v %v %q\n", iiii, ffff, bbbb, ssss)

	/*
	   Some numeric conversions:

	   var i int = 42
	   var f float64 = float64(i)
	   var u uint = uint(f)
	   Or, put more simply:

	   i := 42
	   f := float64(i)
	   u := uint(f)
	*/
	var x, y int = 3, 4
	var fffff float64 = math.Sqrt(float64(x*x + y*y))
	var zzzzz uint = uint(fffff)
	fmt.Println(x, y, zzzzz)

	/*
		When the right hand side of the declaration is typed, the new variable is of that same type:
		var i int
		j := i // j is an int
		But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

		i := 42           // int
		f := 3.142        // float64
		g := 0.867 + 0.5i // complex128
	*/
	vvvvvv := 42 // change me!
	fmt.Printf("vvvvvv is of type %T\n", vvvvvv)

	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Numeric constants are high-precision values.
	// An untyped constant takes the type needed by its context.
	// (An int can store at maximum a 64-bit integer, and sometimes less.)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}

// A function can take zero or more arguments.
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func addV2(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// Return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
