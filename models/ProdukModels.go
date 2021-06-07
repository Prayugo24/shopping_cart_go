package models


type Tb_Produk struct {
	Id int `json:"id", gorm:"primary_key", gorm:autoIncrement` // id
	Kode_produk string `json:"kode_produk"` 
	Kuantitas int `json:"kuantitas"` 
}