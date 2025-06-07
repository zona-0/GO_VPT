package modules

import (
	"CAREER-EDGE/data"
	"fmt"
)

// TODO: Manage experience menu sec
func ManageExperience() {
	var command, index, i int
	var title, company string
	var newExperience data.Experience
	var running bool = true

	for running {
		VPT()
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                      [Experience Manager]                          ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Add]    Tambah pengalaman kerja baru                           ║")
		fmt.Println("║ 2. [Edit]   Ubah pengalaman kerja yang sudah dimasukkan            ║")
		fmt.Println("║ 3. [List]   Lihat semua pengalaman kerja                           ║")
		fmt.Println("║ 4. [Delete] Hapus salah satu pengalaman                            ║")
		fmt.Println("║ 5. [Done]   Selesai                                                ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&command)

		if command == 1 {
			if data.ExperienceCount >= 1000000 {
				Interact(">> [System] Pengalaman sudah mencapai batas maksimum")
			} else {
				SuggestionExperience()
				fmt.Print("[>>] Job Title: ")
				fmt.Scan(&title)

				fmt.Print("[>>] Company Name: ")
				fmt.Scan(&company)

				newExperience.Title = title
				newExperience.Company = company
				data.Experiences[data.ExperienceCount] = newExperience
				data.ExperienceCount += 1

				Interact(">> [System] Pengalaman berhasil ditambahkan")
				Clear()
			}
		} else if command == 2 {
			if data.ExperienceCount == 0 {
				Interact(">> [System] Belum ada pengalaman untuk diubah")
			} else {
				fmt.Println("╔═══════╦════════════════════════╦════════════════════════╗")
				fmt.Println("║  No.  ║        Jabatan         ║        Perusahaan      ║")
				fmt.Println("╠═══════╬════════════════════════╬════════════════════════╣")
				i = 0
				for i < data.ExperienceCount {
					fmt.Printf("║  %-4d ║ %-22s ║ %-22s ║\n", i+1, data.Experiences[i].Title, data.Experiences[i].Company)
					i += 1
				}
				fmt.Println("╚═══════╩════════════════════════╩════════════════════════╝")

				fmt.Print("Masukkan nomor pengalaman yang ingin diubah (1-", data.ExperienceCount, "): ")
				fmt.Scan(&index)
				if index < 1 || index > data.ExperienceCount {
					Interact(">> [System] Nomor tidak valid")
				} else {
					SuggestionExperience()
					fmt.Print("[>>] Job Title Baru: ")
					fmt.Scan(&title)

					fmt.Print("[>>] Company Name Baru: ")
					fmt.Scan(&company)

					data.Experiences[index-1].Title = title
					data.Experiences[index-1].Company = company
					Interact(">> [System] Pengalaman berhasil diubah")
				}
				Clear()
			}
		} else if command == 3 {
			if data.ExperienceCount == 0 {
				Interact(">> [System] Belum ada pengalaman yang dimasukkan")
			} else {
				fmt.Println("╔═══════╦════════════════════════╦════════════════════════╗")
				fmt.Println("║  No.  ║        Jabatan         ║        Perusahaan      ║")
				fmt.Println("╠═══════╬════════════════════════╬════════════════════════╣")
				i = 0
				for i < data.ExperienceCount {
					fmt.Printf("║  %-4d ║ %-22s ║ %-22s ║\n", i+1, data.Experiences[i].Title, data.Experiences[i].Company)
					i += 1
				}
				fmt.Println("╚═══════╩════════════════════════╩════════════════════════╝")
			}
		} else if command == 4 {
			if data.ExperienceCount == 0 {
				Interact(">> [System] Belum ada pengalaman untuk dihapus")
			} else {
				fmt.Println("╔═══════╦════════════════════════╦════════════════════════╗")
				fmt.Println("║  No.  ║        Jabatan         ║        Perusahaan      ║")
				fmt.Println("╠═══════╬════════════════════════╬════════════════════════╣")
				i = 0
				for i < data.ExperienceCount {
					fmt.Printf("║  %-4d ║ %-22s ║ %-22s ║\n", i+1, data.Experiences[i].Title, data.Experiences[i].Company)
					i += 1
				}
				fmt.Println("╚═══════╩════════════════════════╩════════════════════════╝")

				fmt.Print("Masukkan nomor pengalaman yang ingin dihapus (1-", data.ExperienceCount, "): ")
				fmt.Scan(&index)
				if index < 1 || index > data.ExperienceCount {
					Interact(">> [System] Nomor tidak valid")
				} else {
					i = index - 1
					for i < data.ExperienceCount-1 {
						data.Experiences[i] = data.Experiences[i+1]
						i += 1
					}
					data.ExperienceCount -= 1
					Interact(">> [System] Pengalaman berhasil dihapus")
				}
				Clear()
			}
		} else if command == 5 {
			Interact(">> [System] Sesi input pengalaman selesai")
			running = false
		} else {
			Interact(">> [System] Pilihan tidak dikenali. Masukkan angka <1 - 5>")
		}
	}
}
