package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	var allStudent = []Student{
		{name: "Nuraeni", address: "Jl Hartaco Indah", job: "Sofware Engineer", reason: "Ingin mempelajari bahasa pemrograman baru yang belum pernah ia pelajari yaitu golang"},
		{name: "Taufik", address: "Jl Masolo", job: "Programmer", reason: "Hanya penasaran saja"},
		{name: "Fasya", address: "Jl Pongtiku", job: "IT Support", reason: "Merasa tertantang belajar Golang"},
	}
	// for _, v:= range allStudent {
	// 	fmt.Printf("%+v\n", v)
	// }

	if len(os.Args) < 2 {
		fmt.Println("Gunakan perintah: go run biodata.go [no_absen_siswa]")
		return 
	}

	//mengambil nomor absen dari masukan user go run biodata.go [no_absen_siswa]
	numberAbsen, err := strconv.Atoi(os.Args[1]) //strconv.Atoi untuk convert dari string ke int
	
	_ = err
	//konversi numberAbsen sesuai index karena dimulai dari 0 maka dikurangi 1
	index := numberAbsen - 1

	if index < 0 || index >= len(allStudent){
		fmt.Println("Nomor absen invalid atau tidak tersedia")
		return
	}

	student := allStudent[index]

	fmt.Println("Nama : ", student.name)
	fmt.Println("Alamat : ", student.address)
	fmt.Println("Pekerjaan : ", student.job)
	fmt.Println("Alasan memilih kelas Golang : ", student.reason)


	



}