package modules

import "fmt"

func SearchJob() {
	var choose int
	var running bool = true

	for running {
		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                   [MENU] Pencarian Lowongan                  ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. [Search]  Cari berdasarkan kata kunci industri            ║")
		fmt.Println("║ 2. [SEARCH]  Cari berdasarkan title pekerjaan                ║")
		fmt.Println("║ 3. [Done]    Selesai / Kembali                               ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Println("Pilih menu: ")
		fmt.Scan(&choose)

		if choose == 1 {
			//serachJobSequential()
		} else if choose == 2 {
			searchJobBinary()
		} else if choose == 3 {
			running = false
		} else {
			fmt.Println("[HELP] Pilih dari angka 1 sampai 3 untuk memilih menu")
		}
	}
}

// func serachJobSequential() {
// STUCK
// }

func searchJobBinary() {
	if jobCount == 0 {
		fmt.Println("[!] Tidak ada lowongan yang tersedia untuk dicari")
		fmt.Println()
		BackToMenu()
	}

	sortJobTitle()

	var keyword string
	fmt.Println("Masukan judul lowongan yang dicari: ")
	fmt.Scan(&keyword)

	var left int
	var right int = jobCount - 1
	var found bool = false
	var mid int = -1

	for left <= right {
		var m int = (right + left) / 2
		if jobListings[m].Title == keyword {
			mid = m
			found = true
			left = right + 1
		} else if jobListings[m].Title < keyword {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	if found {
		fmt.Println("╔═════════╦════════════════════════╦═════════════════╦══════════════╗")
		fmt.Println("║   No    ║        Title           ║    Industri     ║     Gaji     ║")
		fmt.Println("╠═════════╬════════════════════════╬═════════════════╬══════════════╣")

		var i int = mid
		var no int = 1
		for i >= 0 && jobListings[i].Title == keyword {
			fmt.Printf("╠ %-7d ╠ %-22s ╠ %-15s ╠ Rp %-9d ╠\n",
				no,
				jobListings[i].Title,
				jobListings[i].Industry,
				jobListings[i].Salary)
			i--
			no++
		}

		i = mid + 1
		for i < jobCount && jobListings[i].Title == keyword {
			fmt.Printf("╠ %-7d ╠ %-22s ╠ %-15s ╠ Rp %-9d ╠\n",
				no,
				jobListings[i].Title,
				jobListings[i].Industry,
				jobListings[i].Salary)
			i++
			no++
		}

		fmt.Println("╚═════════╩════════════════════════╩═════════════════╩══════════════╝")
	} else {
		fmt.Println("[!] Tidak ada lowongan dengan judul tersebut")
	}
}

func sortJobTitle() {
	var i int
	for i < jobCount-1 {
		var j int = i + 1
		for j < jobCount {
			if jobListings[i].Title > jobListings[j].Title {
				var temp JobListing
				temp = jobListings[i]
				jobListings[i] = jobListings[j]
				jobListings[j] = temp
			}
			j++
		}
		i++
	}
}
