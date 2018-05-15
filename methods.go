package main

import (
	"fmt"
)

func max(x int) int {
  if(x < 0){
    return x * -1
  }
  return x
}

// go does not have classes
// but you can define methods on types
// Vertex literal is not a constant
type Vertex struct {
  X, Y int
}

type MyInt int

// interface is set of method signatures
type Abser interface {
	Abs() int
}

// type Vertex implements inteface Abser
// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.
func (v *Vertex) Abs() int {
  return max(v.X) + max(v.Y)
}

func (i MyInt) Abs() int {
  return max(int(i))
}

// method is just function with receiver object
// without pointer, operates on copy of original value Vertex (value receiver)
// receiver will be coerced from pointer to object
func (v Vertex) add() int {
  v.X = 3
  return v.X + v.Y
}

// method will coerce object to pointer if needed
// pointers are good for...
// (1) can modify value receiver points to
// (2) avoids inefficient copying of large structs
func (v *Vertex) change() int {
  v.X = 3 // automatically dereferenced
  return v.X + v.Y
}

func changeInterface(num *interface{}) {
	*num = "hello world"
}

// Empty interfaces are used by code that handles values of unknown type.
// one variable can change between different types
func changeType() interface{} {
	// var num int = 2;
	// num = "hello"; compile error
	var num interface{}
	num = 42 // type int
	num = "hello" // type string
	changeInterface(&num)

	return num;
}

func main() {
  v := Vertex{1, 2}
  // same as (&v).add()
  fmt.Println(v.add())
  fmt.Println(v)

  // same as (&v).change()
  fmt.Println(v.change())
  fmt.Println(v)

  var a Abser
  fmt.Printf("%T\n", a) //<nil>
	// a.Abs() will cause runtime error here
	// does not know which concrete method to run

  // var v Abser = Vertex{1, -2}
  // OR var v Vertex = Vertex{1, -2}
  v = Vertex{1, -2}
  i := MyInt(-1)
  // interface values can be thought of as a
  // tuple of a value and a concrete type
  // nil interface value will cause runtime error
  // fmt.Println(a.Abs())
  a = &v
  fmt.Println(a.Abs())
  fmt.Printf("%T\n", a) // *main.Vertex
  a = i
  fmt.Println(a.Abs()) // main.MyInt
  fmt.Printf("%T\n", a)

	fmt.Println(changeType())
}
