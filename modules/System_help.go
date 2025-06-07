package modules

import "fmt"

// TODO: Help menu
// Many applications use a help menu
func Help() {
	Clear()
	VPT()
	fmt.Println()
	// fmt.Println("╔════════════════════════════════════════════════════════════════════════╗")
	// fmt.Println("║                     Vitae PreTrained Transformer                       ║")
	// fmt.Println("║       AI Assistant for Smart Resumes & Cover Letter Creation           ║")
	// fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")
	Interact(" >> [System] This is help menu for user!")
	fmt.Println()

	// TODO: Commands section
	Interact(">> [Skills]   Pada saat menambahkan skill, ketik 'done' untuk berhenti atau selesai")
	Interact(">> [Commands] In main menu you can see multiple number and option, you can select the option using the number")
	Interact("              for example: type '1' for open Manage Skill or type '2' to open Manage Education")
	Interact(">> [Commands] Type 'l-ctrl + c' to stop the program")
	Interact(">> [Commands] Type 'cls' to clear terminal")
	Interact(">> [Commands] Type 'pback' for back to the main menu")
	fmt.Println()
}
