package models

type Mahasiswa struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	NIM     string `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}

type RequestMahasiswa struct{
	NIM     string `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`

}
