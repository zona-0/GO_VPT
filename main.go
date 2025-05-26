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
		fmt.Println("║ [1] Manage Skills                       | Kelola Skill                 ║")
		fmt.Println("║ [2] Manage Experience                   | Kelola Pengalaman            ║")
		fmt.Println("║ [3] Manage Education                    | Kelola Pendidikan            ║")
		fmt.Println("║ [4] Show User Data                      | Tampilkan Data Pengguna      ║")
		fmt.Println("║ [5] Job Recommendation <AI Assistance>  | Rekomendasi Pekerjaan        ║")
		fmt.Println("║ [6] Generate Cover Letter               | Buat Surat Lamaran Otomatis  ║") // SWITCH LAST MENU
		fmt.Println("║ [7] Add Job List                        | Tambahkan Lowongan Pekerjaan ║")
		// fmt.Println("║ [8] Cari Lowongan Pekerjaan             | Search Job Listing  ║")
		// fmt.Println("║ [9] Urutkan Gaji <Selection Sort>       | Sort by Salary      ║")
		// fmt.Println("║ [10] Urutkan Kata Kunci <Insertion>     | Sort by Keyword     ║")
		fmt.Println("║ [99] Help                               | Bantuan                      ║")
		fmt.Println("║ [0] Exit                                | Keluar                       ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════════╝")
		fmt.Print("[!!!] Choose option / Pilih Menu: ")
		fmt.Scanln(&menu)

		if menu == 1 {
			modules.AddSkill()
			modules.BackToMenu()
		} else if menu == 2 {
			modules.ManageExperience()
			modules.BackToMenu()
		} else if menu == 3 {
			modules.ManageEducation()
			modules.BackToMenu()
		} else if menu == 4{
			modules.ShowAllUserData()
			modules.BackToMenu()	
		} else if menu == 5{
			modules.RecommendJob()
			modules.BackToMenu()
		} else if menu == 6{
			modules.CreateResume() // TO BE UPDATE
			modules.BackToMenu()	
		} else if menu == 7{
			modules.ManageJob()
			modules.BackToMenu()	
		} else if menu == 99{
			modules.Help()
			modules.BackToMenu()
		} else if menu == 0 {
			fmt.Println("Ty")
			break
		}
		fmt.Println()
	}
}
