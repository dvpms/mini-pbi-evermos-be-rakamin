package models

import "time"

// User represents the users table in database
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NamaUser     string    `gorm:"type:varchar(255);not null" json:"nama_user"`
	KataSandi    string    `gorm:"type:varchar(255);not null" json:"-"`
	NoTelp       string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"no_telp"`
	TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"type:varchar(20)" json:"jenis_kelamin"`
	Pekerjaan    string    `gorm:"type:varchar(100)" json:"pekerjaan"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	IDProvinsi   string    `gorm:"type:varchar(50)" json:"id_provinsi"`
	IDKota       string    `gorm:"type:varchar(50)" json:"id_kota"`
	Role         string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relationships
	Toko      *Toko       `json:"toko,omitempty" gorm:"foreignKey:UserID"`
	Alamat    []Alamat    `json:"alamat,omitempty" gorm:"foreignKey:UserID"`
	Transaksi []Transaksi `json:"transaksi,omitempty" gorm:"foreignKey:UserID"`
}

// RegisterRequest represents the JSON body for user registration
type RegisterRequest struct {
	Nama         string `json:"nama" binding:"required"`
	KataSandi    string `json:"kata_sandi" binding:"required,min=6"`
	NoTelp       string `json:"no_telp" binding:"required"`
	TanggalLahir string `json:"tanggal_Lahir"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email" binding:"required,email"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
}

// LoginRequest represents the JSON body for user login
type LoginRequest struct {
	NoTelp    string `json:"no_telp" binding:"required"`
	KataSandi string `json:"kata_sandi" binding:"required"`
}

// UpdateProfileRequest represents the JSON body for profile update
type UpdateProfileRequest struct {
	Nama         string `json:"nama"`
	KataSandi    string `json:"kata_sandi"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_Lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IDProvinsi   string `json:"id_provinsi"`
	IDKota       string `json:"id_kota"`
}

// LoginResponseData represents the data payload for successful login
type LoginResponseData struct {
	Nama         string      `json:"nama"`
	NoTelp       string      `json:"no_telp"`
	TanggalLahir string      `json:"tanggal_Lahir"`
	Tentang      string      `json:"tentang"`
	Pekerjaan    string      `json:"pekerjaan"`
	Email        string      `json:"email"`
	IDProvinsi   interface{} `json:"id_provinsi"`
	IDKota       interface{} `json:"id_kota"`
	Token        string      `json:"token"`
}
