package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main()  {
	// rick := data.Instructor{Id: 1, FirstName: "Rick", LastName: "Sanchez", Score: 10}

	morty := data.NewInstructor("Morty", "Smith")

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: morty}

	swiftWS := data.NewWorkshop("Swift with iOS", morty)
	
	// Invoked interface for both workshop and course - This is also an example of polymorphism
	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}

}