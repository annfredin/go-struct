package main

import (
	"encoding/json"
	"fmt"
)


type Employee struct{
	Name string `json:"EmpName"`
	Age int
	Salary float64 `json:"Salary,omitempty"`
}
func main() {

	//empty/default declaration
	emp := Employee{}
	fmt.Printf("Emp => %+v\n", emp)

	//With field names declaration
	empFred := Employee{ Name: "Ann Fredin", Age: 30, Salary: 10000}
	fmt.Printf("Emp Fredin => %+v\n", empFred)

	//Without field names declaration
	empPricks := Employee{  "Ann Pricks", 30,  10000}
	fmt.Printf("Emp Pricks => %v+\n", empPricks)

	//A compiler error will be raised if all values are not provided when field name is not passed
	// empPricks1 := Employee{ "Ann Pricks",  30} //./main.go:25:43: too few values in struct literal
	empPricks1 := Employee{ Name: "Ann Pricks", Age: 30} //valid
	fmt.Printf("Emp Pricks1 => %+v\n", empPricks1)

	//Pointer ref ( 1st Way)
	empPointer := &empPricks1
	fmt.Printf("Emp empPointer => %+v\n", empPointer)
	fmt.Printf("Emp empPointer Values=> name: %v, age: %v\n", empPointer.Name, empPointer.Age)

	//Pointer ref ( 2nd Way "new")
	empP := new(Employee)
	empP1 := &Employee{}
	fmt.Printf("Emp empP => %+v\n", empP)
	fmt.Printf("Emp empP1 => %+v\n", empP1)


	//diff ways to print struct
	fmt.Printf("\nPrint Format\n" )
	fmt.Printf("Emp-> witout fields : %v\n", emp)
    fmt.Printf("Emp-> with fields :  %+v\n", emp)
    fmt.Printf("Emp-> with Object name and fields : %#v\n", emp)


	//Anonymous struct
	anonStr := struct{
		Name string 
		Age int
		}{ Name: "AAAA", Age: 30}
fmt.Println(anonStr)

	//JSON============
	fmt.Printf("\nJSON\n" )

	 //make sure fields should be capitalized to marshal
	 //Fields has JSON Meta or Tag, then will be exported with that name
	empJSON, _ := json.Marshal(empFred)

	empFred.Salary = 0
	empJSON1, _ := json.Marshal(empFred)
	
    fmt.Printf("Marshal function output %s\n", string(empJSON))
    fmt.Printf("Marshal function output %s\n", string(empJSON1))

	//Different Tags for fields...
	// You will often find tags in encoding packages, such as XML, JSON, YAML, ORMs, and Config management.
	type User struct {
		Id         string `csv:"user_id"`
		Name       string `json:"user_name,omitempty"`
		Occupation string `xml:"user_occupation"`
		Password  string    `json:"-"`
	}
	
	// omitempty -> if the value is empty(default value), then will not be exported.
	// - => used to totally ignore export fn, sensitive information not required to export
}