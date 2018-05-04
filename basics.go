package main
import (
  "fmt"
  "math/cmplx"
)

func add(x, y int) int {
	return x + y
}

// function can return any number of results
func swap(x, y string) (string, string) {
  return y, x
}

// return values can alos be named --> used to document the meaning of the name return values
// ie "naked" return
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

// `var` declares a list of variables
// unintialized variables are automatically given a falsey value of that type
var c, python, java bool

// var declaration can also include initializers, one per variable
var i, j int = 1, 2

// you do not have to type variables after `var`
// := can be used instead of var with implicit type
// but it can only be used inside a function
func equals() (int, int, int, bool, bool, string) {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"

  return i, j, k, c, python, java
}

// int 32 bits wide and uintptr 64 bits
func getNum() complex128 {
  return cmplx.Sqrt(-5 + 12i)
}

// implicit variable type is inferred from right
// you cannot leave untyped and undeclared though
func infer() {
  v := 42 // var v = 42
  fmt.Printf("v is of type %T\n", v)
}

// constants are devlared with `const`
// constants cannot be declared with :=
func con() bool {
  const Truth = true
  return Truth
}

// numeric constants are high-precision values
// an untyped constant takes the type needed by its context
const (
  Big = 1 << 100 // 1 followed by 100 zeroes
  Small = Big >> 99
)

func needInt(x int) int {
  return x * 10 + 1
}

func needFloat(x float64) float64 {
  return x * 0.1
}

func print() {
  fmt.Println(needInt(Small))
  // fmt.Println(needInt(Big)) will not print --> constant overflows int

  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}

// ====================MAIN========================================
func main() {
  fmt.Println(add(2, 3));

  a, b := swap("hello", "world")
  fmt.Println(a, b)

  fmt.Println(split(17))

  var i int
  fmt.Println(i, c, python, java)

  var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

  fmt.Println(equals())

  num := getNum()
  fmt.Printf("Type: %T Value: %v\n", num, num)

  infer()

  fmt.Println(con())

  print()
}

// =============================================================
// $ cd ~/go/src/hello
// $ go build
// $ ./hello
