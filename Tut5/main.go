package main

import (
	"fmt"
)

func main(){
	var mystr = []rune("résumé")
	fmt.Printf("%v, %T\n", mystr[1], mystr[1])
	for i, v := range mystr{
		fmt.Println(i, v)
	}


}