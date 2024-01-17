package controllers

import (
	"fmt"
	"go-fiber-test/database"
	"go-fiber-test/models"
	m "go-fiber-test/models"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BodyParser(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	c.Query("search") // "fenny"

	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidateTest(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func CalculateFactorial(c *fiber.Ctx) error {
	result := new(models.APIResponse)
	n, err := strconv.Atoi(c.Params("n"))

	if err != nil {
		result.Data = nil
		result.Message = err.Error()
		result.Success = false
		result.Status = 400
		return c.Status(400).JSON(result)
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}

	result.Data = fmt.Sprintf("%d! = %d", n, ans)
	result.Message = "success"
	result.Success = true
	result.Status = 200

	return c.Status(result.Status).JSON(result)
}

func ChangeStringToAssci(c *fiber.Ctx) error {
	result := new(models.APIResponse)
	taxId := c.Query("tax_id")

	result.Data = fmt.Sprint([]rune(taxId))
	result.Message = "success"
	result.Success = true
	result.Status = 200

	return c.Status(result.Status).JSON(result)
}

func Register(c *fiber.Ctx) error {
	result := new(m.APIResponse)
	member := new(m.Member)
	err := c.BodyParser(&member)
	if err != nil {
		result.Data = nil
		result.Message = err.Error()
		result.Success = false
		result.Status = 500
		return c.Status(result.Status).JSON(result)
	}

	validate := validator.New()
	errors := validate.Struct(member)
	if errors != nil {
		result.Data = nil
		result.Message = err.Error()
		result.Success = false
		result.Status = 500

		return c.Status(result.Status).JSON(result)
	}

	// usernamePattern := `^[a-zA-Z0-9_.]+$`
	// usernameRegex := regexp.MustCompile(usernamePattern)
	// re := usernameRegex.MatchString(member.Username)
	// if !re {
	// 	result.Data = member
	// 	result.Message = "ชื่อผู้ใช้ต้องประกอบไปด้วย ตัวอักษร [A-Z] [a-z] [0-9] [_] [.] เท่านั้น"
	// 	result.Status = 400
	// 	result.Success = false

	// 	return c.Status(result.Status).JSON(result)
	// }

	result.Data = member
	result.Message = "Success"
	result.Status = 201
	result.Success = true

	return c.Status(result.Status).JSON(result)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func GetSoftDeleteDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dog []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dog)

	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNoColor := 0

	for _, v := range dogs { //1 inet 112 //2 inet1 113
		color := ""
		if v.DogID >= 250 && v.DogID <= 200 {
			color = "pink"
			sumPink += 1
		} else if v.DogID >= 150 && v.DogID <= 100 {
			color = "green"
			sumGreen += 1
		} else if v.DogID >= 50 && v.DogID <= 10 {
			color = "red"
			sumRed += 1
		} else {
			color = "no color"
			sumNoColor += 1
		}
		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Color: color,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:       dataResults,
		Name:       "golang-test",
		Count:      len(dogs), //หาผลรวม,
		SumRed:     sumRed,
		SumGreen:   sumGreen,
		SumPink:    sumPink,
		SumNocolor: sumNoColor,
	}
	return c.Status(200).JSON(r)
}

func GetDogsByIdRange(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogScopeRangeBetween50to100).Find(&dogs)

	return c.Status(200).JSON(dogs)
}

func DogScopeRangeBetween50to100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id BETWEEN ? AND ?", 50, 100)
}
