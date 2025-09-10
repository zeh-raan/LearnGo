// go run LearnGo/main.go
package main

import "fmt" // for printing, input

// composite data type with multiple types
type hike struct {
	name       string
	location   string
	difficulty float64
}

// method in Go - attaches hike type to a function as a receiver
func (h hike) printHike() {
	fmt.Printf("Hike: %s\nLocation: %s\nDifficulty: %.1f\n",
		h.name, h.location, h.difficulty)
}

// Add hikes with a slice array
func addHike(slice []hike, h hike) []hike {
	return append(slice, h)
}

// put values in a neat table
func createTable() {}

// must have main function
func main() {
	var hikes []hike

	// slice hard coded
	hikes = addHike(hikes, hike{"Le Morne", "Black River", 3.5})

	// loop without index through each hikes
	for _, hike := range hikes {
		hike.printHike()
	}

	// allocate into map with the help of make
	// m := make(map[int]hike)
	// m[0] = hike{}
}
