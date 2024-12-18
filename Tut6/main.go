package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

type electricEngine struct{
	kwhr uint8
	kw uint8
}

func (e gasEngine) milesleft() uint8{
	return e.gallons*e.mpg
}

func (e electricEngine) milesleft() uint8{
	return e.kwhr*e.kw
}

type engine interface{
	milesleft() uint8
}

func canmakeit(e engine, miles uint8){
	if miles<=e.milesleft(){
		fmt.Println("You can make it there")
	}else{
		fmt.Println("You cannot make it there")
	}
}

func main(){
	var myengine gasEngine = gasEngine{25, 15}
	fmt.Println(myengine.mpg, myengine.gallons, myengine.milesleft())
	var elecengine electricEngine = electricEngine{25, 15}
	fmt.Println(elecengine.kwhr, elecengine.kw, elecengine.milesleft())
	canmakeit(myengine, 30)
	canmakeit(elecengine, 120)
}
