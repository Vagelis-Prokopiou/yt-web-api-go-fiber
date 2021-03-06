package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type User struct {
	Id         int
	Age        int
	First_Name string
	Last_Name  string
	Framework  string
}

func getUsers() [1000]User {
	var users [1000]User

	for i := 1; i < 1001; i++ {
		var stringIndex = strconv.Itoa(i)
		users[i-1] = User{
			Id:         i,
			Age:        25,
			First_Name: "First_name" + stringIndex,
			Last_Name:  "Last_Name" + stringIndex,
			Framework:  "Golang (fiber)  ",
		}
	}

	return users
}

func main() {
	app := fiber.New()

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON(getUsers())
	})

	app.Listen(":3000")
}
