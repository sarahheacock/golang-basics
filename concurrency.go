package main
import (
  "fmt"
  "time"
)

// goroutine is lightweight thread managed by Go runtime
// `go f(x)` starts goroutine `f(x)`
// evaluation of f and x happens in current goroutine
// execution of f happens in new goroutine
// goroutine run in same address space so access to shared memory must be synchronized
func say(s string){
  for i := 0; i < 5; i++ {
    // if we take out sleep, just prints "hello"
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

// channels are a typed conduit that you can send and receive values
// `ch <- v` sends v to channel ch
// `v := <-ch` receives from ch and assigns to v
// `ch := make(chan int)` makes a channel
func sum(s []int, c chan int) {
  total := 0
  for _, val := range s {
    total += val
  }
  c <- total // send total to chan
}

// channels can be buffered
// blocks whenever buffer is full
func buffered() {
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  // ch <- 3 will say goroutines are asleep

  fmt.Println(<-ch) // 1
  fmt.Println(<-ch) // 2
}

// can `close(channel)`
// but do not NEED to--only to say no more values are coming

// `select` lets a goroutine wait on multiple communication operands
// `select` blocks until one of its cases can run
// it chooses one random if multiple are ready
func fib(c, quit chan int){
  x, y := 0, 1
  for {
    select {
    case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
    default:
      fmt.Println("    .")
      //time.Sleep(50 * time.Millisecond)
    }
  }
}

func main() {
  go say("world")
  say("hello")

  s := []int{7, 2, 8, -9, 4, 0}
  c := make(chan int)
  go sum(s[:len(s)/2], c) // 7, 2, 8
  go sum(s[len(s)/2:], c) // -9, 4, 0
  x, y := <-c, <-c // receive from c
  fmt.Println(x, y) // -5, 17

  f := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-f)
		}
		quit <- 0
	}()
	fib(f, quit)
}
