package module

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	"github.com/dimasardnt6/kemahasiswaan/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "kemahasiswaan_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKemahasiswaan(db *mongo.Database, col string, identitas model.Mahasiswa, status_keuangan model.Keuangan, nilai_mhs model.Nilai) (InsertedID interface{}) {
	var kemahasiswaan model.Kemahasiswaan
	kemahasiswaan.Identitas_Mahasiswa = identitas
	kemahasiswaan.Status_Keuangan = status_keuangan
	kemahasiswaan.Nilai_Mahasiswa = nilai_mhs
	return InsertOneDoc(db, col, kemahasiswaan)
}

func InsertDataMahasiswa(db *mongo.Database, col string, npm string, nama string, no_hp string, prodi model.ProgramStudi, jurusan string, kelas string) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Npm = npm
	mahasiswa.Nama = nama
	mahasiswa.Nomor_Handphone = no_hp
	mahasiswa.Prodi = prodi
	mahasiswa.Jurusan = jurusan
	mahasiswa.Kelas = kelas
	return InsertOneDoc(db, col, mahasiswa)
}

func InsertKeuanganMahasiswa(db *mongo.Database, col string, biodata model.Mahasiswa, total_pembayaran int) (InsertedID interface{}) {
	var keuangan model.Keuangan
	keuangan.Biodata = biodata
	keuangan.Total_Pembayaran = total_pembayaran
	return InsertOneDoc(db, col, keuangan)
}

func InsertNilaiMahasiswa(db *mongo.Database, col string, bio_mhs model.Mahasiswa, matakuliah model.Matakuliah, nilai_angka int, nilai_huruf string) (InsertedID interface{}) {
	var nilai model.Nilai
	nilai.Biodata_Mahasiswa = bio_mhs
	nilai.Matakuliah = matakuliah
	nilai.Nilai_Angka = nilai_angka
	nilai.Nilai_Huruf = nilai_huruf
	return InsertOneDoc(db, col, nilai)
}

// test getFunction

func GetKemahasiswaanFromNpm(npm string, db *mongo.Database, col string) (data model.Kemahasiswaan) {
	kemahasiswaan := db.Collection(col)
	filter := bson.M{"identitas.npm": npm}
	err := kemahasiswaan.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getKemahasiswaanFromNpm: %v\n", err)
	}
	return data
}

func GetDataMahasiswaFromNpm(npm string, db *mongo.Database, col string) (data model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataMahasiswaFromNpm: %v\n", err)
	}
	return data
}

func GetKeuanganMahasiswaFromNomorHp(no_hp string, db *mongo.Database, col string) (data model.Keuangan) {
	keuangan := db.Collection(col)
	filter := bson.M{"bio_mahasiswa.no_hp": no_hp}
	err := keuangan.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getKeuanganMahasiswaFromNomorHp: %v\n", err)
	}
	return data
}

func GetNilaiMahasiswaFromNama(nama string, db *mongo.Database, col string) (data model.Nilai) {
	nilai := db.Collection(col)
	filter := bson.M{"bio_mhs.nama": nama}
	err := nilai.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getNilaiKemahasiswaanFromNama: %v\n", err)
	}
	return data
}

//GetFunctionAll

func GetAllKemahasiswaan(db *mongo.Database, col string) (kemahasiswaan []model.Kemahasiswaan) {
	data_kemahasiswaan := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_kemahasiswaan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &kemahasiswaan)
	if err != nil {
		fmt.Println(err)
	}
	return kemahasiswaan
}
func GetAllDataMahasiswa(db *mongo.Database, col string) (mahasiswa []model.Mahasiswa) {
	data_mahasiswa := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &mahasiswa)
	if err != nil {
		fmt.Println(err)
	}
	return mahasiswa
}
func GetAllKeuanganMahasiswa(db *mongo.Database, col string) (keuangan []model.Keuangan) {
	keuangan_mahasiswa := db.Collection(col)
	filter := bson.M{}
	cursor, err := keuangan_mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &keuangan)
	if err != nil {
		fmt.Println(err)
	}
	return keuangan
}
func GetAllNilaiMahasiswa(db *mongo.Database, col string) (nilai []model.Nilai) {
	nilai_mahasiswa := db.Collection(col)
	filter := bson.M{}
	cursor, err := nilai_mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &nilai)
	if err != nil {
		fmt.Println(err)
	}
	return nilai
}
