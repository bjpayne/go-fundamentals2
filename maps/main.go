package main

import "fmt"

func main() {
	dictionary := map[string]string{
		"Go":     "A programming language created by Google",
		"Gopher": "A software engineer who builds with Go.",
	}

	for _, language := range dictionary {
		fmt.Println(language)
	}
}
