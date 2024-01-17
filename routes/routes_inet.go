package routes

import (
	"go-fiber-test/controllers"
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "123456",
		},
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", c.HelloWorld)
	v1.Post("/", c.BodyParser)
	v1.Get("/user/:name", c.ParamTest)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidateTest)

	v1.Get("/fact/:n", c.CalculateFactorial)
	v1.Post("/register", c.Register)

	//CRUD Company
	company := v1.Group("/companies")
	company.Get("", controllers.GetCompanies)
	company.Get("/:id", controllers.GetCompany)
	company.Post("", auth, controllers.CreateCompany)
	company.Put("/:id", auth, controllers.UpdateCompany)
	company.Delete("/:id", auth, controllers.SoftDelteCompany)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", controllers.GetDogs)
	dog.Get("/filter", controllers.GetDog)
	dog.Get("/json", controllers.GetDogsJson)
	dog.Post("", auth, controllers.AddDog)
	dog.Put("/:id", auth, controllers.UpdateDog)
	dog.Delete("/:id", auth, controllers.RemoveDog)
	dog.Get("/deleted", controllers.GetSoftDeleteDogs)
	dog.Get("/range", controllers.GetDogsByIdRange)

	v3 := api.Group("/v3")
	v3.Get("/", c.ChangeStringToAssci)

}
