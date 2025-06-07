package modules

import "fmt"

// TODO: This function is used to return to the main menu
func BackToMenu() {
	var backMenu string
	fmt.Println()
	Interact(">> [System] Type 'pback' to return to the main menu")
	for {
		fmt.Print(">> ")
		fmt.Scan(&backMenu)
		if backMenu == "pback" {
			break
		} else {
			fmt.Println()
		}
	}
	Clear()
}
