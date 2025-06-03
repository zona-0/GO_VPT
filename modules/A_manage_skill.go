package modules

import (
	"CAREER-EDGE/data"
	"fmt"
)

func ManageSkill() {
	var command, index, i int
	var nameSkill string
	var newSkill data.Skill
	var running bool = true
	var selesai bool = false

	for running {
		VPT()
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                          [Skill Manager]                           ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Add]     Tambah skill baru                                     ║")
		fmt.Println("║ 2. [Edit]    Ubah skill yang sudah dimasukkan                      ║")
		fmt.Println("║ 3. [List]    Lihat semua skill saat ini                            ║")
		fmt.Println("║ 4. [Delete]  Hapus salah satu skill                                ║")
		fmt.Println("║ 5. [Done]    Selesai                                               ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&command)

		if command == 1 {
			if data.SkillCount >= 50 {
				Interact(">> [System] Skill sudah mencapai batas maksimum")
			} else {
				SuggestionSkill()
				fmt.Println("Ketikkan beberapa skill satu per satu. Ketik 'done' jika sudah selesai")
				for {
					fmt.Print("Masukkan nama skill >>> ")
					fmt.Scan(&nameSkill)

					if nameSkill == "done" {
						break
					} else {
						newSkill.Name = nameSkill
						data.Skills[data.SkillCount] = newSkill
						data.SkillCount++

						if data.SkillCount == 50 {
							Interact(">> [System] Sudah mencapai maksimum 50 skill")
							break
						}
					}

					if selesai {
						break
					}
				}
				Interact(">> [System] Skill berhasil ditambahkan")
				Clear()
			}
		} else if command == 2 {
			if data.SkillCount == 0 {
				Interact(">> [System] Belum ada skill untuk diubah")
			} else {
				fmt.Println("╔═══════╦════════════════════╗")
				fmt.Println("║  No.  ║      Skill         ║")
				fmt.Println("╠═══════╬════════════════════╣")
				i = 0
				for i < data.SkillCount {
					var skillName string = data.Skills[i].Name
					fmt.Printf("║  %-4d ║ %-18s ║\n", i+1, skillName)
					i += 1
				}

				fmt.Println("╚═══════╩════════════════════╝")

				SuggestionSkill()

				fmt.Print(">> [System] Masukkan nomor skill yang ingin diubah (1-", data.SkillCount, "): ")
				fmt.Scan(&index)
				if index < 1 || index > data.SkillCount {
					Interact(">> [System] Nomor skill tidak valid")
				} else {
					fmt.Print("Masukkan skill baru >>> ")
					fmt.Scan(&nameSkill)
					data.Skills[index-1].Name = nameSkill
					Interact(">> [System] Skill berhasil diubah")
					Clear()
				}
			}
		} else if command == 3 {
			if data.SkillCount == 0 {
				Interact(">> [System] Belum ada skill yang dimasukkan.")
			} else {
				fmt.Println("╔═══════╦════════════════════╗")
				fmt.Println("║  No.  ║      Skill         ║")
				fmt.Println("╠═══════╬════════════════════╣")
				i = 0
				for i < data.SkillCount {
					var skillName string = data.Skills[i].Name
					fmt.Printf("║  %-4d ║ %-18s ║\n", i+1, skillName)
					i += 1
				}
				fmt.Println("╚═══════╩════════════════════╝")
			}
		} else if command == 4 {
			if data.SkillCount == 0 {
				Interact(">> [System] Belum ada skill untuk dihapus")
			} else {
				fmt.Println("╔═══════╦════════════════════╗")
				fmt.Println("║  No.  ║      Skill         ║")
				fmt.Println("╠═══════╬════════════════════╣")
				i = 0
				for i < data.SkillCount {
					var skillName string = data.Skills[i].Name
					fmt.Printf("║  %-4d ║ %-18s ║\n", i+1, skillName)
					i += 1
				}
				fmt.Println("╚═══════╩════════════════════╝")

				fmt.Print("Masukkan nomor skill yang ingin dihapus (1-", data.SkillCount, "): ")
				fmt.Scan(&index)
				if index < 1 || index > data.SkillCount {
					Interact(">> [System] Nomor skill tidak valid")
				} else {
					i = index - 1
					for i < data.SkillCount-1 {
						data.Skills[i] = data.Skills[i+1]
						i += 1
					}
					data.SkillCount = data.SkillCount - 1
					Interact(">> [System] Skill berhasil dihapus")
				}
				Clear()
			}
		} else if command == 5 {
			Interact(">> [System] Sesi input skill selesai")
			running = false
		} else {
			Interact(">> [System] Pilihan tidak dikenali. Masukkan angka <1 - 5>")
		}
	}
}
