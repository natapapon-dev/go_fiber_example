package controllers

import (
	"fmt"
	"go-fiber-test/database"
	"go-fiber-test/models"
	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
)

func CreateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var companies models.Company

	if err := c.BodyParser(&companies); err != nil {
		apiResponse := m.APIResponse{
			Data:    nil,
			Message: err.Error(),
			Status:  503,
			Success: false,
		}
		return c.Status(apiResponse.Status).JSON(apiResponse)
	}

	db.Model(&m.Company{}).Create(&companies).Find("")

	apiResponse := m.APIResponse{
		Data:    companies,
		Message: "Success",
		Status:  201,
		Success: true,
	}

	return c.Status(apiResponse.Status).JSON(apiResponse)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var companies models.Company
	id := c.Params("id")

	if err := c.BodyParser(&companies); err != nil {
		apiResponse := m.APIResponse{
			Data:    nil,
			Message: err.Error(),
			Status:  503,
			Success: false,
		}
		return c.Status(apiResponse.Status).JSON(apiResponse)
	}

	result := db.Where("id = ?", id).Updates(&companies)
	if result.RowsAffected == 0 {
		apiResponse := m.APIResponse{
			Data:    nil,
			Message: fmt.Sprintf("Company ID :: %s is not found", id),
			Status:  404,
			Success: false,
		}
		return c.Status(apiResponse.Status).JSON(apiResponse)
	}

	apiResponse := m.APIResponse{
		Data:    companies,
		Message: "Success",
		Status:  200,
		Success: true,
	}
	return c.Status(200).JSON(apiResponse)
}

func SoftDelteCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var companies m.Company

	result := db.Delete(&companies, id)

	if result.RowsAffected == 0 {
		apiResponse := m.APIResponse{
			Data:    nil,
			Message: fmt.Sprintf("Company ID :: %s is not found", id),
			Status:  404,
			Success: false,
		}
		return c.Status(apiResponse.Status).JSON(apiResponse)
	}

	apiResponse := m.APIResponse{
		Data:    nil,
		Message: fmt.Sprintf("Successfully DELETE Company ID :: %s is not found", id),
		Status:  200,
		Success: true,
	}
	return c.Status(apiResponse.Status).JSON(apiResponse)
}

func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.CompanyResponse

	result := db.Model(&models.Company{}).Find(&companies)

	if result.RowsAffected == 0 {
		apiResponse := m.APIResponse{
			Data:    nil,
			Message: "Company Not Found",
			Status:  404,
			Success: true,
		}
		c.Status(apiResponse.Status).JSON(apiResponse)
	}

	apiResponse := m.APIResponse{
		Data:    companies,
		Message: "success",
		Status:  200,
		Success: true,
	}

	return c.Status(apiResponse.Status).JSON(apiResponse)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.CompanyResponse
	id := c.Params("id")

	result := db.Model(&m.Company{}).Where("id = ?", id).Find(&company)

	if result.RowsAffected == 0 {
		apiResponse := m.APIResponse{
			Data:    nil,
			Message: fmt.Sprintf("Company ID :: %s is not found", id),
			Status:  404,
			Success: false,
		}
		return c.Status(apiResponse.Status).JSON(apiResponse)
	}

	apiResponse := m.APIResponse{
		Data:    company,
		Message: "Success",
		Status:  200,
		Success: true,
	}

	return c.Status(apiResponse.Status).JSON(apiResponse)
}
