package main

import (
	"CAREER-EDGE/modules"
	"fmt"
)

func main() {
	var menu int

	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                 <>   SYSTEM INITIALIZATION    <>              ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
	fmt.Println("║ >> Initializing Systems...                                    ║")
	fmt.Println("║ >> Loading modules: Resume Engine, CoverLetter Generator...   ║")
	fmt.Println("║ >> Welcome to your personal career assistant                  ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      [!!]    NOTICE    [!!]                      ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ This system is a specialized assistant designed ONLY to help     ║")
	fmt.Println("║ you build smart resumes and job-specific cover letters           ║")
	fmt.Println("║                                                                  ║")
	fmt.Println("║ [X] This is NOT a general chatbot                                ║")
	fmt.Println("║ [V] Use the menu options to interact effectively                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	for {
		fmt.Println("╔════════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                     C A R E E R  E D G E                               ║")
		fmt.Println("║       AI Assistant for Smart Resumes & Cover Letter Creation           ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ [1] Manage Profile                      | Kelola Profil                ║")
		fmt.Println("║ [2] Manage Job                          | Kelola Pekerjaan             ║")
		fmt.Println("║ [3] Job Recommendation <AI Assistance>  | Rekomendasi Pekerjaan        ║")

		// fmt.Println("║ [7] Add Job List                        | Tambahkan Lowongan Pekerjaan ║")
		// fmt.Println("║ [8] Cari Lowongan Pekerjaan             | Search Job Listing           ║")
		//fmt.Println("║ [9] Cari Posisi Pekerjaan               | Searh Job                    ║")
		// fmt.Println("║ [10] Urutkan Kata Kunci <Insertion>     | Sort by Keyword     ║")
		fmt.Println("║ [9] Help                                | Bantuan                      ║")
		fmt.Println("║ [0] Exit                                | Keluar                       ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("[!!!] Choose option / Pilih Menu: ")
		fmt.Scanln(&menu)

		if menu == 1 {
			modules.ManageProfile()
			modules.BackToMenu()
		} else if menu == 2 {
			modules.ManageJob()
			modules.BackToMenu()
		} else if menu == 3 {
			modules.RecommendJob()
			modules.BackToMenu()
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
