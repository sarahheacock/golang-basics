package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"io"
	"image"
	// "golang.org/x/tour/pic"
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

// can perform type assertion to interface's concrete value
func assert() (int, bool){
	var i interface{} = 0
	s, ok := i.(int)
	return s, ok // returns the value and if it matches type
}

// A type switch is a construct that permits several type assertions in series
// cases in a type switch specify types not values
func typeSwitch() {
	var i interface{} = 0

	switch v := i.(type) {
	case int:
		fmt.Printf("%v is an int\n", v)
	default:
		fmt.Printf("%v is NOT an int", v)
	}
}

// Stringer is an interface defined by fmt package
// Stinger is a type that can describe itself as a string
// type Stringer interface {
// 	String() string
// }
type Person struct {
	Name string
	Age  int
}

// type Person implements interface String
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer() {
	var p Person = Person{"Arthur Dent", 42}
	fmt.Println(p)
}

type MyError struct {
	When time.Time
	What string
}
// error type is built-in interface similar to fmt.Stringer
// type error interface {
// 	Error() string
// }
func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
// handle errors by testing whether error equals nil

func run() error {
	return MyError{
		time.Now(),
		"ruh roh",
	}
}

// io package specifies io.Reader interface which represents read end of stream of data
// io.Reader interface has Read method
// func (T) Read(b []byte) (n int, err error) -->
// returns number of bytes populated with an error value
// return io.EOF error when the stream ends
func read() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) // makes a zeroed array and a slice that references it

	// prints...
	// "Hello, R"
	// "eader!"
	// "" <-- with EOF error
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// io.Reader that wraps another io.Reader
// modifying the stream in some way
type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	l,e := r.r.Read(b)
	for i,c := range(b) {
      if c <= 'Z' && c >='A'  {
        b[i] = (c - 'A' + 13)%26 + 'A'
      } else if c >= 'a' && c <= 'z' {
        b[i] = (c - 'a' + 13)%26 + 'a'
      }
    }
	return l, e
}

// Package image defines the Image interface
// type Image interface {
//   ColorModel() color.Model
//   Bounds() Rectangle
//   At(x, y int) color.Color
// }
func createImage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

// create image
// type Image struct{
// 	Width, Height int
// 	colr uint8
// }
//
// func (r *Image) Bounds() image.Rectangle {
// 	return image.Rect(0, 0, r.Width, r.Height)
// }
//
// func (r *Image) ColorModel() color.Model {
// 	return color.RGBAModel
// }
//
// func (r *Image) At(x, y int) color.Color {
// 	return color.RGBA{r.colr+uint8(x), r.colr+uint8(y), 255, 255}
// }

// func createMyImage() {
// 	m := Image{100, 100, 128}
// 	pic.ShowImage(&m)
// }

func main() {
  v := Vertex{1, 2}
  // same as (&v).add()
  fmt.Println(v.add())
  fmt.Println(v)

  // same as (&v).change()
  fmt.Println(v.change())
  fmt.Println(v)

  var a Abser
	// a.Abs() --> there is no type inside the interface
	// tuple to indicate which concrete method to call
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
	fmt.Println(assert())
	typeSwitch()

	stringer()

	if err := run(); err != nil {
		fmt.Println(err)
	}

	read()

	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	createImage()
	// createMyImage()
}
