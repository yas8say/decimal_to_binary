package main

import (
	"fmt"
	"os"
	"strconv"
)
func decimalToBinary(num int) string {
  //base cases
  if num == 1 {
    return "1"
  }
  if num == 0 {
    return "0"
  }

  return decimalToBinary(num/2) + strconv.Itoa(num%2)
}

func main() {
  //trying to convert to int and also capture the error if not convertible
  num, err := strconv.Atoi(os.Args[1])

  if err != nil {
    fmt.Println("Invalid Number!", os.Args[1])
    return
  }

  binary := decimalToBinary(num)
  fmt.Printf("Decimal: %d, Binary: %s\n", num, binary)
  
}