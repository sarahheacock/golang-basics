package main
import (
  "fmt"
)

func pointers() int {
  var i int = 21
  var p *int = &i

  *p = 42
  fmt.Println(p) // print address of i
  return *p // print value at address p
}

// struct is collection of fields
// strings are immutable
// structs and ints are mutable
type Vertex struct {
  X int
  Y int
}

// will mutate the argument
// but will not mutate the original value
func mutate(v Vertex) Vertex {
  // p := &v
  // p.X = 5
  v.X = 5
  return v
}

// by sending pointer, we can mutate Vertex directly
func change(p *Vertex) {
  // *p = 42
  // struct pointer automatically dereferenced
  p.X = 5
}

func declare() (Vertex, Vertex, Vertex, *Vertex) {
  var (
    v1 = Vertex{1, 2}
    v2 = Vertex{Y: 2}
    v3 = Vertex{}
    p = &Vertex{1, 2}
  )
  return v1, v2, v3, p // &{1, 2}
}

// size is part of variable's type so
// array cannot be resized
func initArray() [2]string {
  var arr [2]string
  arr[0] = "hello"
  arr[1] = "world"
  return arr
}

func assignArray() [3]int {
  return [3]int{1, 2, 3}
}

// slice is dynamically sized view of elements of an array
// it does not store any data slice just represents array
func slicer() []int {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Printf("%T\n", primes)

	var s []int = primes[1:4] // first i and last i not included
  fmt.Printf("%T\n", s)
	return s
}

func changeSlicer() ([]string, []string) {
  names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
  names[3] = "world"

  a, b := names[0:2], names[0:2]
  a[0] = "hello"

  fmt.Println("Original:", names)
  return a, b
}

// slice literal is like array literal without the length
func sliceLit() []int {
  // array literal
  var arr [3]int = [3]int{1, 2, 3}
  arr[0] = 3
  fmt.Println(arr, len(arr), cap(arr))

  // slice literal -- builds array and then slices to reference it
  slice := []int{1, 2, 3}
  fmt.Println(slice, len(slice), cap(slice))
  return slice
}

// slice has len(s) and cap()
// length is number of elements in slice
// capacity is number of elements in array starting with first element in slice
func makeFunc() ([]int, []int) {
  b := make([]int, 3, 5) // len(b)=0, cap(b)=5
  return b[0:], b[0:cap(b)]
}

func sliceOfSlice() [][]string {
  return [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
}

// append adds new element to slice
// if underlying array has a large enough capacity, length increases
// if array is too small, then a new array is allocated
// the capacity will increase by 3 with certain numbers
func myAppend() []int {
  var s []int
  s = append(s, 1, 2)
  return s
}

func mutateArr(a [3]int) {
  a[2] = 5
  return
}

func mutateSlice(s []int) {
  s[2] = 5
}

func printRange() {
  pow := make([]int, 5)
  for i := range pow {
    pow[i] = i
  }

  for _, val := range pow {
    fmt.Printf("%d ", val)
  }
  fmt.Printf("\n")
}

// slice and map are both reference types
func mapSlice(m map[string]Vertex) {
  m["Google"] = Vertex{5, 3}
  delete(m, "foo")
  // m["Google"].X = 2 does not work
}

func mapLiteral(m map[string]Vertex) map[string]Vertex {
  // m["Google"].X = 5 does not work
  m["Google"] = Vertex{5, 3}
  m["foo"] = Vertex{}
  delete(m, "boo")

  return m
}

// functions are values too
// they can be passed around as arguments
func myFunc(x int, y int) int {
  return x + y
}

func caller(adder func(int, int) int, x int, y int) int {
  return adder(x, y)
}

// Go has closures!
func myClosure() func(int) int {
  counter := 0
  return func(i int) int {
    counter += i
    return counter
  }
}

// ========================================
func main() {
  fmt.Println(pointers())

  var v Vertex = Vertex{1, 2}
  p := &v
  p.X = 3
  mutate(v) // will not change v.X
  // v = mutate(v) will copy over
  fmt.Println(v)

  var x Vertex = Vertex{1, 2}
  change(&x)
  fmt.Println(x)

  fmt.Println(declare())
  fmt.Println(initArray())
  fmt.Println(assignArray())
  fmt.Println(slicer())
  fmt.Println(changeSlicer())
  fmt.Println(sliceLit())
  fmt.Println(makeFunc())
  fmt.Println(sliceOfSlice())

  arr := [3]int{1,2,3}
  arr[0] = 0
  mutateArr(arr)
  fmt.Println(arr) // 0, 2, 3

  s := arr[0:]
  mutateSlice(s)
  fmt.Println(s) // 0, 2, 5

  printRange()

  // can make slice of a map
  var sliceMap = make(map[string]Vertex)
  sliceMap["foo"] = Vertex{}
  mapSlice(sliceMap)
  fmt.Println(sliceMap)

  // can change a map's values
  var m = map[string]Vertex{
  	"Bell Labs": Vertex{
  		0, 1,
  	},
  	"Google": {
  		2, 3,
  	},
  }
  m["boo"] = Vertex{}
  mapLiteral(m)
  fmt.Println(m)

  fmt.Println(caller(myFunc, 1, 2))

  count := myClosure()
  for i := 0; i < 3; i++ {
    fmt.Println(count(i))
  }
}
