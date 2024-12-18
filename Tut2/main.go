package main

import "fmt"

func main() {

	// int: 8, 16, 32, 64
	// int8: (-128, 127)
	// uint8: (0, 255)

	var intNum int16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var intNum2 uint16 = 32767
	intNum2 = intNum2 + 1
	fmt.Println(intNum2)

	// float: 32, 64
	var floatnum float64 = 12345678.9
	fmt.Println(floatnum)

	var floatnum2 float32 = 10.2
	var intnum3 int8 = 2
	var result float32 = floatnum2 + float32(intnum3)
	fmt.Println(result)

	var mystring string = "Hello world" + " " + "nice"
	fmt.Println(mystring)

	var mybool bool = true
	fmt.Println(mybool)

	myvar := "text"
	fmt.Println(myvar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myvalue int = 9
	fmt.Println(myvalue)
}