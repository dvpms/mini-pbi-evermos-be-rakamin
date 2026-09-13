package models

import "time"

// Category represents categories table in database
type Category struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NamaCategory string    `gorm:"type:varchar(255);not null" json:"nama_category"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	// Relationships
	Produk []Produk `json:"produk,omitempty" gorm:"foreignKey:CategoryID"`
}

// CategoryRequest represents input body for creating or updating a category
type CategoryRequest struct {
	NamaCategory string `json:"nama_category" binding:"required"`
}

// CategorySimpleResponse represents clean view of category
type CategorySimpleResponse struct {
	ID           uint   `json:"id"`
	NamaCategory string `json:"nama_category"`
}
