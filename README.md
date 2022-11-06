# UDACITY - Go Fundamentals 1

Maps, Structs, Interfaces, Goroutines

### Maps

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := map[int]string{
		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}

	// Not necessarily in order...weird
	for index, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Printf("%d: %s\n", index, course)
		}
	}

	fmt.Println("---------------------------")

	courses[4] = "Algorithms"
	courses[8] = "Spanish"

	delete(courses, 1)

	// Not necessarily in order...weird
	for index, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Printf("%d: %s\n", index, course)
		}
	}
}
```

### Structs

```go
package main

import "fmt"

type Student struct {
	id   uint16
	name string
}

type Classroom struct {
	id          uint16
	capacity    uint16
	subject     string
	studentList []Student
}

func main() {
	students := []Student{
		{
			id:   1,
			name: "Ben",
		},
		{
			id:   2,
			name: "Bob",
		},
		{
			id:   3,
			name: "Sally",
		},
		{
			id:   4,
			name: "Sam",
		},
	}

	c1 := Classroom{
		id:          1,
		capacity:    4,
		subject:     "History",
		studentList: students,
	}

	c2 := new(Classroom)

	c2.id = 2
	c2.capacity = 4
	c2.subject = "Math"
	c2.studentList = students

	fmt.Println(c1)
	fmt.Println(c2)
}
```
