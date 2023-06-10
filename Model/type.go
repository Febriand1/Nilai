package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NPM          int                `bson:"npm,omitempty" json:"npm,omitempty"`
	Phone_Number string             `bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
}

type Matakuliah struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_MK  string             `bson:"nama_mk,omitempty" json:"nama_mk,omitempty"`
	SKS      int                `bson:"sks,omitempty" json:"sks,omitempty"`
	Jadwal   Waktu              `bson:"jadwal,omitempty" json:"jadwal,omitempty"`
	Pengampu Dosen              `bson:"pengampu,omitempty" json:"pengampu,omitempty"`
}

type Waktu struct {
	Jam_Masuk  string   `bson:"jammasuk,omitempty" json:"jammasuk,omitempty"`
	Jam_Keluar string   `bson:"jamkeluar,omitempty" json:"jamkeluar,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
}

type Dosen struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Dosen    string             `bson:"namadosen,omitempty" json:"namadosen,omitempty"`
	NIK           string             `bson:"nik,omitempty" json:"nik,omitempty"`
	Phone_NumberD string             `bson:"phonenumberd,omitempty" json:"phonenumberd,omitempty"`
}

type Tugas struct {
	Tugas1 int `bson:"tugas1,omitempty" json:"tugas1,omitempty"`
	Tugas2 int `bson:"tugas2,omitempty" json:"tugas2,omitempty"`
	Tugas3 int `bson:"tugas3,omitempty" json:"tugas3,omitempty"`
	Tugas4 int `bson:"tugas4,omitempty" json:"tugas4,omitempty"`
	Tugas5 int `bson:"tugas5,omitempty" json:"tugas5,omitempty"`
}

type Nilai struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	All_Tugas Tugas              `bson:"alltugas,omitempty" json:"alltugas,omitempty"`
	UTS       int                `bson:"uts,omitempty" json:"uts,omitempty"`
	UAS       int                `bson:"uas,omitempty" json:"uas,omitempty"`
	Grade     Grade              `bson:"grade,omitempty" json:"grade,omitempty"`
	Kategori  Matakuliah         `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Absensi   Presensi           `bson:"absensi,omitempty" json:"absensi,omitempty"`
}

type Grade struct {
	Nama_Grade string `bson:"namagrade,omitempty" json:"namagrade,omitempty"`
	Skala      string `bson:"skala,omitempty" json:"skala,omitempty"`
}

type Presensi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jumlah_Kehadiran int                `bson:"jumlahkehadiran,omitempty" json:"jumlahkehadiran,omitempty"`
	Biodata          Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
