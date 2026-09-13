package models

// Province represents province item from EMSIFA API
type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// City represents city item from EMSIFA API
type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}
