package main

import (
	"github.com/bxcodec/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
)

type Item struct {
	Id    uint
	Name  string
	Email string
	Phone string
	Price int
}

func main() {

	// Connect to MySQL Database

	db, err := gorm.Open(mysql.Open("root:RootPassword/fakedata"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	// Create Database Table (Items)

	db.AutoMigrate(&Item{})

	// Initiate GoFiber server

	app := fiber.New()
	app.Use(cors.New())

	// Post handler request

	app.Post("/api/item/create", func(c *fiber.Ctx) error {

		for i := 0; i < 500; i++ {
			db.Create(&Item{

				// Create Fake Name with faker repo
				Name: faker.Word(),
				// Create Fake Email with faker repo
				Email: faker.Email(),
				// Create Fake Phone No.
				Phone: faker.Phonenumber(),
				// Create fake Price using random EQ
				Price: rand.Intn(140) + 10,
			})

		}
		// Return message if success

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
		})
	})
	// Get request
	app.Get("/api/item/all", func(c *fiber.Ctx) error {

		var items []Item
		db.Find(&items)
		return c.Status(200).JSON(items)
	})

	// Listen to port 8000

	app.Listen(":8000")
}
