package main

import "fmt"

func main() {

	empSal := make(map[string]int) //declaration
	empSal = map[string]int{       // initialization
		"Neha":   2000,
		"Atul":   2500,
		"Gaurav": 45000,
	}

	//declaration with initialization
	empSal2 := map[string]int{

		"Rekha":  2000,
		"Jaya":   5000,
		"Sushma": 4000,
	}

	emp := empSal2

	empSal2["Rekha"] = 4000

	fmt.Println("Copy Value of one array to another Array", emp)
	fmt.Println("Original Array", empSal2)

	fmt.Println(empSal2)
	fmt.Println(len(empSal2))

	//to delete the element in array 
	delete(empSal2, "Jaya")
	fmt.Println(empSal2)


	//If value present in array then flag true otherwise false UnderScore basically define or symbolize that array we want to check 
	_, flag := empSal2["Gaurav"]
	fmt.Println(flag)

	fmt.Println(empSal2["Gaurav"]) //Trap Not having value but it will gives zero



	fmt.Println(empSal2["Rekha"])

	//decalartion
	fmt.Println(empSal)
}
