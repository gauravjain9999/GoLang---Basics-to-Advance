package main

import "fmt"

func main() {

	num := 7

	//Switch Case
	switch num {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	default:
		fmt.Println("None")
	}

	// if num % 2==0{
	// 	fmt.Println("Even")
	// }else{
	// 	fmt.Println("Odd")
	// }

	// if num < 5 {
	// 	fmt.Println("Hi", num)
	// 	num++
	// } else {
	// 	fmt.Println("Bye" , num)
	// }

	// i :=1

	// for i<=5 {
	//   fmt.Println("Gaurav", i)
	//   i++;
	// }

	for i := 1; i <= 5; i++ {
		fmt.Println("Gaurav", i)
	}
}
