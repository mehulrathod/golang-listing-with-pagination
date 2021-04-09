package models

type Employees struct {
	ID   uint   `gorm:"primary_key" json:"id,omitempty"`
	Name string `json:"name"`
	City string `json:"city"`
}
