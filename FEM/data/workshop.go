package data

import "time"

type Workshop struct {
	Course // Embedding type within another type
	Instructor
	Date time.Time
}

func (w Workshop) SignUp() bool {
	return true
}

// Factory
func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop {}
	w.Name = name
	w.Instructor = instructor
	return w
}