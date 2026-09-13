package models

import "time"

// Transaksi represents transaksis table in database
type Transaksi struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	HargaTotal       float64           `gorm:"type:decimal(15,2);not null" json:"harga_total"`
	KodeInvoice      string            `gorm:"type:varchar(100);uniqueIndex;not null;column:kode_invoice" json:"kode_invoice"`
	MethodBayar      string            `gorm:"type:varchar(50);not null;column:method_bayar" json:"method_bayar"`
	AlamatPengiriman uint              `gorm:"not null;column:alamat_kirim_id" json:"-"`
	UserID           uint              `gorm:"not null;index" json:"user_id,omitempty"`
	CreatedAt        time.Time         `json:"created_at,omitempty"`
	UpdatedAt        time.Time         `json:"updated_at,omitempty"`

	// Relationships
	AlamatKirim Alamat            `json:"alamat_kirim" gorm:"foreignKey:AlamatPengiriman;references:ID"`
	DetailTrx   []DetailTransaksi `json:"detail_trx" gorm:"foreignKey:TransaksiID;references:ID"`
}

// DetailTransaksi represents detail_transaksis table in database
type DetailTransaksi struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"`
	LogProdukID uint      `gorm:"not null;index;column:log_produk_id" json:"-"`
	TransaksiID uint      `gorm:"not null;index;column:transaksi_id" json:"-"`
	Kuantitas   int       `gorm:"type:int;not null" json:"kuantitas"`
	HargaTotal  float64   `gorm:"type:decimal(15,2);not null" json:"harga_total"`
	TokoID      uint      `gorm:"not null;column:toko_id" json:"-"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`

	// Relationships
	Product LogProduk `json:"product" gorm:"foreignKey:LogProdukID;references:ID"`
	Toko    Toko      `json:"toko" gorm:"foreignKey:TokoID;references:ID"`
}

// TransactionItemInput represents each item in checkout input
type TransactionItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Kuantitas int  `json:"kuantitas" binding:"required,min=1"`
}

// CreateTransaksiRequest represents input JSON for creating transaction
type CreateTransaksiRequest struct {
	MethodBayar string                 `json:"method_bayar" binding:"required"`
	AlamatKirim uint                   `json:"alamat_kirim" binding:"required"`
	DetailTrx   []TransactionItemInput `json:"detail_trx" binding:"required,dive"`
}
