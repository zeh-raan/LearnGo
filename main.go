package main

import "fmt" // for printing text

// must have main function
func main() {
	const name = "Khan"
	p_language := "GO"
	// %v - replace the default parameter (int, string, etc...)
	fmt.Printf("Hello %v from %v\n", name, p_language)
	// %s - replace the parameter(s) into a string
	fmt.Printf("You are going to be a %s developer\n", "full stack")
	// %d - replace the parameter(s) into a decimal
	fmt.Printf("You have to be a %d/%d\n", 10, 10)
	// %f - replace the parameter(s) into a float64 and significant figure can be specified
	fmt.Printf("You are %.2f there", 99.9888)
}
