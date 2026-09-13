package models

import "time"

// LogProduk represents snapshot of product state at transaction time
type LogProduk struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	ProdukID      uint         `gorm:"not null;index" json:"product_id"`
	NamaProduk    string       `gorm:"type:varchar(255);not null" json:"nama_produk"`
	Slug          string       `gorm:"type:varchar(255);not null" json:"slug"`
	HargaReseller float64      `gorm:"type:decimal(15,2);not null" json:"harga_reseler"`
	HargaKonsumen float64      `gorm:"type:decimal(15,2);not null" json:"harga_konsumen"`
	Deskripsi     string       `gorm:"type:text" json:"deskripsi"`
	TokoID        uint         `gorm:"not null;index" json:"toko_id,omitempty"`
	CategoryID    uint         `gorm:"not null;index" json:"category_id,omitempty"`
	CreatedAt     time.Time    `json:"created_at,omitempty"`
	UpdatedAt     time.Time    `json:"updated_at,omitempty"`

	// Relationships
	Toko     Toko         `json:"toko" gorm:"foreignKey:TokoID;references:ID"`
	Category Category     `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	Photos   []FotoProduk `json:"photos" gorm:"-"`
}
