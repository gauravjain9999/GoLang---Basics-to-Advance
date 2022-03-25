package main

import "fmt"

func main(){

	// var arr [4]int = [4]int{1,2,3,4} //In this we initialize or fix the size of the array

	//Another way 
	var arr = [...]int{1,2,3,5,6, 7,8}


	//Multi dimension Array 

	var matrix[3][3]int = [3][3] int{{1,2,3,}, {4,5,6}, {7,8,9}}

	fmt.Println(matrix)
	fmt.Println(len(matrix))
	fmt.Println(matrix[1][1])

	arr[1]= 20

	arr2 := arr
	fmt.Println(arr2)
	fmt.Println(arr)
	fmt.Println(len(arr))
}