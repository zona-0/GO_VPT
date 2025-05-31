package modules

import (
	"fmt"
	"time"
)

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

func Delay() {
	time.Sleep(1 * time.Second)
}
