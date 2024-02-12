package data

import "fmt"

type location string

// Method
func (origin location) DistanceTo(destination location) distance {
	// TODO caulations
	fmt.Printf("Origin %v Destination %v", origin, destination)
	return 10
}

func locationTest() {
	nyc := location("33.4234, 34.234")
	nyc.DistanceTo(location("-23, -44"))
	print(nyc)
}