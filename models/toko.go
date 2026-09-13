package models

import "time"

// Toko represents the tokos table in database
type Toko struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NamaToko  string    `gorm:"type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string    `gorm:"type:varchar(255);column:url_foto" json:"url_foto"`
	UserID    uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	User   *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Produk []Produk `json:"produk,omitempty" gorm:"foreignKey:TokoID"`
}

// TokoSimpleResponse represents a simplified store view in lists
type TokoSimpleResponse struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
}
