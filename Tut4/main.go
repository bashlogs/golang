package main

import (
	"fmt"
)

func main(){
	// Array
	var intArr [3]int8 = [3]int8{1,2,3}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])
	fmt.Println(intArr[:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])


	var intslice []int8 = []int8{1,2,3}
	fmt.Printf("The length of the array is %v and capacity %v", len(intslice), cap(intslice))
	intslice = append(intslice, 1)
	fmt.Printf("\nThe length of the array is %v and capacity %v\n", len(intslice), cap(intslice))

	// var intslice2 []int8 = make(int8[], 8, 10)

	// Maps
	var map1 map[string]uint8 = make(map[string]uint8)
	map1 = map[string]uint8{"Adam": 23, "Mayur": 20}
	fmt.Println(map1["Adam"])

	// delete(map1, "Mayur")
	var age, ok = map1["Mayur"]
	if ok{
		fmt.Printf("This is age of %v", age)
	}else{
		fmt.Printf("Invalid Name")
	}

	for name := range map1{
		fmt.Printf("\nName: %v", name)
	}

	for i, v := range intArr{
		fmt.Printf("\nIndex - %v, Value - %v", i, v)
	}

}