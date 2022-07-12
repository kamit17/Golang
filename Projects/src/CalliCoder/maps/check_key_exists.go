/*
When you retrieve the value assigned to a given key using the syntax map[key], it returns an additional boolean value as well which is true if the key exists in the map, and false if it doesn’t exist.

syntax:
value,ok :=m[key]
*/
package main

import "fmt"

func main() {
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}
