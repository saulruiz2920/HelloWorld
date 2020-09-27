package main

import "fmt"

func main() {
  var age int
  fmt.Print("Input your age: ")
  fmt.Scan(&age)
  fmt.Printf("You are %d years old\n", age)
}