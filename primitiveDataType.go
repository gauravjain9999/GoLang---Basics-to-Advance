//Primitive Data Type
package main

import "fmt"

const (

	//by deafult it is zero and it increment by 1
	x = iota
	y = iota
)


func main(){

	// var a int32 = 32
	// a := 3.2
	// a := 4+5i

	//If there is no real value 
    // var a complex64 = 5i
	 
	//Passing Data Directly

	// c := complex(2,5)


	// const i int = 5645;
	// const j int = iota;

	fmt.Println(x);
	fmt.Println(y);





	// fmt.Printf("%v, %T", c, c)
	// fmt.Printf("%v, %v", real(c), imag(c))
	// fmt.Printf("%v, %T", a, a)
}