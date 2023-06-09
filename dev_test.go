package NPM

import (
	"fmt"
	"testing"

	model "github.com/Febriand1/Nilai/Model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	module "github.com/Febriand1/Nilai/Module"
)

func TestInsertMahasiswa(t *testing.T) {
	nama := "jeki"
	npm := 1294653
	phonenumber := "081234543421"

	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, npm, phonenumber)
	fmt.Println(hasil)
}

func TestInsertWaktu(t *testing.T) {
	jammasuk := "06.00"
	jamkeluar := "10.00"
	hari := []string{"kamis, jumat, sabtu"}

	hasil := module.InsertWaktu(module.MongoConn, "waktu", jammasuk, jamkeluar, hari)
	fmt.Println(hasil)
}

func TestInsertDosen(t *testing.T) {
	namadosen := "mark"
	nik := "808.808.808"
	phonenumberd := "084532524566"

	hasil := module.InsertDosen(module.MongoConn, "dosen", namadosen, nik, phonenumberd)
	fmt.Println(hasil)
}

func TestInsertMatakuliah(t *testing.T) {
	namamk := ""
	sks := 4
	jadwal := model.Waktu{
		Jam_Masuk:  "06.00",
		Jam_Keluar: "10.00",
		Hari:       []string{"kamis, jumat, sabtu"},
	}
	pengampu := model.Dosen{
		Nama_Dosen:    "mark",
		NIK:           "808.808.808",
		Phone_NumberD: "084532524566",
	}
	hasil := module.InsertMatakuliah(module.MongoConn, "matakuliah", namamk, sks, jadwal, pengampu)
	fmt.Println(hasil)
}

func TestInsertTugas(t *testing.T) {
	tugas1 := 90
	tugas2 := 100
	tugas3 := 80
	tugas4 := 70
	tugas5 := 70
	hasil := module.InsertTugas(module.MongoConn, "tugas", tugas1, tugas2, tugas3, tugas4, tugas5)
	fmt.Println(hasil)
}

func TestInsertNilai(t *testing.T) {
	alltugas := model.Tugas{
		Tugas1: 90,
		Tugas2: 100,
		Tugas3: 80,
		Tugas4: 70,
		Tugas5: 70,
	}
	uts := 70
	uas := 80
	grade := model.Grade{
		Nama_Grade: "B",
		Skala:      "80-95",
	}
	kategori := model.Matakuliah{
		Nama_MK: "",
		SKS:     4,
		Jadwal: model.Waktu{
			Jam_Masuk:  "06.00",
			Jam_Keluar: "10.00",
			Hari:       []string{"kamis, jumat, sabtu"},
		},
		Pengampu: model.Dosen{
			Nama_Dosen:    "mark",
			NIK:           "808.808.808",
			Phone_NumberD: "084532524566",
		},
	}
	absensi := model.Presensi{
		Jumlah_Kehadiran: 7,
		Biodata: model.Mahasiswa{
			Nama:         "jeki",
			NPM:          1294653,
			Phone_Number: "081234543421",
		},
	}

	hasil := module.InsertNilai(module.MongoConn, "nilai", alltugas, uts, uas, grade, kategori, absensi)
	fmt.Println(hasil)
}

func TestInsertGrade(t *testing.T) {
	namagrade := "B"
	skala := "80-95"
	hasil := module.InsertGrade(module.MongoConn, "grade", namagrade, skala)
	fmt.Println(hasil)
}

func TestInsertPresensi(t *testing.T) {
	jumlahkehadiran := 7
	biodata := model.Mahasiswa{
		Nama:         "jeki",
		NPM:          1294653,
		Phone_Number: "081234543421",
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
	jammasuk := ""
	jam := module.GetMatakuliahFromJadwal(jammasuk, module.MongoConn, "matakuliah")
	fmt.Println(jam)
}

func TestGetMatakuliahFromNIK(t *testing.T) {
	nik := "123.123.123"
	matkul := module.GetMatakuliahFromNIK(nik, module.MongoConn, "matakuliah")
	fmt.Println(matkul)
}

func TestGetNilaiFromMatakuliah(t *testing.T) {
	namamk := "algoritma"
	data := module.GetNilaiFromMatakuliah(namamk, module.MongoConn, "nilai")
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

func TestGetAll(t *testing.T) {
	data := module.GetAllNilai(module.MongoConn, "nilai")
	fmt.Println(data)
}


func TestGetNilaiFromID(t *testing.T) {
	id := "6426d338c606ea58bc4f8ebd"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	mhs, err := module.GetNilaiFromID(objectID, module.MongoConn, "nilai")
	if err != nil {
		t.Fatalf("error calling GetNilaiFromID: %v", err)
	}
	fmt.Println(mhs)
}

func TestDeleteMahasiswaByID(t *testing.T) {
	id := "6426d338c606ea58bc4f8ebd" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteNilaiByID(objectID, module.MongoConn, "nilai")
	if err != nil {
		t.Fatalf("error calling DeleteNilaiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetNilaiFromID(objectID, module.MongoConn, "nilai")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
