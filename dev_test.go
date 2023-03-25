package kemahasiswaan1214054

import (
	"fmt"
	"testing"

	"github.com/dimasardnt6/kemahasiswaan/model"
	"github.com/dimasardnt6/kemahasiswaan/module"
)

func TestInsertKemahasiswaan(t *testing.T) {
	identitas := model.Mahasiswa{
		Npm:             "1214054",
		Nama:            "Dimas Ardianto",
		Nomor_Handphone: "6289647129890",
		Prodi: model.ProgramStudi{
			Kode_Prodi: "012021",
			Nama_Prodi: "D4",
		},
		Jurusan: "Teknik Informatika",
		Kelas:   "2B",
	}
	status_keuangan := model.Keuangan{
		Total_Pembayaran: 7700000,
	}
	nilai_mhs := model.Nilai{
		Matakuliah: model.Matakuliah{
			Nama_Matkul: "Pemrograman",
			Nama_Dosen:  "Indra Riksa",
		},
		Nilai_Angka: 90,
		Nilai_Huruf: "A",
	}
	hasil := module.InsertKemahasiswaan(module.MongoConn, "kemahasiswaan", identitas, status_keuangan, nilai_mhs)
	fmt.Println(hasil)
}

func TestInsertDataMahasiswa(t *testing.T) {
	Npm := "1214054"
	Nama := "Dimas Ardianto"
	Nomor_Handphone := "6289647129890"
	Prodi := model.ProgramStudi{
		Kode_Prodi: "012021",
		Nama_Prodi: "D4",
	}
	Jurusan := "Teknik Informatika"
	Kelas := "2B"
	hasil := module.InsertDataMahasiswa(module.MongoConn, "data_mahasiswa", Npm, Nama, Nomor_Handphone, Prodi, Jurusan, Kelas)
	fmt.Println(hasil)
}

func TestInsertKeuanganMahasiswa(t *testing.T) {
	Biodata := model.Mahasiswa{
		Npm:             "1214054",
		Nama:            "Dimas Ardianto",
		Nomor_Handphone: "6289647129890",
		Prodi: model.ProgramStudi{
			Kode_Prodi: "012021",
			Nama_Prodi: "D4",
		},
		Jurusan: "Teknik Informatika",
		Kelas:   "2B",
	}
	Total_Pembayaran := 7700000
	hasil := module.InsertKeuanganMahasiswa(module.MongoConn, "keuangan_mahasiswa", Biodata, Total_Pembayaran)
	fmt.Println(hasil)
}

func TestInsertNilaiMahasiswa(t *testing.T) {
	Biodata_Mahasiswa := model.Mahasiswa{
		Npm:             "1214054",
		Nama:            "Dimas Ardianto",
		Nomor_Handphone: "6289647129890",
		Prodi: model.ProgramStudi{
			Kode_Prodi: "012021",
			Nama_Prodi: "D4",
		},
		Jurusan: "Teknik Informatika",
		Kelas:   "2B",
	}
	Matakuliah := model.Matakuliah{
		Nama_Matkul: "Pemrograman",
		Nama_Dosen:  "Indra Riksa",
	}
	Nilai_Angka := 90
	Nilai_Huruf := "A"
	hasil := module.InsertNilaiMahasiswa(module.MongoConn, "nilai_mahasiswa", Biodata_Mahasiswa, Matakuliah, Nilai_Angka, Nilai_Huruf)
	fmt.Println(hasil)
}

// test getFunction

func TestGetKemahasiswaanFromNpm(t *testing.T) {
	Npm := "1214054"
	kemahasiswaan := module.GetKemahasiswaanFromNpm(Npm, module.MongoConn, "kemahasiswaan")
	fmt.Println(kemahasiswaan)
}

func TestGetDataMahasiswaFromNpm(t *testing.T) {
	Npm := "1214054"
	mahasiswa := module.GetDataMahasiswaFromNpm(Npm, module.MongoConn, "data_mahasiswa")
	fmt.Println(mahasiswa)
}

func TestGetKeuanganMahasiswaFromNomorHp(t *testing.T) {
	Nomor_Handphone := "6289647129890"
	keuangan := module.GetKeuanganMahasiswaFromNomorHp(Nomor_Handphone, module.MongoConn, "keuangan_mahasiswa")
	fmt.Println(keuangan)
}

func TestGetNilaiMahasiswaFromNama(t *testing.T) {
	Nama := "Dimas Ardianto"
	nilai := module.GetNilaiMahasiswaFromNama(Nama, module.MongoConn, "nilai_mahasiswa")
	fmt.Println(nilai)
}
