package main

import (
	"errors"
	"fmt"
)

func main(){
  var helloworld string = "Hello World"
  print(helloworld)

  var numerator int = 10
  var denominator int = 0
  var division, remainder, err = division(numerator, denominator)
  switch{
    case err != nil:
      fmt.Println(err.Error())
    case remainder == 0:
      fmt.Printf("Result of the operations: Division - %v", division)
    default:
      fmt.Printf("Result of the operations: Division - %v Remainder - %v", division, remainder)   
  }
}


func print(printvalue interface{}){
  fmt.Println(printvalue)
}

func division(numerator int, denominator int) (int, int, error) {
  var err error
  if denominator == 0{
    err = errors.New("cannot divide by zero")
    return 0, 0, err
  }
  var division int = numerator / denominator
  var remainder int = numerator % denominator
  return division, remainder, err
}


