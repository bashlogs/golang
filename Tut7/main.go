package main

import "fmt"

func square(thing1 [5]float64) [5]float64{
	for i := range thing1{
		thing1[i] = thing1[i] * thing1[i]
	}
	return thing1
}

func square2(thing1 *[5]float64) [5]float64{
	for i := range thing1{
		thing1[i] = thing1[i] * thing1[i]
	}
	return *thing1
}

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Println(&p, &i)
	p = &i
	*p = 1
	fmt.Println(*p, i)


	var thing1 = [5]float64{1,2,3,4,5}
	var thing2 = square(thing1)
	fmt.Printf("Memory of thing2 %p and thing1 %p\n", &thing2, &thing1)
	fmt.Println(thing2, thing1)

	square2(&thing1)
	fmt.Printf("Memory of thing1 %p\n", &thing1)
	fmt.Println(thing1)
}