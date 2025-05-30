package modules

import (
	"CAREER-EDGE/data"
	"fmt"
)

func MainAI() {
	var i int
	var user string
	//SuggestionAI()

	Clear()
	fmt.Print(">> ")
	fmt.Scan(&user)
	// Clear()
	// TODO: AI resume sec
	if user == "resume" || user == "RESUME" || user == "Resume" {
		fmt.Println("----------------------------------------------------------------------------------------------")
		fmt.Println(">> Tentu! untuk membuat resume yang efektif, aku butuh informasi dasar terlebih dahulu")
		fmt.Println(">> Silahkan isi data berikut: ")
		fmt.Println("----------------------------------------------------------------------------------------------")
		fmt.Println("[A] Informasi Pribadi: ")
		fmt.Printf("	[X] Nama: \n	[X] Alamat: \n	[X] No. HP: \n	[X] Email: \n")
		fmt.Println("[B] Jelaskan tentang dirimu: ")
		fmt.Println("	[X] About Me: ")
		// TODO: TAKE EDUCATION DATA
		fmt.Println("[C] Pendidikan: ")
		if data.EducationCount == 0 {
			fmt.Println("	[System] Anda belum mengisi data pendidikan. Silakan isi melalui menu utama!")
		} else {
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("	[System] Sekolah/Universitas: %s\n	[System] Jenjang: %s\n	[System] Tahun Lulus: %d\n",
					data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
		}

		// TODO: TAKE SKILLS DATA
		fmt.Println("[D] Keterampilan: ")
		if data.SkillCount == 0 {
			fmt.Println("	[System] Anda belum mengisi data keterampilan. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.SkillCount {
				fmt.Printf("	[System] %d. %s\n", i+1, data.Skills[i].Name)
				i += 1
			}
		}

		// TODO: PENGALAMAN KERJA
		fmt.Println("[E] Pengalaman Kerja: ")
		if data.ExperienceCount == 0 {
			fmt.Println("	[System] Anda belum mengisi data pengalaman kerja. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.ExperienceCount {
				fmt.Printf("	[System] %d. %s %s\n", i+1, data.Experiences[i].Title, data.Experiences[i].Company)
				i += 1
			}
		}

		// TODO: CERTIFICATION SEC
		fmt.Println("[F] Sertifikat: ")

		// TODO: END SEC
		fmt.Println()
		fmt.Printf(">> Silahkan isi informasi di atas atau cukup jawab yang kamu punya sekarang.\n   Setelah itu, aku akan buatkan resume versi teks yang rapi dan siap pakai\n")
		fmt.Println("----------------------------------------------------------------------------------------------")
		fmt.Println()

		// TODO: USER INPUT SEC
		var nama, alamat, hp, email string

		fmt.Print(">> Nama: ")
		fmt.Scan(&nama)
		fmt.Print(">> Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print(">> No. Hp: ")
		fmt.Scan(&hp)
		fmt.Print(">> Email: ")
		fmt.Scan(&email)
		fmt.Println()

		var kata, aboutme string
		var selesai bool = false

		fmt.Println(">> Tentang Saya: ")
		fmt.Println("   [System] Akhiri dengan tanda '.' untuk mengakhiri")
		for !selesai {
			fmt.Scan(&kata)

			if kata == "." {
				selesai = true
			} else {
				if aboutme == "" {
					aboutme = kata
				} else {
					aboutme = aboutme + " " + kata
				}
			}
		}

		var sertifikat string
		fmt.Print(">> Sertifikat: ")
		fmt.Scan(&sertifikat)

		// TODO: RESUME PRINT SECTION
		fmt.Printf("\n\n")
		fmt.Println(">> Berikut ini adalah resume yang berhasil saya susun berdasarkan data yang Anda berikan.")
		fmt.Println(">> Silakan periksa di bawah ini. Jika ada yang ingin direvisi, Anda bisa kembali kapan saja.")
		fmt.Printf("\n\n")
		fmt.Println("==================================================================")
		fmt.Printf("Nama: %s\nAlamat: %s\n\nTelp: %s | Email: %s\n", nama, alamat, hp, email)

		fmt.Println("------------------------------------------------------------")
		fmt.Println("Tentang Saya: ")
		fmt.Println(aboutme)

		fmt.Println("------------------------------------------------------------")
		fmt.Println("KETERAMPILAN:")
		if data.SkillCount == 0 {
			fmt.Println("  [System] Belum ada keterampilan")
		} else {
			for i = 0; i < data.SkillCount; i++ {
				fmt.Printf("  - %s\n", data.Skills[i].Name)
			}
		}

		fmt.Println("------------------------------------------------------------------")

		if data.EducationCount == 0 && data.ExperienceCount == 0 {
			fmt.Println("Pendidikan:")
			fmt.Println("  [System] Belum ada pendidikan")
			fmt.Println("Pengalaman Kerja:")
			fmt.Println("  [System] Belum ada pengalaman kerja")
		} else if data.EducationCount == 0 {
			fmt.Println("PENDIDIKAN:")
			fmt.Println("  [System] Belum ada pendidikan")
			fmt.Println("Pengalaman Kerja:")
			for i = 0; i < data.ExperienceCount; i++ {
				fmt.Printf("  - %s di %s\n", data.Experiences[i].Title, data.Experiences[i].Company)
			}
		} else if data.ExperienceCount == 0 {
			fmt.Println("Pendidikan:")
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("  - %s, %s (%d)\n", data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
			fmt.Println("Pengalaman Kerja:")
			fmt.Println("  [System] Belum ada pengalaman kerja")
		} else {
			fmt.Printf("%-38s| %-38s\n", ">> Pendidikan: ", ">> Pengalaman Kerja: ")

			var maxRows int
			if data.EducationCount > data.ExperienceCount {
				maxRows = data.EducationCount
			} else {
				maxRows = data.ExperienceCount
			}

			for i = 0; i < maxRows; i++ {
				var eduStr, expStr string

				if i < data.EducationCount {
					eduStr = fmt.Sprintf("- %s, %s %d", data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
				}
				if i < data.ExperienceCount {
					expStr = fmt.Sprintf("- %s di %s", data.Experiences[i].Title, data.Experiences[i].Company)
				}

				fmt.Printf("%-38s| %-38s\n", eduStr, expStr)
			}
		}

		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Sertifikat:")
		if sertifikat != "" {
			fmt.Printf("  - %s\n", sertifikat)
		} else {
			fmt.Println("  [System] Tidak ada sertifikat")
		}

		fmt.Printf("\n\n")
		fmt.Println(">> Resume telah selesai dibuat. Semoga harimu menyenangkan ", nama)
		fmt.Println()

		//endSec()

		//fmt.Println("TEST: ", aboutme)
		// fmt.Println("[A] Informasi Pribadi: ")
		// fmt.Printf("	[X] Nama: %s\n	[X] Alamat: %s\n	[X] No. HP: %s\n	[X] Email: %s\n", nama, alamat, hp, email)
		//ManageEducation()
	} else if user == "suratkerja" || user == "SuratKerja" || user == "Surat_Kerja" {
		fmt.Println()
		fmt.Println(">> Tentu! Sebelum aku membuat surat lamaran kerja, aku mau kasih contoh formatnya dulu, coba kamu baca ya!")
		fmt.Println(">> Silahkan isi data berikut: ")

		fmt.Println("-------------------------------------------------------------")
		fmt.Println("                        SURAT LAMARAN PEKERJAAN              ")
		fmt.Println("-------------------------------------------------------------")
		fmt.Println("Jakarta, 30 Mei 2025")
		fmt.Println()

		fmt.Println("Hal       : Lamaran Pekerjaan")
		fmt.Println("Lampiran  : 1 (satu) berkas")
		fmt.Println()

		fmt.Println("Kepada Yth.")
		fmt.Println("Bapak/Ibu HRD")
		fmt.Println("PT Maju Jaya Abadi, Jakarta Pusat")
		fmt.Println("Jl. Merdeka No. 123")
		fmt.Println()

		fmt.Println("Dengan hormat,")
		fmt.Println("Berdasarkan informasi lowongan kerja yang saya peroleh dari situs JobStreet pada tanggal 28 Mei 2025,")
		fmt.Println("saya bermaksud mengajukan lamaran untuk posisi \033[1mSoftware Engineer\033[0m di PT Maju Jaya Abadi.")
		fmt.Println("Saya sangat tertarik dengan visi perusahaan Bapak/Ibu dalam menghadirkan solusi teknologi inovatif.")
		fmt.Println()

		fmt.Println("Berikut data diri saya:")
		fmt.Println("-------------------------------------------------------------")
		fmt.Printf("%-20s: %s\n", "Nama", "Zona")
		fmt.Printf("%-20s: %s\n", "Tempat/Tgl Lahir", "Surakarta, 7 April 2006")
		fmt.Printf("%-20s: %s\n", "Pendidikan Terakhir", "S1 Teknik Informatika – Universitas Telkom")
		fmt.Printf("%-20s: %s\n", "Alamat", "Jl. Kenanga No. 45, Jakarta Timur")
		fmt.Printf("%-20s: %s\n", "Telepon", "0812-3456-7890")
		fmt.Printf("%-20s: %s\n", "Email", "zona@email.com")
		fmt.Println("-------------------------------------------------------------")
		fmt.Println()

		fmt.Println("Saya memiliki pengalaman 2 tahun sebagai Backend Developer di PT Teknologi Cerdas,")
		fmt.Println("dengan keahlian utama dalam REST API (Go, Node.js), PostgreSQL, dan arsitektur microservices.")
		fmt.Println("Saya terbiasa bekerja dalam tim, mampu belajar cepat, dan beradaptasi dengan lingkungan kerja.")
		fmt.Println()

		fmt.Println("Sebagai bahan pertimbangan, saya lampirkan:")
		fmt.Println("  1. Daftar Riwayat Hidup (CV)")
		fmt.Println("  2. Fotokopi Ijazah Terakhir")
		fmt.Println("  3. Fotokopi Transkrip Nilai")
		fmt.Println("  4. Sertifikat Keahlian")
		fmt.Println("  5. Pas Foto Terbaru")
		fmt.Println()

		fmt.Println("Demikian surat lamaran ini saya buat. Saya berharap dapat mengikuti tahapan seleksi berikutnya.")
		fmt.Println("Atas perhatian dan kesempatan yang diberikan, saya ucapkan terima kasih.")
		fmt.Println()

		fmt.Println("Hormat saya,")
		fmt.Printf("\n\n")
		fmt.Println("Zona")
		fmt.Println("-------------------------------------------------------------")
		fmt.Printf("\n\n")
		fmt.Println(">> Surat lamaran kerja di atas adalah contoh format yang umum digunakan")
		fmt.Println(">> Jika anda ingin membuat surat lamaran pekerjaan versi anda, saya dengan senang hati akan membantu!")
		fmt.Println()
		fmt.Println(">> Apakah anda ingin membuat surat lamaran kerja sekarang?")
		fmt.Println(">> [System] Ketik 'ya' untuk melanjutkan atau 'tidak' untuk kembali ke menu utama")

		var userInput string
		fmt.Print(">> ")
		fmt.Scan(&userInput)
		if userInput == "ya" || userInput == "Ya" || userInput == "YA" {
			fmt.Println()
			fmt.Println(">> Perhatikan data yang anda isi sebelumnya!")
			ShowAllUserData()

			fmt.Println()
			fmt.Println("Saya membutuhkan data lebih lanjut berkaitan dengan data diri anda: ")
			fmt.Println(">> [Input] Tempat, Tanggal Bulan Tahun: ")
			fmt.Println(">> [HELP] Jakarta, 30 Mei 2025")

			var tempat, bulan string
			var tanggal, tahun int

			fmt.Print(">> ")
			fmt.Scan(&tempat, &tanggal, &bulan, &tahun)

			// fmt.Printf("%s %d %s %d", tempat, tanggal, bulan, tahun)

		} else if userInput == "tidak" || userInput == "Tidak" || userInput == "TIDAK" {
			BackToMenu()
		}

	} else {
		fmt.Println()
		fmt.Println("[System] ketik 'resume' untuk bantuan membuat resume atau ketik 'suratkerja' untuk bantuan membuat surat lamaran kerja")
	}
}
