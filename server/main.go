package main

import (
	"go-mysql-react/helpers"
	"go-mysql-react/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllUsers(c *fiber.Ctx) error {
	var users models.Users
	var arrUser []models.Users
	var response models.Response

	db := helpers.ConnectDB()
	defer db.Close()

	rows, err := db.Query("Select id, name ,email, password from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Name, &users.Email, &users.Password); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to retrieve data",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "data retrieved 🔥",
		"data":    response,
	})
}

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./out")
	app.Get("/api/getUser/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	app.Get("/api/getAllUsers", getAllUsers)

	app.Listen(":13322")

}
