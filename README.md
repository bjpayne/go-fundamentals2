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
