package module

import (
	"context"
	"fmt"
	"github.com/Febriand1/Nilai/Model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "db_penilaian",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMahasiswa(db *mongo.Database, col string, nama string, npm int, phone_number string) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.NPM = npm
	mahasiswa.Phone_Number = phone_number
	return InsertOneDoc(db, col, mahasiswa)
}

func InsertMatakuliah(db *mongo.Database, col string, nama_mk string, sks int, jadwal model.Waktu, pengampu model.Dosen) (InsertedID interface{}) {
	var matakuliah model.Matakuliah
	matakuliah.Nama_MK = nama_mk
	matakuliah.SKS = sks
	matakuliah.Jadwal = jadwal
	matakuliah.Pengampu = pengampu
	return InsertOneDoc(db, col, matakuliah)
}

func InsertWaktu(db *mongo.Database, col string, jam_masuk string, jam_keluar string, hari []string) (InsertedID interface{}) {
	var waktu model.Waktu
	waktu.Jam_Masuk = jam_masuk
	waktu.Jam_Keluar = jam_keluar
	waktu.Hari = hari
	return InsertOneDoc(db, col, waktu)
}

func InsertDosen(db *mongo.Database, col string, nama_dosen string, nik string, phonenumberd string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.Nama_Dosen = nama_dosen
	dosen.NIK = nik
	dosen.Phone_NumberD = phonenumberd
	return InsertOneDoc(db, col, dosen)
}

func InsertNilai(db *mongo.Database, col string, all_tugas model.Tugas, uts int, uas int, grade model.Grade, kategori model.Matakuliah, biodata model.Mahasiswa) (InsertedID interface{}) {
	var nilai model.Nilai
	nilai.All_Tugas = all_tugas
	nilai.UTS = uts
	nilai.UAS = uas
	nilai.Grade = grade
	nilai.Kategori = kategori
	nilai.Biodata = biodata
	return InsertOneDoc(db, col, nilai)
}

func InsertTugas(db *mongo.Database, col string, tugas1 int, tugas2 int, tugas3 int, tugas4 int, tugas5 int) (InsertedID interface{}) {
	var tugas model.Tugas
	tugas.Tugas1 = tugas1
	tugas.Tugas2 = tugas2
	tugas.Tugas3 = tugas3
	tugas.Tugas4 = tugas4
	tugas.Tugas5 = tugas5
	return InsertOneDoc(db, col, tugas)
}

func InsertPresensi(db *mongo.Database, col string, jumlahkehadiran int, biodata model.Mahasiswa) (InsertedID interface{}) {
	var presensi model.Presensi
	presensi.Jumlah_Kehadiran = jumlahkehadiran
	presensi.Biodata = biodata
	return InsertOneDoc(db, col, presensi)
}

func InsertGrade(db *mongo.Database, col string, namagrade string, skala string) (InsertedID interface{}) {
	var grade model.Grade
	grade.Nama_Grade = namagrade
	grade.Skala = skala
	return InsertOneDoc(db, col, grade)
}

func GetPresensiFromMahasiswa(nama string, db *mongo.Database, col string) (identitas model.Presensi) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"biodata.nama": nama}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&identitas)
	if err != nil {
		fmt.Printf("GetPresensiFromMahasiswa: %v\n", err)
	}
	return identitas
}

func GetUASFromMahasiswa(npm int, db *mongo.Database, col string) (nilai model.Nilai) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"biodata.npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&nilai)
	if err != nil {
		fmt.Printf("GetUASFromMahasiswa: %v\n", err)
	}
	return nilai
}

func GetGradeFromMahasiswa(npm int, db *mongo.Database, col string) (nilai model.Nilai) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"biodata.npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&nilai)
	if err != nil {
		fmt.Printf("GetGradeFromMahasiswa: %v\n", err)
	}
	return nilai
}

func GetMatakuliahFromJadwal(jam_masuk string, db *mongo.Database, col string) (jam model.Matakuliah) {
	jadwal := db.Collection(col)
	filter := bson.M{"jadwal.jam_masuk": jam_masuk}
	err := jadwal.FindOne(context.TODO(), filter).Decode(&jam)
	if err != nil {
		fmt.Printf("GetMatakuliahFromJadwal: %v\n", err)
	}
	return jam
}

func GetMatakuliahFromNIK(nik string, db *mongo.Database, col string) (matkul model.Matakuliah) {
	dosen := db.Collection(col)
	filter := bson.M{"pengampu.nik": nik}
	err := dosen.FindOne(context.TODO(), filter).Decode(&matkul)
	if err != nil {
		fmt.Printf("GetMatakuliahFromNIK: %v\n", err)
	}
	return matkul
}

func GetNilaiFromMatakuliah(nama_mk string, db *mongo.Database, col string) (data model.Nilai) {
	matakuliah := db.Collection(col)
	filter := bson.M{"kategori.nama_mk": nama_mk}
	err := matakuliah.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetNilaiFromMatakuliah: %v\n", err)
	}
	return data
}

func GetNilaiFromNamaMahasiswa(nama string, db *mongo.Database, col string) (data model.Nilai) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"biodata.nama": nama}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetNilaiFromNamaMahasiswa: %v\n", err)
	}
	return data
}

func GetAllNilaiFromNamaMahasiswa(nama string, db *mongo.Database, col string) (data []model.Nilai) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"biodata.nama": nama}
	cursor, err := mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
