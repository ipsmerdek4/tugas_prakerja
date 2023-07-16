package models

type T_Donasi_Buku struct {
	Nim       int    `gorm:"primary_key;"`
	Nama_buku string `json:"nama_buku"`
	Name      string `json:"name"`
	Nomer_hp  string `json:"nomer_hp"`
	Jurusan   string `json:"jurusan"`
	Alamat    string `json:"alamat"`
}
