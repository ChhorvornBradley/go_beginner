package main


import (
"fmt"
"strconv"
)


func fizzbuzz(i int) string {
  if i%15 == 0{
	return "FizzBuzz"
	} else if i%5 == 0{
        return "Buzz"
        } else if i%3 ==0 {
        return "Fizz"
        }
  return strconv.Itoa(i)
}

func main() {
	for i :=1; i <= 50; i++{
          fmt.Printf(fizzbuzz(i))
          fmt.Printf("\n")
        }
	fmt.Printf("Hello, world.\n")
}


