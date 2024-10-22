package main

import (
	"bufio"   // Membaca inputan dari perngguna
	"fmt"     // Melakukan format output
	"os"      // Untuk fungsi sistem operasi membaca input
	"strconv" // Untuk mengkonversi dari string ke tipe data numerik dan sebaliknya.
	"strings" // Untuk fungsi operasi string.
)

// Struct untuk menyimpan data Calon
type Calon struct {
	Nama      string
	Partai    string
	Suara     int
	Threshold int
}

// Struct untuk menyimpan data Pemilih
type Pemilih struct {
	Nama  string
	Suara int
}

const MaxCalon = 5
const MaxPemilih = 100

var (
	calonList     [MaxCalon]Calon
	pemilihList   [MaxPemilih]Pemilih
	jumlahCalon   int
	jumlahPemilih int
	isAdmin       bool
	votingOpen    bool
)

// Fungsi untuk input dari pengguna
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Fungsi untuk menambah calon baru
func tambahCalon() {
	if jumlahCalon >= MaxCalon {
		fmt.Println("Tidak bisa menambah calon lagi.")
		return
	}
	nama := input("Masukkan nama calon: ")
	partai := input("Masukkan partai calon: ")
	threshold, _ := strconv.Atoi(input("Masukkan threshold calon: "))
	calonList[jumlahCalon] = Calon{nama, partai, 0, threshold}
	jumlahCalon++
	fmt.Println("Calon berhasil ditambahkan.")
}

// Fungsi untuk mengubah data calon
func ubahCalon() {
	tampilkanCalon()
	calonIndex, _ := strconv.Atoi(input(fmt.Sprintf("Pilih calon yang ingin diubah (1-%d): ", jumlahCalon)))
	if calonIndex > 0 && calonIndex <= jumlahCalon {
		calonList[calonIndex-1].Nama = input("Masukkan nama baru calon: ")
		calonList[calonIndex-1].Partai = input("Masukkan partai baru calon: ")
		threshold, _ := strconv.Atoi(input("Masukkan threshold baru calon: "))
		calonList[calonIndex-1].Threshold = threshold
		fmt.Println("Data calon berhasil diubah.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk menghapus calon
func hapusCalon() {
	tampilkanCalon()
	calonIndex, _ := strconv.Atoi(input(fmt.Sprintf("Pilih calon yang ingin dihapus (1-%d): ", jumlahCalon)))
	if calonIndex > 0 && calonIndex <= jumlahCalon {
		for i := calonIndex - 1; i < jumlahCalon-1; i++ {
			calonList[i] = calonList[i+1]
		}
		jumlahCalon--
		fmt.Println("Calon berhasil dihapus.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk menambah pemilih baru
func tambahPemilih() {
	if !votingOpen {
		fmt.Println("Voting saat ini ditutup.")
		return
	}
	if jumlahPemilih >= MaxPemilih {
		fmt.Println("Tidak bisa menambah pemilih lagi.")
		return
	}
	namaPemilih := input("Masukkan nama pemilih: ")
	suara, _ := strconv.Atoi(input(fmt.Sprintf("Masukkan pilihan suara (1-%d): ", jumlahCalon)))
	if suara > 0 && suara <= jumlahCalon {
		pemilihList[jumlahPemilih] = Pemilih{namaPemilih, suara}
		calonList[suara-1].Suara++
		jumlahPemilih++
		fmt.Println("Pemilih berhasil ditambahkan.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk menghapus pemilih
func hapusPemilih() {
	tampilkanPemilih()
	pemilihIndex, _ := strconv.Atoi(input(fmt.Sprintf("Pilih pemilih yang ingin dihapus (1-%d): ", jumlahPemilih)))
	if pemilihIndex > 0 && pemilihIndex <= jumlahPemilih {
		suara := pemilihList[pemilihIndex-1].Suara
		calonList[suara-1].Suara--
		for i := pemilihIndex - 1; i < jumlahPemilih-1; i++ {
			pemilihList[i] = pemilihList[i+1]
		}
		jumlahPemilih--
		fmt.Println("Pemilih berhasil dihapus.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// Fungsi untuk menampilkan daftar calon
func tampilkanCalon() {
	fmt.Println("\n--- Daftar Calon ---")
	for i := 0; i < jumlahCalon; i++ {
		fmt.Printf("%d. %s (%s)\n", i+1, calonList[i].Nama, calonList[i].Partai)
	}
}

// Fungsi untuk menampilkan daftar pemilih
func tampilkanPemilih() {
	fmt.Println("\n--- Daftar Pemilih ---")
	for i := 0; i < jumlahPemilih; i++ {
		fmt.Printf("%d. %s\n", i+1, pemilihList[i].Nama)
	}
}

// Fungsi untuk menampilkan hasil voting
func tampilkanHasilVoting() {
	fmt.Println("\n--- Hasil Voting ---")
	var pemenang Calon
	maxSuara := 0
	for i := 0; i < jumlahCalon; i++ {
		calon := calonList[i]
		fmt.Printf("Calon %d: %s (%s) - Suara: %d\n", i+1, calon.Nama, calon.Partai, calon.Suara)
		if calon.Suara >= calon.Threshold {
			fmt.Println("(Lolos)")
		} else {
			fmt.Println("(Tidak Lolos)")
		}
		if calon.Suara > maxSuara && calon.Suara >= calon.Threshold {
			maxSuara = calon.Suara
			pemenang = calon
		}
	}
	if maxSuara >= pemenang.Threshold {
		fmt.Printf("Pemenang: %s (%s) - Suara: %d (Lolos)\n", pemenang.Nama, pemenang.Partai, pemenang.Suara)
	} else {
		fmt.Println("Pemenang: Tidak ada pemenang (belum ada calon yang lolos)")
	}
}

// Fungsi untuk mencari calon berdasarkan nama atau partai
func cariCalon() {
	query := strings.ToLower(input("Masukkan nama calon atau partai: "))
	fmt.Println("\n--- Hasil Pencarian ---")
	for i := 0; i < jumlahCalon; i++ {
		calon := calonList[i]
		if strings.Contains(strings.ToLower(calon.Nama), query) || strings.Contains(strings.ToLower(calon.Partai), query) {
			fmt.Printf("%d. %s (%s) - Suara: %d\n", i+1, calon.Nama, calon.Partai, calon.Suara)
		}
	}
}

// Fungsi untuk mencari calon dengan sequential search
func sequentialSearchCalon(nama string) int {
	for i := 0; i < jumlahCalon; i++ {
		if strings.ToLower(calonList[i].Nama) == strings.ToLower(nama) {
			return i
		}
	}
	return -1
}

// Fungsi untuk mencari calon dengan binary search
func binarySearchCalon(nama string) int {
	left, right := 0, jumlahCalon-1
	for left <= right {
		mid := (left + right) / 2
		if strings.ToLower(calonList[mid].Nama) < strings.ToLower(nama) {
			left = mid + 1
		} else if strings.ToLower(calonList[mid].Nama) > strings.ToLower(nama) {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// Fungsi untuk mengurutkan calon dengan selection sort berdasarkan nama
func selectionSortCalonByNama(asc bool) {
	for i := 0; i < jumlahCalon-1; i++ {
		idx := i
		for j := i + 1; j < jumlahCalon; j++ {
			if asc {
				if calonList[j].Nama < calonList[idx].Nama {
					idx = j
				}
			} else {
				if calonList[j].Nama > calonList[idx].Nama {
					idx = j
				}
			}
		}
		calonList[i], calonList[idx] = calonList[idx], calonList[i]
	}
}

// Fungsi untuk mengurutkan calon dengan insertion sort berdasarkan nama
func insertionSortCalonByNama(asc bool) {
	for i := 1; i < jumlahCalon; i++ {
		key := calonList[i]
		j := i - 1
		for j >= 0 && ((asc && calonList[j].Nama > key.Nama) || (!asc && calonList[j].Nama < key.Nama)) {
			calonList[j+1] = calonList[j]
			j--
		}
		calonList[j+1] = key
	}
}

// Fungsi utama untuk menjalankan program
func main() {
	// Inisialisasi calon default awal
	calonList[jumlahCalon] = Calon{"Anies", "NASDEM", 0, 5}
	jumlahCalon++
	calonList[jumlahCalon] = Calon{"Prabowo", "GERINDRA", 0, 5}
	jumlahCalon++
	calonList[jumlahCalon] = Calon{"Ganjar", "PDIP", 0, 5}
	jumlahCalon++

	// Menu untuk admin
	menuAdmin := map[string]func(){
		"1": tambahCalon,
		"2": ubahCalon,
		"3": hapusCalon,
		"4": tambahPemilih,
		"5": hapusPemilih,
		"6": func() {
			votingOpen = !votingOpen
			if votingOpen {
				fmt.Println("Voting dibuka.")
			} else {
				fmt.Println("Voting ditutup.")
			}
		},
		"7": tampilkanHasilVoting,
		"8": cariCalon,
		"9": func() { isAdmin = false },
	}

	// Menu untuk pemilih
	menuPemilih := map[string]func(){
		"1": tambahPemilih,
		"2": tampilkanHasilVoting,
		"3": cariCalon,
		"4": func() {
			if input("Masukkan password: ") == "pilpres2024" {
				isAdmin = true
				fmt.Println("Anda sekarang berstatus admin.")
			} else {
				fmt.Println("Password salah. Silakan coba lagi.")
			}
		},
	}

	// Memilih role kembali
	for {
		fmt.Println("\n--- Menu Awal ---")
		fmt.Println("1. Petugas KPU")
		fmt.Println("2. Pemilih")

		switch input("Pilih role: ") {
		case "1":
			isAdmin = true
		case "2":
			isAdmin = false
		default:
			fmt.Println("Pilihan role tidak valid.")
			continue
		}
		break
	}

	// Menampilkan menu berdasarkan role
	for {
		fmt.Println("\n--- Menu Utama ---")
		if isAdmin {
			fmt.Println("1. Tambah Calon")
			fmt.Println("2. Ubah Calon")
			fmt.Println("3. Hapus Calon")
			fmt.Println("4. Tambah Pemilih")
			fmt.Println("5. Hapus Pemilih")
			fmt.Println("6. Buka/Tutup Voting")
			fmt.Println("7. Tampilkan Hasil Voting")
			fmt.Println("8. Cari Calon")
			fmt.Println("9. Ganti Role ke Pemilih")
			fmt.Println("0. Keluar")
			pilihan := input("Pilih menu: ")
			if action, exists := menuAdmin[pilihan]; exists {
				action()
			} else if pilihan == "0" {
				fmt.Println("Terima kasih telah menggunakan program voting.")
				break
			} else {
				fmt.Println("Pilihan menu tidak valid. Silakan pilih menu 1-9.")
			}
		} else {
			fmt.Println("1. Tambah Pemilih")
			fmt.Println("2. Tampilkan Hasil Voting")
			fmt.Println("3. Cari Calon")
			fmt.Println("4. Ganti Role ke Admin")
			fmt.Println("0. Keluar")
			pilihan := input("Pilih menu: ")
			if action, exists := menuPemilih[pilihan]; exists {
				action()
			} else if pilihan == "0" {
				fmt.Println("Terima kasih telah menggunakan program voting.")
				break
			} else {
				fmt.Println("Pilihan menu tidak valid. Silakan pilih menu 1-4.")
			}
		}
	}
}
