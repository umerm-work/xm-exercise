package models

import "gorm.io/gorm"

type Company struct {
	ID         int    `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Code       string `json:"code"`
	Country    string `json:"country"`
	Website    string `json:"website"`
	Phone      string `json:"phone"`
	gorm.Model `json:"-"`
}

// NewCompany ...
func NewCompany(name, code, country, website, phone string) *Company {
	table := &Company{
		Name:    name,
		Code:    code,
		Country: country,
		Website: website,
		Phone:   phone,
		Model:   gorm.Model{},
	}
	return table
}
