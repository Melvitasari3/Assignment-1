package main

import (
	"fmt"
	"os"
	"strconv"
)

type biodata struct {
	absen     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	Program := os.Args
	index, _ := strconv.Atoi(Program[1])

	var student = []biodata{
		{absen: 1, nama: "Melvita", alamat: "Dusun 1 Desa Lubuk Saban Kecamatan Pantai Cermin Kabupaten Serdang Bedagai Provinsi Sumatera Utara", pekerjaan: "Mahasiswa", alasan: "Ingin belajar pemograman golang dan mendapatkan lebih banyak teman"},
		{absen: 2, nama: "Yola", alamat: "Dusun 5 Desa wonosari Kecamatan Pantai Cermin Kabupaten Serdang Bedagai Provinsi Sumatera Utara", pekerjaan: "Mahasiswa", alasan: "Ingin belajar pemograman golang "},
		{absen: 3, nama: "Rahma", alamat: "Dusun 1 Desa jagung Kecamatan Pantai Cermin Kabupaten Serdang Bedagai Provinsi Sumatera Utara", pekerjaan: "Mahasiswa", alasan: "Ingin mendapatkan lebih banyak teman"},
	}
	getStudent(student, index)
}

func getStudent(student []biodata, absen int) {
	var index int
	for i := 0; i < len(student); i++ {
		if student[i].absen == absen {
			index = i
			goto suc
		} else if i == (len(student) - 1) {
			goto empty
		}
	}

suc:
	fmt.Printf("Nama : %v\n", student[index].nama)
	fmt.Printf("Alamat : %v\n", student[index].alamat)
	fmt.Printf("Pekerjaan : %v\n", student[index].pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang : %v\n", student[index].alasan)
	return

empty:
	fmt.Println("Nomor absen yang anda cari tidak terdaftar")
	return
}
