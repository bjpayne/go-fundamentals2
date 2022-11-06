package main

import (
	"fmt"
	"time"
)

func showMessages(messages []string) {
	for _, message := range messages {
		time.Sleep(1 * time.Second)

		fmt.Println(message)
	}
}

func main() {
	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}

	go showMessages(colorNames)

	showMessages(colorCodes)
}
