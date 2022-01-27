//Maps

/*Create a map that uses a string representing a course code as key.
The value in the map needs to be a structure with basic information about the course.
The following main routine:
*/

package main

import "fmt"

// Define a suitable structure
type Course struct {
	NStudents int
	Professor string
	Avg       float64
}

func main() {
	// Create a dynamic map m
	m := make(map[string]Course)

	// Add the courses CSI2120 and CSI2110 to the map
	m["CSI2110"] = Course{186, "Lang", 79.50}
	m["CSI2120"] = Course{211, "Moura", 81.0}

	for k, v := range m {
		fmt.Printf("Course Code: %s\n", k)
		fmt.Printf("Number of students: %d\n", v.NStudents)
		fmt.Printf("Professor: %s\n", v.Professor)
		fmt.Printf("Average: %f\n\n", v.Avg)
	}
}
