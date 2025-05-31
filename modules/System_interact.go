package modules

import (
	"fmt"
	"time"
)

// TODO: Interact func for simulating user interaction with AI
func Interact(text string) {
	var i int = 0
	for {
		if i >= len(text) {
			break
		}
		fmt.Printf("%c", text[i])
		time.Sleep(50 * time.Millisecond)
		i++
	}
	fmt.Println()
}

// TODO: Delay func for simulating a delay
func Delay() {
	time.Sleep(1 * time.Second)
}

// TODO: Simulasi ges
