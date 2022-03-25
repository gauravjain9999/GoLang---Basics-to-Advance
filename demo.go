package main

import (
	"fmt"
	"strconv"
)


var val int = 100

func main(){

	// var a int    //declartion
	// a = 55     // Initilization

	// var b int = 500  //declaration with initlization
	// var val = 300

	//3rd Method
	// d :=10000

	var variableType = 500
	var str int = 65
	// var d string = string(str)
	var d string = strconv.Itoa(str) //Itoa -> Interger to a 


	fmt.Printf("%v, %T", d , d);
	fmt.Printf("%v, %T", variableType, variableType)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(val)
	// fmt.Println(d)
}

