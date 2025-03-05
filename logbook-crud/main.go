package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struktur data untuk menyimpan informasi karakter
type Character struct {
	Nama   string
	Umur   int
	Bounty int
	Hobi   string
}

// Slice untuk menyimpan data karakter
var characters []Character

// Membuat reader global untuk input dengan spasi
var reader = bufio.NewReader(os.Stdin)

// Fungsi untuk membaca input dengan spasi
func readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi untuk menampilkan semua karakter
func getAllCharacters() {
	fmt.Println("\n📜 Daftar Karakter:")
	if len(characters) == 0 {
		fmt.Println("Tidak ada data karakter.")
		return
	}
	for i, c := range characters {
		fmt.Printf("%d. Nama: %s | Umur: %d | Bounty: %d | Hobi: %s\n", i+1, c.Nama, c.Umur, c.Bounty, c.Hobi)
	}
}

// Fungsi untuk menambahkan karakter baru
func addCharacter() {
	nama := readInput("Masukkan Nama: ")
	umurStr := readInput("Masukkan Umur: ")
	bountyStr := readInput("Masukkan Bounty: ")
	hobi := readInput("Masukkan Hobi: ")

	umur, err1 := strconv.Atoi(umurStr)
	bounty, err2 := strconv.Atoi(bountyStr)

	if err1 != nil || err2 != nil {
		fmt.Println("⚠️ Umur dan bounty harus berupa angka!")
		return
	}

	characters = append(characters, Character{Nama: nama, Umur: umur, Bounty: bounty, Hobi: hobi})
	fmt.Println("\n✅ Karakter berhasil ditambahkan!")
}

// Fungsi untuk memperbarui karakter berdasarkan indeks
func updateCharacter() {
	getAllCharacters()
	indexStr := readInput("\nMasukkan nomor karakter yang ingin diperbarui: ")
	index, err := strconv.Atoi(indexStr)

	if err != nil || index < 1 || index > len(characters) {
		fmt.Println("⚠️ Indeks tidak valid.")
		return
	}

	index-- // Sesuaikan dengan indeks slice (mulai dari 0)

	nama := readInput("Masukkan Nama Baru: ")
	umurStr := readInput("Masukkan Umur Baru: ")
	bountyStr := readInput("Masukkan Bounty Baru: ")
	hobi := readInput("Masukkan Hobi Baru: ")

	umur, err1 := strconv.Atoi(umurStr)
	bounty, err2 := strconv.Atoi(bountyStr)

	if err1 != nil || err2 != nil {
		fmt.Println("⚠️ Umur dan bounty harus berupa angka!")
		return
	}

	characters[index] = Character{Nama: nama, Umur: umur, Bounty: bounty, Hobi: hobi}
	fmt.Println("\n✅ Karakter berhasil diperbarui!")
}

// Fungsi untuk menghapus karakter berdasarkan indeks
func deleteCharacter() {
	getAllCharacters()
	indexStr := readInput("\nMasukkan nomor karakter yang ingin dihapus: ")
	index, err := strconv.Atoi(indexStr)

	if err != nil || index < 1 || index > len(characters) {
		fmt.Println("⚠️ Indeks tidak valid.")
		return
	}

	index-- // Sesuaikan dengan indeks slice (mulai dari 0)
	characters = append(characters[:index], characters[index+1:]...)
	fmt.Println("\n🗑️ Karakter berhasil dihapus!")
}

// Fungsi utama untuk menu interaktif
func main() {
	for {
		fmt.Println("\n==== MENU CRUD KARAKTER ====")
		fmt.Println("1. Tambah Karakter")
		fmt.Println("2. Lihat Semua Karakter")
		fmt.Println("3. Perbarui Karakter")
		fmt.Println("4. Hapus Karakter")
		fmt.Println("5. Keluar")
		pilihanStr := readInput("Pilih menu (1-5): ")

		pilihan, err := strconv.Atoi(pilihanStr)
		if err != nil {
			fmt.Println("⚠️ Pilihan harus berupa angka!")
			continue
		}

		switch pilihan {
		case 1:
			addCharacter()
		case 2:
			getAllCharacters()
		case 3:
			updateCharacter()
		case 4:
			deleteCharacter()
		case 5:
			fmt.Println("🚀 Program selesai. Terima kasih!")
			return
		default:
			fmt.Println("⚠️ Pilihan tidak valid, coba lagi.")
		}
	}
}
