package modules

import "fmt"

func ManageJob() {
	var input int
	var running bool = true

	Clear()
	for running {
		fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                           [Job Manager]                            ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Menu]   Setting data lowongan pekerjaan                        ║")
		fmt.Println("║ 2. [Search] Cari data lowongan pekerjaan                           ║")
		fmt.Println("║ 3. [Search] Cari posisi pekerjaan berdasarkan gaji                 ║")
		fmt.Println("║ 4. [Show]   List data pekerjaan                                    ║")
		fmt.Println("║ 5. [Done]   Selesai dan simpan                                     ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&input)

		if input == 1 {
			Clear()
			JobMenu()
		} else if input == 2 {
			Clear()
			SearchJob()
		} else if input == 3 {
			Clear()
		} else if input == 4 {
			listJobListings()
			Clear()
		} else if input == 5 {
			running = false
		} else {
			fmt.Println("[HELP] Pilihan anda tidak ada di menu. Pilihlah menu yang tersedia dari angka 1 sampai 5")
		}
	}
}
