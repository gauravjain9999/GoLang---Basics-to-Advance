package main

import "fmt"

type Employee struct{

	EmpId int           // naming convention is in Capital or Pascal case to access in the globally scope 
	EmpName string
	DeptName string
	EmpMobile []string 

}

func main(){

	// fmt.Println("This is Demo of Struct ")

	emp := Employee{
		101,
	    "Gaurav",
		"IT",
		[]string{"9773754885", "9778521440"},
	}
	fmt.Println(emp)
}