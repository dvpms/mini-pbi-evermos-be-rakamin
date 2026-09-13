package models

import "time"

// Produk represents produks table in database
type Produk struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	NamaProduk    string       `gorm:"type:varchar(255);not null" json:"nama_produk"`
	Slug          string       `gorm:"type:varchar(255);not null;index" json:"slug"`
	HargaReseller float64      `gorm:"type:decimal(15,2);not null" json:"harga_reseler"`
	HargaKonsumen float64      `gorm:"type:decimal(15,2);not null" json:"harga_konsumen"`
	Stok          int          `gorm:"type:int;not null" json:"stok"`
	Deskripsi     string       `gorm:"type:text" json:"deskripsi"`
	TokoID        uint         `gorm:"not null;index" json:"toko_id,omitempty"`
	CategoryID    uint         `gorm:"not null;index" json:"category_id,omitempty"`
	CreatedAt     time.Time    `json:"created_at,omitempty"`
	UpdatedAt     time.Time    `json:"updated_at,omitempty"`

	// Relationships
	Toko     Toko         `json:"toko" gorm:"foreignKey:TokoID;references:ID"`
	Category Category     `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	Photos   []FotoProduk `json:"photos" gorm:"foreignKey:ProductID;references:ID"`
}

// FotoProduk represents foto_produks table in database
type FotoProduk struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null;index;column:product_id" json:"product_id"`
	URL       string    `gorm:"type:varchar(255);not null;column:url" json:"url"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// CreateProdukRequest represents input data for creating a product
type CreateProdukRequest struct {
	NamaProduk    string  `form:"nama_produk" binding:"required"`
	CategoryID    uint    `form:"category_id" binding:"required"`
	HargaReseller float64 `form:"harga_reseller" binding:"required"`
	HargaKonsumen float64 `form:"harga_konsumen" binding:"required"`
	Stok          int     `form:"stok" binding:"required"`
	Deskripsi     string  `form:"deskripsi"`
}

// UpdateProdukRequest represents input data for updating a product
type UpdateProdukRequest struct {
	NamaProduk    string  `form:"nama_produk"`
	CategoryID    uint    `form:"category_id"`
	HargaReseller float64 `form:"harga_reseller"`
	HargaKonsumen float64 `form:"harga_konsumen"`
	Stok          int     `form:"stok"`
	Deskripsi     string  `form:"deskripsi"`
}
