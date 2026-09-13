package models

import "time"

// Alamat represents the alamats table in database
type Alamat struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	JudulAlamat  string    `gorm:"type:varchar(255)" json:"judul_alamat"`
	NamaPenerima string    `gorm:"type:varchar(255);not null" json:"nama_penerima"`
	NoTelp       string    `gorm:"type:varchar(50);not null" json:"no_telp"`
	DetailAlamat string    `gorm:"type:text;not null" json:"detail_alamat"`
	UserID       uint      `gorm:"not null;index" json:"user_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	// Relationships
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// CreateAlamatRequest represents input body for adding shipping address
type CreateAlamatRequest struct {
	JudulAlamat  string `json:"judul_alamat" binding:"required"`
	NamaPenerima string `json:"nama_penerima" binding:"required"`
	NoTelp       string `json:"no_telp" binding:"required"`
	DetailAlamat string `json:"detail_alamat" binding:"required"`
}

// UpdateAlamatRequest represents input body for updating shipping address
type UpdateAlamatRequest struct {
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}
