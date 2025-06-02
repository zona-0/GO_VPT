package modules

import (
	"CAREER-EDGE/data"
	"fmt"
)

var kata, aboutme, sertifikat, nama, alamat, hp, email string
var selesai bool = false

func MainAI() {
	var i int
	var user string
	//SuggestionAI()

	Clear()
	fmt.Println()
	fmt.Println(" ██        ██  ██████  ████████ ")
	fmt.Println("  ██      ██   ██  ██     ██	   ")
	fmt.Println("   ██    ██    ██████     ██    ")
	fmt.Println("    ██  ██     ██         ██    ")
	fmt.Println("      ██       ██	  ██    ")

	fmt.Println()
	fmt.Println(">> Vitae Pre-Trained Transformer")
	fmt.Println(">> Your AI Assistant for Smart Resumes, Job Reccomendation and Cover Letter Creation")
	Interact(">> Ketik 'resume' untuk bantuan membuat resume atau ketik 'suratkerja' untuk bantuan membuat surat lamaran kerja")
	fmt.Println()
	fmt.Print(">> ")
	fmt.Scan(&user)
	// Clear()
	// TODO: AI resume sec
	if user == "resume" || user == "RESUME" || user == "Resume" {
		fmt.Println("----------------------------------------------------------------------------------------------")
		Interact(">> Tentu! untuk membuat resume yang efektif, aku butuh informasi dasar terlebih dahulu")
		Interact(">> Silahkan isi data berikut: ")
		fmt.Println("----------------------------------------------------------------------------------------------")
		Interact("[A] Informasi Pribadi: ")
		fmt.Printf("	[X] Nama: \n	[X] Alamat: \n	[X] No. HP: \n	[X] Email: \n")
		Interact("[B] Jelaskan tentang dirimu: ")
		Interact("	[X] About Me: ")
		// TODO: TAKE EDUCATION DATA
		Interact("[C] Pendidikan: ")
		if data.EducationCount == 0 {
			Interact("	>> [System] Anda belum mengisi data pendidikan. Silakan isi melalui menu utama!")
		} else {
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("	>> [V] Sekolah/Universitas: %s\n	>> [V] Jenjang: %s\n	>> [V] Tahun Lulus: %d\n",
					data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
		}

		// TODO: TAKE SKILLS DATA
		Interact("[D] Keterampilan: ")
		if data.SkillCount == 0 {
			Interact("	>> [System] Anda belum mengisi data keterampilan. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.SkillCount {
				fmt.Printf("	>> [V] %d. %s\n", i+1, data.Skills[i].Name)
				i += 1
			}
		}

		// TODO: PENGALAMAN KERJA
		Interact("[E] Pengalaman Kerja: ")
		if data.ExperienceCount == 0 {
			Interact("	>> [System] Anda belum mengisi data pengalaman kerja. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.ExperienceCount {
				fmt.Printf("	>> [V] %d. %s %s\n", i+1, data.Experiences[i].Title, data.Experiences[i].Company)
				i += 1
			}
		}

		// TODO: CERTIFICATION SEC
		Interact("[F] Sertifikat: ")

		// TODO: END SEC
		fmt.Println()
		fmt.Printf(">> Silahkan isi informasi di atas atau cukup jawab yang kamu punya sekarang.\n   Setelah itu, aku akan buatkan resume versi teks yang rapi dan siap pakai\n")
		fmt.Println("----------------------------------------------------------------------------------------------")
		fmt.Println()

		// TODO: USER INPUT SEC
		// var nama, alamat, hp, email string

		fmt.Print(">> Nama: ")
		fmt.Scan(&nama)
		fmt.Print(">> Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print(">> No. Hp: ")
		fmt.Scan(&hp)
		fmt.Print(">> Email: ")
		fmt.Scan(&email)
		fmt.Println()

		// var kata, aboutme string
		// var selesai bool = false

		Interact(">> Tentang Saya: ")
		Interact("   >> [System] Akhiri dengan klik 'enter' lalu ketik '.' untuk mengakhiri")
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

		// var sertifikat string
		fmt.Print(">> Sertifikat: ")
		fmt.Scan(&sertifikat)

		// TODO: RESUME PRINT SECTION
		fmt.Printf("\n\n")
		Interact(">> Berikut ini adalah resume yang berhasil saya susun berdasarkan data yang kamu berikan.")
		Interact(">> Silakan periksa di bawah ini. Jika ada yang ingin direvisi, Anda bisa kembali kapan saja.")
		fmt.Printf("\n\n")
		fmt.Println("==================================================================")
		fmt.Printf("Nama: %s\nAlamat: %s\n\nTelp: %s | Email: %s\n", nama, alamat, hp, email)

		fmt.Println("------------------------------------------------------------------")
		Interact("Tentang Saya: ")
		Interact(aboutme)

		fmt.Println("------------------------------------------------------------------")
		Interact("KETERAMPILAN:")
		if data.SkillCount == 0 {
			Interact("  >> [X] Belum ada keterampilan")
		} else {
			for i = 0; i < data.SkillCount; i++ {
				fmt.Printf("  - %s\n", data.Skills[i].Name)
			}
		}

		fmt.Println("------------------------------------------------------------------")

		if data.EducationCount == 0 && data.ExperienceCount == 0 {
			Interact("Pendidikan:")
			Interact("  >> [X] Belum ada pendidikan")
			Interact("Pengalaman Kerja:")
			Interact("  >> [X] Belum ada pengalaman kerja")
		} else if data.EducationCount == 0 {
			Interact("PENDIDIKAN:")
			Interact("  >> [X] Belum ada pendidikan")
			Interact("Pengalaman Kerja:")
			for i = 0; i < data.ExperienceCount; i++ {
				fmt.Printf("  - %s di %s\n", data.Experiences[i].Title, data.Experiences[i].Company)
			}
		} else if data.ExperienceCount == 0 {
			Interact("Pendidikan:")
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("  - %s, %s (%d)\n", data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
			Interact("Pengalaman Kerja:")
			Interact("  >> [X] Belum ada pengalaman kerja")
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
		Interact("Sertifikat:")
		if sertifikat != "" {
			fmt.Printf("  - %s\n", sertifikat)
		} else {
			Interact("  >> [X] Tidak ada sertifikat")
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
		Interact(">> Tentu! Sebelum aku membuat surat lamaran kerja, aku mau kasih contoh formatnya dulu, coba kamu baca ya!")

		Interact("-------------------------------------------------------------")
		Interact("                    SURAT LAMARAN PEKERJAAN                  ")
		Interact("-------------------------------------------------------------")
		Interact("Jakarta, 30 Mei 2025")
		fmt.Println()

		Interact("Hal       : Lamaran Pekerjaan")
		Interact("Lampiran  : 1 (satu) berkas")
		fmt.Println()

		Interact("Kepada Yth.")
		Interact("Bapak/Ibu HRD")
		Interact("PT Maju Jaya Abadi, Jakarta Pusat")
		Interact("Jl. Merdeka No. 123")
		fmt.Println()

		Interact("Dengan hormat,")
		Interact("Berdasarkan informasi lowongan kerja yang saya peroleh dari situs JobStreet pada tanggal 28 Mei 2025,")
		Interact("saya bermaksud mengajukan lamaran untuk posisi Fullstack Developer di PT Telkom.")
		Interact("Saya sangat tertarik dengan visi perusahaan Bapak/Ibu dalam menghadirkan solusi teknologi inovatif.")
		fmt.Println()

		Interact("Berikut data diri saya:")
		Interact("-------------------------------------------------------------")
		fmt.Printf("%-20s: %s\n", "Nama", "Zona")
		fmt.Printf("%-20s: %s\n", "Tempat/Tgl Lahir", "Surakarta, 7 April 2006")
		fmt.Printf("%-20s: %s\n", "Pendidikan Terakhir", "S1 Teknik Informatika – Universitas Telkom")
		fmt.Printf("%-20s: %s\n", "Alamat", "Jl. Kenanga No. 45, Jakarta Timur")
		fmt.Printf("%-20s: %s\n", "Telepon", "0812-3456-7890")
		fmt.Printf("%-20s: %s\n", "Email", "zona@email.com")
		Interact("-------------------------------------------------------------")
		fmt.Println()

		Interact("Saya memiliki pengalaman 2 tahun sebagai Backend Developer di PT Teknologi Cerdas,")
		Interact("dengan keahlian utama dalam REST API (Go, Node.js), PostgreSQL, dan arsitektur microservices.")
		Interact("Saya terbiasa bekerja dalam tim, mampu belajar cepat, dan beradaptasi dengan lingkungan kerja.")
		fmt.Println()

		Interact("Sebagai bahan pertimbangan, saya lampirkan:")
		Interact("  1. Daftar Riwayat Hidup (CV)")
		Interact("  2. Fotokopi Ijazah Terakhir")
		Interact("  3. Fotokopi Transkrip Nilai")
		Interact("  4. Sertifikat Keahlian")
		Interact("  5. Pas Foto Terbaru")
		fmt.Println()

		Interact("Demikian surat lamaran ini saya buat. Saya berharap dapat mengikuti tahapan seleksi berikutnya.")
		Interact("Atas perhatian dan kesempatan yang diberikan, saya ucapkan terima kasih.")
		fmt.Println()

		Interact("Hormat saya,")
		fmt.Printf("\n\n")
		Interact("Zona")
		Interact("-------------------------------------------------------------")
		fmt.Printf("\n\n")
		Interact(">> Surat lamaran kerja di atas adalah contoh format yang umum digunakan")
		Interact(">> Jika kamu ingin membuat surat lamaran pekerjaan versi kamu, aku dengan senang hati bakal bantu!")
		fmt.Println()
		Interact(">> Apakah kamu ingin membuat surat lamaran kerja sekarang?")
		Interact(">> Ketik 'ya' untuk melanjutkan atau 'tidak' untuk kembali ke menu utama")

		var userInput string
		fmt.Print(">> ")
		fmt.Scan(&userInput)
		if userInput == "ya" || userInput == "Ya" || userInput == "YA" {
			fmt.Println()
			Interact("Isi data diri kamu ya supaya aku bisa bantu buat surat lamaran kerjanya")
			fmt.Println()
			fmt.Println("-------------------------------------------------------------")
			fmt.Printf("%30s\n", "SURAT LAMARAN PEKERJAAN")
			fmt.Println("-------------------------------------------------------------")

			fmt.Println("______, __ ____ ____")
			var tempatS, bulanS string
			var tanggalS, tahunS int

			fmt.Print(">> ")
			fmt.Scan(&tempatS, &tanggalS, &bulanS, &tahunS)
			fmt.Println()

			fmt.Println("Hal       : Lamaran Pekerjaan")
			fmt.Println("Lampiran  : 1 (satu) berkas")
			fmt.Println()

			fmt.Println("Kepada Yth.")
			fmt.Println("Bapak/Ibu HRD")

			fmt.Println("PT _______, _______")
			var namaPT, alamatPT string
			fmt.Print(">> ")
			fmt.Scan(&namaPT, &alamatPT)

			fmt.Println("Jl. _____ No. __")
			var namaJalan string
			var noJalan int
			fmt.Print(">> ")
			fmt.Scan(&namaJalan, &noJalan)

			fmt.Println("Dengan hormat,")
			fmt.Println("Berdasarkan informasi lowongan kerja yang saya peroleh dari situs _______ pada tanggal __ ____ ____,")
			var situs, bulanA string
			var tanggalA, tahunA int
			fmt.Print(">> ")
			fmt.Scan(&situs, &tanggalA, &bulanA, &tahunA)

			fmt.Printf("saya bermaksud mengajukan lamaran untuk posisi ________ di PT %s\n.", namaPT)
			fmt.Print(">> ")
			var posisiPekerjaan string
			fmt.Scan(&posisiPekerjaan)

			Interact("Saya sangat tertarik dengan visi perusahaan Bapak/Ibu dalam menghadirkan solusi teknologi inovatif.")
			fmt.Println()

			fmt.Println("-------------------------------------------------------------")
			var namaLengkap, tempatLahir, bulanLahir, alamatJ, alamatK, email string
			var tanggalLahir, tahunLahir, noJ, noHp int
			fmt.Println("Nama                : _______")
			fmt.Print(">> ")
			fmt.Scan(&namaLengkap)

			fmt.Println("Tempat/Tgl Lahir    : _______, __ ______ ____")
			fmt.Print(">> ")
			fmt.Scan(&tempatLahir, &tanggalLahir, &bulanLahir, &tahunLahir)

			// fmt.Println("Pendidikan Terakhir : __ __________ - ______")
			if data.EducationCount > 0 {
				var last = data.Educations[data.EducationCount-1]
				fmt.Printf("Pendidikan Terakhir : %s - %s\n", last.Degree, last.Degree)
			} else {
				fmt.Println("Pendidikan Terakhir : -")
			}

			fmt.Println("Alamat              : Jl. _____ No. __, _____")
			fmt.Print(">> ")
			fmt.Scan(&alamatJ, &noJ, &alamatK)

			fmt.Println("Telepon             : _______________")
			fmt.Print(">> ")
			fmt.Scan(&noHp)

			fmt.Println("Email               : __________@email.com")
			fmt.Print(">> ")
			fmt.Scan(&email)
			fmt.Println("-------------------------------------------------------------")
			fmt.Println()

			if data.ExperienceCount > 0 {
				var lastExp = data.Experiences[data.ExperienceCount-1]
				var pengalaman int

				fmt.Printf("Saya memiliki pengalaman __ tahun sebagai %s di PT %s,\n", lastExp.Title, lastExp.Company)
				fmt.Print(">> ")
				fmt.Scan(&pengalaman)

				//fmt.Printf("Saya memiliki pengalaman %d tahun sebagai %s di %s,\n", pengalaman, last.Title, last.Company)
			} else {
				fmt.Println("")
			}

			if data.SkillCount > 0 {
				Interact("dengan keahlian utama dalam ")

				var i int
				for i = 0; i < data.SkillCount; i++ {
					Interact(data.Skills[i].Name)

					if i == data.SkillCount-2 {
						fmt.Print(" dan ")
					} else if i < data.SkillCount-2 {
						fmt.Print(", ")
					}
				}
				fmt.Println()
			} else {
				fmt.Println("")
			}

			Interact("Saya terbiasa bekerja dalam tim, mampu belajar cepat, dan beradaptasi dengan lingkungan kerja.")
			fmt.Println()

			fmt.Println("Sebagai bahan pertimbangan, saya lampirkan:")
			fmt.Println("  1. Daftar Riwayat Hidup (CV)")
			fmt.Println("  2. Fotokopi Ijazah Terakhir")
			fmt.Println("  3. Fotokopi Transkrip Nilai")
			fmt.Println("  4. Sertifikat Keahlian")
			fmt.Println("  5. Pas Foto Terbaru")
			fmt.Println()

			Interact("Demikian surat lamaran ini saya buat. Saya berharap dapat mengikuti tahapan seleksi berikutnya.")
			Interact("Atas perhatian dan kesempatan yang diberikan, saya ucapkan terima kasih.")
			fmt.Println()

			fmt.Println("Hormat saya,")
			fmt.Println()
			fmt.Printf("%s\n", namaLengkap)
			fmt.Println("-------------------------------------------------------------")

			Interact(">> Terima kasih telah mengisi data diri kamu! Berikut ini adalah surat lamaran kerja yang telah aku buat berdasarkan data yang kamu berikan.")

			Interact("-------------------------------------------------------------")
			Interact("                    SURAT LAMARAN PEKERJAAN                  ")
			Interact("-------------------------------------------------------------")
			fmt.Printf("%s %d %s %d", tempatS, tanggalS, bulanS, tahunS)
			fmt.Println()

			Interact("Hal       : Lamaran Pekerjaan")
			Interact("Lampiran  : 1 (satu) berkas")
			fmt.Println()

			Interact("Kepada Yth.")
			Interact("Bapak/Ibu HRD")
			fmt.Printf("PT %s, %s\n", namaPT, alamatPT)
			fmt.Printf("Jl. %s No. %d\n", namaJalan, noJalan)

			Interact("Dengan hormat,")
			fmt.Printf("Berdasarkan informasi lowongan kerja yang saya peroleh dari situs %s pada tanggal %d %s %d", situs, tanggalA, bulanA, tahunA)
			fmt.Printf("saya bermaksud mengajukan lamaran untuk posisi %s di PT %s.", posisiPekerjaan, namaPT)
			Interact("Saya sangat tertarik dengan visi perusahaan Bapak/Ibu dalam menghadirkan solusi teknologi inovatif.")
			fmt.Println()

			Interact("Berikut data diri saya:")
			Interact("-------------------------------------------------------------")
			fmt.Printf("%-20s: %s\n", "Nama", namaLengkap)
			fmt.Printf("%-20s: %s, %d %s %d\n", "Tempat/Tgl Lahir", tempatLahir, tanggalLahir, bulanLahir, tahunLahir)
			// fmt.Printf("%-20s: %s\n", "Pendidikan Terakhir", "S1 Teknik Informatika – Universitas Telkom")
			if data.EducationCount > 0 {
				var last = data.Educations[data.EducationCount-1]
				fmt.Printf("Pendidikan Terakhir : %s - %s\n", last.Degree, last.Degree)
			} else {
				fmt.Println("Pendidikan Terakhir : -")
			}

			fmt.Printf("%-20s: %s %d %s\n", "Alamat", alamatJ, noJ, alamatK)
			fmt.Printf("%-20s: %d\n", "Telepon", noHp)
			fmt.Printf("%-20s: %s\n", "Email", email)
			Interact("-------------------------------------------------------------")
			fmt.Println()

			if data.ExperienceCount > 0 {
				var lastExp = data.Experiences[data.ExperienceCount-1]
				var pengalaman int

				fmt.Printf("Saya memiliki pengalaman __ tahun sebagai %s di PT %s,\n", lastExp.Title, lastExp.Company)
				fmt.Print(">> ")
				fmt.Scan(&pengalaman)

				//fmt.Printf("Saya memiliki pengalaman %d tahun sebagai %s di %s,\n", pengalaman, last.Title, last.Company)
			} else {
				fmt.Println("")
			}

			if data.SkillCount > 0 {
				Interact("dengan keahlian utama dalam ")

				var i int
				for i = 0; i < data.SkillCount; i++ {
					Interact(data.Skills[i].Name)

					if i == data.SkillCount-2 {
						fmt.Print(" dan ")
					} else if i < data.SkillCount-2 {
						fmt.Print(", ")
					}
				}
				fmt.Println()
			} else {
				fmt.Println("")
			}

			Interact("Saya terbiasa bekerja dalam tim, mampu belajar cepat, dan beradaptasi dengan lingkungan kerja.")
			fmt.Println()

			Interact("Sebagai bahan pertimbangan, saya lampirkan:")
			Interact("  1. Daftar Riwayat Hidup (CV)")
			Interact("  2. Fotokopi Ijazah Terakhir")
			Interact("  3. Fotokopi Transkrip Nilai")
			Interact("  4. Sertifikat Keahlian")
			Interact("  5. Pas Foto Terbaru")
			fmt.Println()

			Interact("Demikian surat lamaran ini saya buat. Saya berharap dapat mengikuti tahapan seleksi berikutnya.")
			Interact("Atas perhatian dan kesempatan yang diberikan, saya ucapkan terima kasih.")
			fmt.Println()

			Interact("Hormat saya,")
			fmt.Printf("\n\n")
			fmt.Printf("%s\n", namaLengkap)
			Interact("-------------------------------------------------------------")

			fmt.Printf("\n\n")
			Interact(">> Surat lamaran kerja di atas hanyalah rekomendasi dari aku! Kamu bisa mengubahnya sesuai dengan kreatifitas kamu")
			Delay()
			fmt.Printf(">> Terima kasih telah menggunakan layanan VPT, semoga sukses dalam mencari pekerjaan ya %s", namaLengkap)
		} else if userInput == "tidak" || userInput == "Tidak" || userInput == "TIDAK" {
			BackToMenu()
		}
	} else {
		fmt.Println()
		Interact(">> [System] ketik 'resume' untuk bantuan membuat resume atau ketik 'suratkerja' untuk bantuan membuat surat lamaran kerja")
	}
}
