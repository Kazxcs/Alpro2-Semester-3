package main

import (
	"fmt"
	"strings"
)

const MAX_DATA = 20

type Pabrikan struct {
	ID           int
	Nama         string
	Negara       string
	TahunBerdiri int
	Alamat       string
}

type Mobil struct {
	ID             int
	IDPabrikan     int
	Nama           string
	Model          string
	TahunProduksi  int
	Harga          float64
	KapasitasMesin float64
	WarnaTersedia  [3]string
	Rating         float64
	Favorit        bool
	Penjualan      int
}

type FilterMobil struct {
	NamaPabrikan    string
	NamaMobil       string
	HargaMin        float64
	HargaMax        float64
	UrutBerdasarkan string
	UrutNaikTurun   string
}

type Laporan struct {
	LaporanPabrikan      [MAX_DATA]LaporanPabrikan
	LaporanMobil         [MAX_DATA]LaporanMobil
	PabrikanTeratas      [3]PabrikanTeratas
	MobilTeratas         [3]MobilTeratas
	TotalPabrikan        int
	TotalMobil           int
	RataMobilPerPabrikan float64
	TotalPendapatan      float64
}

type LaporanPabrikan struct {
	NamaPabrikan string
	JumlahMobil  int
}

type LaporanMobil struct {
	Nama          string
	TahunProduksi int
	Pabrikan      string
	Harga         float64
}

type PabrikanTeratas struct {
	Nama           string
	TotalPenjualan int
}

type MobilTeratas struct {
	Nama   string
	Rating float64
}

type LaporanPenjualan struct {
	NamaMobil        string
	UnitTerjual      int
	Pendapatan       float64
	PendapatanDealer float64
}

var pabrikan [MAX_DATA]Pabrikan
var mobil [MAX_DATA]Mobil
var jumlahPabrikan, jumlahMobil int

func login() bool {
	var username, password string
	for {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println(" ")
		fmt.Printf("%-10s%s\n", "", "Selamat datang di Dealer Mobil Kelompok 1")
		fmt.Println(" ")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println(" ")
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username == "Kelompok1Alpro" && password == "Semester3" {
			fmt.Println(" ")
			fmt.Println("------------------------ Login berhasil! ------------------------")
			return true
		} else {
			fmt.Println(" ")
			fmt.Println("Alert: ")
			fmt.Println("------- Username atau password salah! Silakan coba lagi! --------")
			fmt.Println(" ")
		}
	}
}

func tambahPabrikan(data *[MAX_DATA]Pabrikan, jumlah *int) {
	if *jumlah >= MAX_DATA {
		fmt.Println("Data sudah penuh, tidak bisa menambah pabrikan baru.")
		return
	}

	var pabrikan Pabrikan
	fmt.Println("\n--- Tambah Pabrikan ---")
	fmt.Print("Masukkan ID Pabrikan: ")
	fmt.Scanln(&pabrikan.ID)
	fmt.Print("Masukkan Nama Pabrikan: ")
	fmt.Scanln(&pabrikan.Nama)
	fmt.Print("Masukkan Negara Pabrikan: ")
	fmt.Scanln(&pabrikan.Negara)
	fmt.Print("Masukkan Tahun Berdiri: ")
	fmt.Scanln(&pabrikan.TahunBerdiri)
	fmt.Print("Masukkan Alamat Pabrikan: ")
	fmt.Scanln(&pabrikan.Alamat)

	data[*jumlah] = pabrikan
	*jumlah++

	fmt.Println("Pabrikan berhasil ditambahkan!")
}

func updatePabrikan(data *[MAX_DATA]Pabrikan, jumlah int) {
	var id int
	fmt.Println("\n--- Update Pabrikan ---")
	fmt.Print("Masukkan ID Pabrikan yang ingin diupdate: ")
	fmt.Scanln(&id)

	for i := 0; i < jumlah; i++ {
		if data[i].ID == id {
			var pabrikan Pabrikan
			fmt.Print("Masukkan Nama Pabrikan: ")
			fmt.Scanln(&pabrikan.Nama)
			fmt.Print("Masukkan Negara Pabrikan: ")
			fmt.Scanln(&pabrikan.Negara)
			fmt.Print("Masukkan Tahun Berdiri: ")
			fmt.Scanln(&pabrikan.TahunBerdiri)
			fmt.Print("Masukkan Alamat Pabrikan: ")
			fmt.Scanln(&pabrikan.Alamat)

			data[i] = pabrikan
			fmt.Println("Pabrikan berhasil diupdate!")
			return
		}
	}

	fmt.Println("Pabrikan dengan ID tersebut tidak ditemukan.")
}

func deletePabrikan(data *[MAX_DATA]Pabrikan, jumlah *int) {
	var id int
	fmt.Println("\n--- Hapus Pabrikan ---")
	fmt.Print("Masukkan ID Pabrikan yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i := 0; i < *jumlah; i++ {
		if data[i].ID == id {
			for j := i; j < *jumlah-1; j++ {
				data[j] = data[j+1]
			}
			*jumlah--
			fmt.Println("Pabrikan berhasil dihapus!")
			return
		}
	}

	fmt.Println("Pabrikan dengan ID tersebut tidak ditemukan.")
}

func pengelolaanPabrikan(data *[MAX_DATA]Pabrikan, jumlah *int) {
	var pilihan int
	for {
		fmt.Println("\n--- Pengelolaan Data Pabrikan ---")
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Update Pabrikan")
		fmt.Println("3. Hapus Pabrikan")
		fmt.Println("4. Kembali ke Menu Utama")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPabrikan(data, jumlah)
		case 2:
			updatePabrikan(data, *jumlah)
		case 3:
			deletePabrikan(data, jumlah)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahMobil(data *[MAX_DATA]Mobil, jumlah *int) {
	if *jumlah >= MAX_DATA {
		fmt.Println("Data sudah penuh, tidak bisa menambah mobil baru.")
		return
	}

	var mobil Mobil
	fmt.Println("\n--- Tambah Mobil ---")
	fmt.Print("Masukkan ID Mobil: ")
	fmt.Scanln(&mobil.ID)
	fmt.Print("Masukkan ID Pabrikan: ")
	fmt.Scanln(&mobil.IDPabrikan)
	fmt.Print("Masukkan Nama Mobil: ")
	fmt.Scanln(&mobil.Nama)
	fmt.Print("Masukkan Model Mobil: ")
	fmt.Scanln(&mobil.Model)
	fmt.Print("Masukkan Tahun Produksi: ")
	fmt.Scanln(&mobil.TahunProduksi)
	fmt.Print("Masukkan Harga Mobil: ")
	fmt.Scanln(&mobil.Harga)
	fmt.Print("Masukkan Kapasitas Mesin (dalam liter): ")
	fmt.Scanln(&mobil.KapasitasMesin)
	fmt.Print("Masukkan Warna Tersedia (3 warna): ")
	for i := 0; i < 3; i++ {
		fmt.Print("Warna ke-", i+1, ": ")
		fmt.Scanln(&mobil.WarnaTersedia[i])
	}
	fmt.Print("Masukkan Rating Mobil (0-5): ")
	fmt.Scanln(&mobil.Rating)
	fmt.Print("Apakah mobil ini favorit? (true/false): ")
	fmt.Scanln(&mobil.Favorit)
	fmt.Print("Masukkan Jumlah Penjualan: ")
	fmt.Scanln(&mobil.Penjualan)

	data[*jumlah] = mobil
	*jumlah++

	fmt.Println("Mobil berhasil ditambahkan!")
}

func updateMobil(data *[MAX_DATA]Mobil, jumlah int) {
	var id int
	fmt.Println("\n--- Update Mobil ---")
	fmt.Print("Masukkan ID Mobil yang ingin diupdate: ")
	fmt.Scanln(&id)

	for i := 0; i < jumlah; i++ {
		if data[i].ID == id {
			var mobil Mobil
			fmt.Print("Masukkan ID Pabrikan: ")
			fmt.Scanln(&mobil.IDPabrikan)
			fmt.Print("Masukkan Nama Mobil: ")
			fmt.Scanln(&mobil.Nama)
			fmt.Print("Masukkan Model Mobil: ")
			fmt.Scanln(&mobil.Model)
			fmt.Print("Masukkan Tahun Produksi: ")
			fmt.Scanln(&mobil.TahunProduksi)
			fmt.Print("Masukkan Harga Mobil: ")
			fmt.Scanln(&mobil.Harga)
			fmt.Print("Masukkan Kapasitas Mesin (dalam liter): ")
			fmt.Scanln(&mobil.KapasitasMesin)
			fmt.Print("Masukkan Warna Tersedia (3 warna): ")
			for i := 0; i < 3; i++ {
				fmt.Print("Warna ke-", i+1, ": ")
				fmt.Scanln(&mobil.WarnaTersedia[i])
			}
			fmt.Print("Masukkan Rating Mobil (0-5): ")
			fmt.Scanln(&mobil.Rating)
			fmt.Print("Apakah mobil ini favorit? (true/false): ")
			fmt.Scanln(&mobil.Favorit)
			fmt.Print("Masukkan Jumlah Penjualan: ")
			fmt.Scanln(&mobil.Penjualan)

			data[i] = mobil
			fmt.Println("Mobil berhasil diupdate!")
			return
		}
	}

	fmt.Println("Mobil dengan ID tersebut tidak ditemukan.")
}

func deleteMobil(data *[MAX_DATA]Mobil, jumlah *int) {
	var id int
	fmt.Println("\n--- Hapus Mobil ---")
	fmt.Print("Masukkan ID Mobil yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i := 0; i < *jumlah; i++ {
		if data[i].ID == id {
			for j := i; j < *jumlah-1; j++ {
				data[j] = data[j+1]
			}
			*jumlah--
			fmt.Println("Mobil berhasil dihapus!")
			return
		}
	}

	fmt.Println("Mobil dengan ID tersebut tidak ditemukan.")
}

func pengelolaanMobil(data *[MAX_DATA]Mobil, jumlah *int) {
	var pilihan int
	for {
		fmt.Println("\n--- Pengelolaan Data Mobil ---")
		fmt.Println("1. Tambah Mobil")
		fmt.Println("2. Update Mobil")
		fmt.Println("3. Hapus Mobil")
		fmt.Println("4. Kembali ke Menu Utama")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahMobil(data, jumlah)
		case 2:
			updateMobil(data, *jumlah)
		case 3:
			deleteMobil(data, jumlah)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func pencarianDanFilterMobil(data [MAX_DATA]Mobil, jumlah int) {
	var pilihan int
	var filter FilterMobil
	fmt.Println("\n--- Pencarian dan Filter Mobil ---")
	fmt.Println("1. Filter Berdasarkan Harga")
	fmt.Println("2. Filter Berdasarkan Tahun Produksi")
	fmt.Println("3. Filter Berdasarkan Nama Mobil")
	fmt.Println("4. Kembali ke Menu Utama")

	fmt.Print("Pilih menu: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan Harga Minimum: ")
		fmt.Scanln(&filter.HargaMin)
		fmt.Print("Masukkan Harga Maksimum: ")
		fmt.Scanln(&filter.HargaMax)
		for i := 0; i < jumlah; i++ {
			if data[i].Harga >= filter.HargaMin && data[i].Harga <= filter.HargaMax {
				fmt.Printf("%+v\n", data[i])
			}
		}
	case 2:
		fmt.Print("Masukkan Tahun Produksi Minimum: ")
		var tahunMin int
		fmt.Scanln(&tahunMin)
		for i := 0; i < jumlah; i++ {
			if data[i].TahunProduksi >= tahunMin {
				fmt.Printf("%+v\n", data[i])
			}
		}
	case 3:
		fmt.Print("Masukkan Nama Mobil: ")
		fmt.Scanln(&filter.NamaMobil)
		for i := 0; i < jumlah; i++ {
			if strings.Contains(data[i].Nama, filter.NamaMobil) {
				fmt.Printf("%+v\n", data[i])
			}
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func laporanDanStatistik(data [MAX_DATA]Mobil, jumlah int) {
	var pilihan int
	fmt.Println("\n--- Laporan dan Statistik ---")
	fmt.Println("1. Mobil dengan Penjualan Tertinggi")
	fmt.Println("2. Mobil Favorit")
	fmt.Println("3. Rata-rata Harga Mobil")
	fmt.Println("4. Kembali ke Menu Utama")

	fmt.Print("Pilih menu: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		var maxPenjualan int
		var mobilTertinggi Mobil
		for i := 0; i < jumlah; i++ {
			if data[i].Penjualan > maxPenjualan {
				maxPenjualan = data[i].Penjualan
				mobilTertinggi = data[i]
			}
		}
		fmt.Println("Mobil dengan Penjualan Tertinggi: ", mobilTertinggi)
	case 2:
		fmt.Println("Mobil Favorit:")
		for i := 0; i < jumlah; i++ {
			if data[i].Favorit {
				fmt.Printf("%+v\n", data[i])
			}
		}
	case 3:
		var totalHarga float64
		for i := 0; i < jumlah; i++ {
			totalHarga += data[i].Harga
		}
		fmt.Printf("Rata-rata Harga Mobil: %.2f\n", totalHarga/float64(jumlah))
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func fiturFavoritMobil(data *[MAX_DATA]Mobil, jumlah int) {
	var id int
	var pilihan int
	fmt.Println("\n--- Fitur Favorit untuk Mobil ---")
	fmt.Println("1. Tandai Mobil sebagai Favorit")
	fmt.Println("2. Hapus Mobil dari Favorit")
	fmt.Println("3. Lihat Mobil Favorit")
	fmt.Println("4. Kembali ke Menu Utama")

	fmt.Print("Pilih menu: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("Masukkan ID Mobil yang ingin ditandai sebagai favorit: ")
		fmt.Scanln(&id)
		for i := 0; i < jumlah; i++ {
			if data[i].ID == id {
				data[i].Favorit = true
				fmt.Println("Mobil berhasil ditandai sebagai favorit!")
				return
			}
		}
		fmt.Println("Mobil dengan ID tersebut tidak ditemukan.")
	case 2:
		fmt.Print("Masukkan ID Mobil yang ingin dihapus dari favorit: ")
		fmt.Scanln(&id)
		for i := 0; i < jumlah; i++ {
			if data[i].ID == id {
				data[i].Favorit = false
				fmt.Println("Mobil berhasil dihapus dari favorit!")
				return
			}
		}
		fmt.Println("Mobil dengan ID tersebut tidak ditemukan.")
	case 3:
		fmt.Println("Mobil Favorit:")
		for i := 0; i < jumlah; i++ {
			if data[i].Favorit {
				fmt.Printf("%+v\n", data[i])
			}
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func tampilkanPabrikan(data [MAX_DATA]Pabrikan, jumlah int) {
	fmt.Println("\nDaftar Pabrikan:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("ID: %d, Nama: %s, Negara: %s, Tahun Berdiri: %d, Alamat: %s\n",
			data[i].ID, data[i].Nama, data[i].Negara, data[i].TahunBerdiri, data[i].Alamat)
	}
}

func tampilkanMobil(data [MAX_DATA]Mobil, jumlah int) {
	fmt.Println("\nDaftar Mobil:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("ID: %d, Nama: %s, Model: %s, Tahun Produksi: %d, Harga: %.2f, Kapasitas Mesin: %.1f, Rating: %.1f, Favorit: %t, Penjualan: %d\n",
			data[i].ID, data[i].Nama, data[i].Model, data[i].TahunProduksi, data[i].Harga, data[i].KapasitasMesin, data[i].Rating, data[i].Favorit, data[i].Penjualan)
	}
}

func main() {
	pabrikan[0] = Pabrikan{ID: 1, Nama: "Toyota", Negara: "Jepang", TahunBerdiri: 1937, Alamat: "Toyota City, Japan"}
	pabrikan[1] = Pabrikan{ID: 2, Nama: "Honda", Negara: "Jepang", TahunBerdiri: 1948, Alamat: "Tokyo, Japan"}
	jumlahPabrikan = 2

	mobil[0] = Mobil{ID: 1, IDPabrikan: 1, Nama: "Toyota Camry", Model: "Sedan", TahunProduksi: 2023, Harga: 500000000, KapasitasMesin: 2.5, WarnaTersedia: [3]string{"Hitam", "Putih", "Perak"}, Rating: 4.5, Favorit: true, Penjualan: 100}
	mobil[1] = Mobil{ID: 2, IDPabrikan: 2, Nama: "Honda Civic", Model: "Sedan", TahunProduksi: 2023, Harga: 400000000, KapasitasMesin: 1.8, WarnaTersedia: [3]string{"Merah", "Putih", "Biru"}, Rating: 4.7, Favorit: false, Penjualan: 120}
	jumlahMobil = 2

	login()
	var pilihan int
	for {
		fmt.Println(" ")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println(" ")
		fmt.Printf("%-13s%s\n", "", "Menu Dealer Mobil Kelompok 1")
		fmt.Println(" ")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println(" ")
		fmt.Println("1. Pengelolaan Data Pabrikan")
		fmt.Println("2. Pengelolaan Data Mobil")
		fmt.Println("3. Tampilkan Data Pabrikan dan Mobil")
		fmt.Println("4. Pencarian dan Filter Data")
		fmt.Println("5. Laporan dan Statistik")
		fmt.Println("6. Fitur Favorit untuk Mobil ")
		fmt.Println("7. Sistem Rating untuk Mobil ")
		fmt.Println("8. Filter Mobil Berdasarkan Harga")
		fmt.Println("9. Penghitung Total Penjualan Mobil")
		fmt.Println("10. Keluar")
		fmt.Println(" ")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println(" ")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)
		fmt.Println(" ")
		fmt.Println("-----------------------------------------------------------------")

		switch pilihan {
		case 1:
			pengelolaanPabrikan(&pabrikan, &jumlahPabrikan)
		case 2:
			pengelolaanMobil(&mobil, &jumlahMobil)
		case 3:
			tampilkanPabrikan(pabrikan, jumlahPabrikan)
			tampilkanMobil(mobil, jumlahMobil)
		case 4:
			pencarianDanFilterMobil(mobil, jumlahMobil)
		case 5:
			laporanDanStatistik(mobil, jumlahMobil)
		case 6:
			fiturFavoritMobil(&mobil, jumlahMobil)
		case 7:
		case 8:
		case 9:
		case 10:
			fmt.Println(" ")
			fmt.Println("Terima kasih telah menggunakan aplikasi Dealer Mobil Kelompok 1.")
			fmt.Printf("%-5s%s\n", "", "Semoga aplikasi ini dapat membantu pekerjaan Anda.")
			fmt.Println(" ")
			fmt.Printf("%-10s%s\n", "", "Sampai jumpa lagi dan selamat bekerja!")
			fmt.Println(" ")
			fmt.Println("-----------------------------------------------------------------")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
