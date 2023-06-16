package module

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	model "github.com/Febriand1/Nilai/Model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
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

func InsertMahasiswa(db *mongo.Database, col string, nama string, npm int, phonenumber string) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.NPM = npm
	mahasiswa.Phone_Number = phonenumber
	return InsertOneDoc(db, col, mahasiswa)
}

func InsertMatakuliah(db *mongo.Database, col string, namamk string, sks int, jadwal model.Waktu, pengampu model.Dosen) (InsertedID interface{}) {
	var matakuliah model.Matakuliah
	matakuliah.Nama_MK = namamk
	matakuliah.SKS = sks
	matakuliah.Jadwal = jadwal
	matakuliah.Pengampu = pengampu
	return InsertOneDoc(db, col, matakuliah)
}

func InsertWaktu(db *mongo.Database, col string, jammasuk string, jamkeluar string, hari []string) (InsertedID interface{}) {
	var waktu model.Waktu
	waktu.Jam_Masuk = jammasuk
	waktu.Jam_Keluar = jamkeluar
	waktu.Hari = hari
	return InsertOneDoc(db, col, waktu)
}

func InsertDosen(db *mongo.Database, col string, namadosen string, nik string, phonenumberd string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.Nama_Dosen = namadosen
	dosen.NIK = nik
	dosen.Phone_NumberD = phonenumberd
	return InsertOneDoc(db, col, dosen)
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

func InsertGrade(db *mongo.Database, col string, namagrade string, rata_rata string) (InsertedID interface{}) {
	var grade model.Grade
	grade.Nama_Grade = namagrade
	grade.Rata_Rata = rata_rata
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

func GetMatakuliahFromJadwal(jammasuk string, db *mongo.Database, col string) (jam model.Matakuliah) {
	jadwal := db.Collection(col)
	filter := bson.M{"jadwal.jammasuk": jammasuk}
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

func GetNilaiFromMatakuliah(namamk string, db *mongo.Database, col string) (data model.Nilai) {
	matakuliah := db.Collection(col)
	filter := bson.M{"kategori.namamk": namamk}
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


//TB
func GetAllNilai(db *mongo.Database, col string) (data []model.Nilai) {
	nilai := db.Collection(col)
	filter := bson.M{}
	cursor, err := nilai.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetNilaiFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mhs model.Nilai, errs error) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mhs, fmt.Errorf("no data found for ID %s", _id)
		}
		return mhs, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mhs, nil
}

// func InsertNilai(db *mongo.Database, col string, alltugas model.Tugas, uts int, uas int, grade model.Grade, kategori model.Matakuliah, absensi model.Presensi) (InsertedID interface{}) {
// 	var nilai model.Nilai
// 	nilai.All_Tugas = alltugas
// 	nilai.UTS = uts
// 	nilai.UAS = uas
// 	nilai.Grade = grade
// 	nilai.Kategori = kategori
// 	nilai.Absensi = absensi
// 	return InsertOneDoc(db, col, nilai)
// }

func InsertNilai(db *mongo.Database, col string, alltugas model.Tugas, uts int, uas int, grade model.Grade, kategori model.Matakuliah, absensi model.Presensi) (insertedID primitive.ObjectID, err error) {
	nilai := bson.M{
		"alltugas":	alltugas,
		"uts":     	uts,
		"uas":     	uas,
		"grade":   	grade,
		"kategori": kategori,
		"absensi":  absensi,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), nilai)
	if err != nil {
		fmt.Printf("InsertNilai: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func DeleteNilaiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	mahasiswa := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := mahasiswa.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func UpdateNilai(db *mongo.Database, col string, id primitive.ObjectID, alltugas model.Tugas, uts int, uas int, grade model.Grade, kategori model.Matakuliah, absensi model.Presensi) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"alltugas"	: alltugas,
			"uts"		: uts,
			"uas"		: uas,
			"grade"		: grade,
			"kategori"	: kategori,
			"absensi"	: absensi,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateNilai: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}
//TB

//login
func InsertAdmin(db *mongo.Database, col string, username string, password string) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
		"username":	username,
		"password": password,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAdminFromID(_id primitive.ObjectID, db *mongo.Database, col string) (adm model.Admin, errs error) {
	admin := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := admin.FindOne(context.TODO(), filter).Decode(&adm)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return adm, fmt.Errorf("no data found for ID %s", _id)
		}
		return adm, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return adm, nil
}

func InsertUser(db *mongo.Database, col string, usernamem string, passwordm string) (insertedID primitive.ObjectID, err error) {
	user := bson.M{
		"usernamem": usernamem,
		"passwordm": passwordm,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database, col string) (usr model.User, errs error) {
	user := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := user.FindOne(context.TODO(), filter).Decode(&usr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return usr, fmt.Errorf("no data found for ID %s", _id)
		}
		return usr, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return usr, nil
}

//login