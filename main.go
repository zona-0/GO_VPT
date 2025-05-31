package main

import (
	"CAREER-EDGE/modules"
	"fmt"
)

// Copyright © 2025 Zona
// dont use this code without permission!

func main() {
	var menu int

	modules.Clear()
	// SOURCE: https://patorjk.com/software/taag/#p=display&f=ANSI%20Regular&t=VPT
	// ANSI Reguler Font
	// SOURCE for UI: https://openai.com/index/chatgpt
	fmt.Println()
	fmt.Println("██ ███    ██ ████████ ██████   ██████  ██████  ██    ██  ██████ ██ ███    ██  ██████      ██        ██  ██████  ████████ ")
	fmt.Println("██ ████   ██    ██    ██   ██ ██    ██ ██   ██ ██    ██ ██      ██ ████   ██ ██            ██      ██   ██  ██     ██	   ")
	fmt.Println("██ ██ ██  ██    ██    ██████  ██    ██ ██   ██ ██    ██ ██      ██ ██ ██  ██ ██   ███       ██    ██    ██████     ██    ")
	fmt.Println("██ ██  ██ ██    ██    ██   ██ ██    ██ ██   ██ ██    ██ ██      ██ ██  ██ ██ ██    ██        ██  ██     ██         ██    ")
	fmt.Println("██ ██   ████    ██    ██   ██  ██████  ██████   ██████   ██████ ██ ██   ████  ███████          ██       ██	   ██    ")

	fmt.Println()
	modules.Interact(">> Vitae Pre-Trained Transformer")
	modules.Interact(">> Your AI Assistant for Smart Resumes & Cover Letter Creation")
	fmt.Println()

	// fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	// fmt.Println("║                 <>   SYSTEM INITIALIZATION    <>              ║")
	// fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
	// fmt.Println("║ >> Initializing Systems...                                    ║")
	// fmt.Println("║ >> Loading modules: Resume Engine, CoverLetter Generator...   ║")
	// fmt.Println("║ >> Welcome to your personal career assistant                  ║")
	// fmt.Println("╚═══════════════════════════════════════════════════════════════╝")

	// SOURCE: https://symbl.cc/en/unicode/blocks/box-drawing/
	modules.Delay()
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      [!!]    NOTICE    [!!]                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	modules.Delay()
	fmt.Println("║ This system is a specialized assistant designed ONLY to help     ║")
	modules.Delay()
	fmt.Println("║ you build smart resumes and job-specific cover letters           ║")
	fmt.Println("║                                                                  ║")
	modules.Delay()
	fmt.Println("║ [X] This is NOT a general chatbot                                ║")
	modules.Delay()
	fmt.Println("║ [V] Use the menu options to interact effectively                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	for {
		modules.Delay()
		fmt.Println("╔═══════════════════════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║               V I T A E  P R E - T R A I N E D  T R A N S F O R M E R                     ║")
		fmt.Println("║               AI Assistant for Smart Resumes & Cover Letter Creation                      ║")
		fmt.Println("╠═══════════════════════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ [1] Manage Profile                                         | Kelola Profil                ║")
		fmt.Println("║ [2] Manage Job                                             | Kelola Pekerjaan             ║")
		fmt.Println("║ [3] Job Recommendation             <AI Assistance>         | Rekomendasi Pekerjaan        ║")
		fmt.Println("║ [4] Create Resume or Cover Letters <AI Assistance>         | Buat Resume                  ║")
		fmt.Println("║ [9] Help                                                   | Bantuan                      ║")
		fmt.Println("║ [0] Exit                                                   | Keluar                       ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════════════════════════════════╝")
		// fmt.Print(">> Choose option / Pilih Menu: ")
		fmt.Print(">> Pilih menu:")
		fmt.Scan(&menu)

		if menu == 1 {
			modules.ManageProfile()
			modules.BackToMenu()
		} else if menu == 2 {
			modules.ManageJob()
			modules.BackToMenu()
		} else if menu == 3 {
			modules.RecommendJob()
			modules.BackToMenu()
		} else if menu == 4 {
			modules.MainAI()
			modules.BackToMenu()
		} else if menu == 9 {
			modules.Help()
			modules.BackToMenu()
		} else if menu == 0 {
			modules.Interact(">> Thank you for using Vitae Pre-Trained Transformer!")
			modules.Interact(">> Exiting the system...")
			break
		} else {
			modules.Interact("[System] Pilihan tidak valid. Silakan pilih menu yang tersedia")
		}

		// } else if menu == 6 {
		// 	modules.CreateResume() // TO BE UPDATE
		// 	modules.BackToMenu()
		// } else if menu == 7 {
		// 	modules.ManageJob()
		// 	modules.BackToMenu()
		// } else if menu == 8 {
		// 	modules.SearchJob()
		// } else if menu == 99 {
		// 	modules.Help()
		// 	modules.BackToMenu()
		// } else if menu == 0 {
		// 	fmt.Println("Ty")
		// 	break
		// }
		fmt.Println()
	}
}
