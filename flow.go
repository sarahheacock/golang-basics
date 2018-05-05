package main
import (
  "fmt"
  "math"
  "runtime"
)

// := is declaration and assignment
// = is only assignment
func forLoop() int {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }

  return sum
}

// drop the semicolons and we get a while loop
func whileLoop() int {
  sum := 1
  for sum < 1000 {
    sum += sum
  }

  return sum
}

// if statements do not need parentheses
// sqrt takes in float64
func sqrt(x float64) string {
  if x < 0 {
    num := sqrt(-x)
    return num + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

// variables can declared within if statement and then last
// till the end of if statement
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  return lim // cannot ref v here
}

func Sqrt(x float64) float64 {
  z := x
	num := z - 0.1
  buffer := 0.0000000000000004

	for x < num - buffer || x > num + buffer {
		z -= (z * z - x) / (2 * z)
		num = z * z
	}

	return z
}

// break is automatically added in switch case
// switch cases also do not need to be constant
// if no switch variable, then default
func os() string {
  var result string;

  fmt.Print("Go runs on ")

  // switch os := runtime.GOOS; os {

  os := runtime.GOOS
  switch os {
    case "darwin":
      result = "OS X"
    case "linux":
      result = "Linux"
    default:
      result = os
  }

  return result
}

// defer holds execution of a function until surrounding function returns
// deferred call's arguments evaluated immediately but call no executed until
// surrounding function returns
func hold() {
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}

// ==========================================================
func main() {
  fmt.Println(forLoop())
  fmt.Println(whileLoop())

  // num starts as "2i"
  // but then become -4 in sqrt func
  num := sqrt(-4)
  fmt.Println(sqrt(2), num)

  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20), // need , for new line
  )

  fmt.Println(Sqrt(2))

  fmt.Println(os())

  hold()
}

// go run flow.go
