package main

import (
	"fmt"
	"math"
) 
	


var a =9 //Package Level 
func demo(){

	var a=8  //function level 
	fmt.Println("in demo", a)
}

// func cal(x, y int) (int, int){
	
// 	var out1 = x + y
// 	var out2 = x - y;
// 	return out1, out2
// }

func main(){

	var num float64 = 9
	var result = math.Sqrt(num)

	var intResult = math.Round(result)
	fmt.Println(intResult)

	fmt.Println(num)

	

	demo()
	fmt.Println("in main" ,a)
	// num1 :=5;
	// num2 :=8;
	


	
	// result1, result2 := cal(num1,num2)
	// fmt.Println(result1, result2)
}