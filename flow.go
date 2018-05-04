package main
import "fmt"


func forLoop() int {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }

  return sum
}

func main() {
  fmt.Println(forLoop())
}

// go run flow.go
