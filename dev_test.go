package NPM

import (
	"fmt"
	model "github.com/Febriand1/Nilai/Model"
	"github.com/Febriand1/Nilai/Module"
	"testing"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "budiman"
	npm := 121395
	phone_number := "083543242546"

	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, phone_number)
	fmt.Println(hasil)
}

func TestInsertWaktu(t *testing.T) {
	jam_masuk := "07.00"
	jam_keluar := "10.00"
	hari := []string{"sabtu"}

	hasil := module.InsertWaktu(module.MongoConn, "waktu", jam_masuk, jam_keluar, hari)
	fmt.Println(hasil)
}

func TestInsertDosen(t *testing.T) {
	nama_dosen := "indra riksa herlambang"
	nik := "123.123.123"
	phonenumberd := "08123422542"

	hasil := module.InsertDosen(module.MongoConn, "dosen", nama_dosen, nik, phonenumberd)
	fmt.Println(hasil)
}

func TestInsertMatakuliah(t *testing.T) {
	nama_mk := "algoritma"
	sks := 2
	jadwal := model.Waktu{
		Jam_Masuk:  "07.00",
		Jam_Keluar: "10.00",
		Hari:       []string{"sabtu"},
	}
	pengampu := model.Dosen{
		Nama_Dosen:    "indra riksa herlambang",
		NIK:           "123.123.123",
		Phone_NumberD: "08123422542",
	}
	hasil := module.InsertMatakuliah(module.MongoConn, "matakuliah", nama_mk, sks, jadwal, pengampu)
	fmt.Println(hasil)
}

func TestInsertTugas(t *testing.T) {
	tugas1 := 100
	tugas2 := 80
	tugas3 := 80
	tugas4 := 80
	tugas5 := 100
	hasil := module.InsertTugas(module.MongoConn, "tugas", tugas1, tugas2, tugas3, tugas4, tugas5)
	fmt.Println(hasil)
}

func TestInsertNilai(t *testing.T) {
	all_tugas := model.Tugas{
		Tugas1: 100,
		Tugas2: 80,
		Tugas3: 80,
		Tugas4: 80,
		Tugas5: 100,
	}
	uts := 70
	uas := 70
	grade := model.Grade{
		Nama_Grade: "A",
		Skala:      "95-100",
	}
	kategori := model.Matakuliah{
		Nama_MK: "algoritma",
		SKS:     2,
		Jadwal: model.Waktu{
			Jam_Masuk:  "07.00",
			Jam_Keluar: "10.00",
			Hari:       []string{"sabtu"},
		},
		Pengampu: model.Dosen{
			Nama_Dosen:    "indra riksa herlambang",
			NIK:           "123.123.123",
			Phone_NumberD: "08123422542",
		},
	}
	biodata := model.Mahasiswa{
		Nama:         "budiman",
		NPM:          121395,
		Phone_Number: "083543242546",
	}
	hasil := module.InsertNilai(module.MongoConn, "nilai", all_tugas, uts, uas, grade, kategori, biodata)
	fmt.Println(hasil)
}

func TestInsertGrade(t *testing.T) {
	namagrade := "A"
	skala := "95-100"
	hasil := module.InsertGrade(module.MongoConn, "grade", namagrade, skala)
	fmt.Println(hasil)
}

func TestInsertPresensi(t *testing.T) {
	jumlahkehadiran := 7
	biodata := model.Mahasiswa{
		Nama:         "budiman",
		NPM:          121395,
		Phone_Number: "083543242546",
	}
	hasil := module.InsertPresensi(module.MongoConn, "presensi", jumlahkehadiran, biodata)
	fmt.Println(hasil)
}

func TestGetPresensiFromMahasiswa(t *testing.T) {
	nama := "budiman"
	identitas := module.GetPresensiFromMahasiswa(nama, module.MongoConn, "presensi")
	fmt.Println(identitas)
}

func TestGetUASFromMahasiswa(t *testing.T) {
	npm := 121395
	nilai := module.GetUASFromMahasiswa(npm, module.MongoConn, "nilai")
	fmt.Println(nilai)
}

func TestGetGradeFromMahasiswa(t *testing.T) {
	npm := 121395
	nilai := module.GetGradeFromMahasiswa(npm, module.MongoConn, "nilai")
	fmt.Println(nilai)
}

func TestGetMatakuliahFromJadwal(t *testing.T) {
	jam_masuk := "07.00"
	jam := module.GetMatakuliahFromJadwal(jam_masuk, module.MongoConn, "matakuliah")
	fmt.Println(jam)
}

func TestGetMatakuliahFromNIK(t *testing.T) {
	nik := "123.123.123"
	matkul := module.GetMatakuliahFromNIK(nik, module.MongoConn, "matakuliah")
	fmt.Println(matkul)
}

func TestGetNilaiFromMatakuliah(t *testing.T) {
	nama_mk := "algoritma"
	data := module.GetNilaiFromMatakuliah(nama_mk, module.MongoConn, "nilai")
	fmt.Println(data)
}

func TestGetNilaiFromNamaMahasiswa(t *testing.T) {
	nama := "budiman"
	data := module.GetNilaiFromNamaMahasiswa(nama, module.MongoConn, "nilai")
	fmt.Println(data)
}

func TestGetAllNilaiFromNamaMahasiswa(t *testing.T) {
	nama := "budiman"
	data1 := module.GetAllNilaiFromNamaMahasiswa(nama, module.MongoConn, "nilai")
	fmt.Println(data1)
}
