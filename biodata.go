package main

import (
	"fmt"
	"os"
	"strconv"
)

type namaSiswa struct {
	absensi   int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	Program := os.Args
	idSiswa, _ := strconv.Atoi(Program[1])
	getDataKelas(idSiswa)
	// fmt.Println(namaProgram)
}

var dataKelas []namaSiswa

func getDataKelas(noabsen int) {
	dataKelas = []namaSiswa{
		{absensi: 1, nama: "Melvita", alamat: "Dusun 1 Desa Lubuk Saban Kecamatan Pantai Cermin Kabupaten Serdang Bedagai Provinsi Sumatera Utara", pekerjaan: "Mahasiswa", alasan: "Ingin belajar pemograman golang dan mendapatkan lebih banyak teman"},
		{absensi: 2, nama: "Yola", alamat: "Dusun 5 Desa wonosari Kecamatan Pantai Cermin Kabupaten Serdang Bedagai Provinsi Sumatera Utara", pekerjaan: "Mahasiswa", alasan: "Ingin belajar pemograman golang "},
		{absensi: 3, nama: "Rahma", alamat: "Dusun 1 Desa jagung Kecamatan Pantai Cermin Kabupaten Serdang Bedagai Provinsi Sumatera Utara", pekerjaan: "Mahasiswa", alasan: "Ingin mendapatkan lebih banyak teman"},
	}

	for _, siswa := range dataKelas {

		if siswa.absensi == noabsen {
			fmt.Println("Nama : ", siswa.nama)
			fmt.Println("Alamat : ", siswa.alamat)
			fmt.Println("Pekerjaan : ", siswa.pekerjaan)
			fmt.Println("Alasan memilih kelas Golang : ", siswa.alasan)

		}
	}
}
