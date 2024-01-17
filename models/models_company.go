package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName    string `json:"companyName" validate:"required"`
	CompanyAddress string `json:"comapanyAddress" validate:"required"`
	CompanyTel     string `json:"comapanyTel" validate:"required"`
	CompanySite    string `json:"comapanySite" validate:"required"`
	CompanyEmail   string `json:"comapanyEmail" validate:"required"`
}

type CompanyResponse struct {
	Id             int
	CompanyName    string
	CompanyAddress string
	CompanyTel     string
	CompanyEmail   string
	CompanySite    string
}
