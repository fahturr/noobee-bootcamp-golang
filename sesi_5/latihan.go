package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var dataUser = []User{}

func main() {
	app := fiber.New()

	app.Use(Trace())

	app.Post("/users", CreateUser())
	app.Get("/users", GetListUser())
	app.Put("/users/:id", UpdateUser())
	app.Delete("/users/:id", DeleteUser())

	log.Fatal(app.Listen(":4444"))
}

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := c.Method()
		uri := c.Request().URI().Path()
		traceId := uuid.New().String()

		c.Response().Header.Add("X-Request-Id", traceId)

		log.Printf(`message="incoming request" method="%s" uri="%s" trace_id="%s"`, method, uri, traceId)
		err := c.Next()
		if err != nil {
			log.Printf(`message="%s" method="%s" uri="%s" trace_id="%s"`, err.Error(), method, uri, traceId)
		}
		log.Printf(`message="finish request" method="%s" uri="%s" trace_id="%s"`, method, uri, traceId)

		return err
	}
}

func CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := User{}

		err := c.BodyParser(&user)
		if err != nil {
			return err
		}

		dataUser = append(dataUser, user)

		return c.JSON(fiber.Map{
			"success":     true,
			"status_code": fiber.StatusCreated,
			"message":     "created success",
		})
	}
}

func GetListUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := []User{}
		for i, v := range dataUser {
			user := v
			user.Id = i + 1

			data = append(data, user)
		}

		return c.JSON(fiber.Map{
			"success":     true,
			"status_code": fiber.StatusOK,
			"message":     "get all success",
			"payload":     data,
		})
	}
}

func UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		user := User{}

		if id == 0 || id > len(dataUser) {
			return fiber.NewError(fiber.StatusBadRequest, "id not registered")
		}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		dataUser[id-1] = user

		return c.JSON(fiber.Map{
			"success":     true,
			"status_code": fiber.StatusOK,
			"message":     "get all success",
		})
	}
}

func DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		if id == 0 || id > len(dataUser) {
			return fiber.NewError(fiber.StatusBadRequest, "id not registered")
		}

		if len(dataUser) == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "data user is empty")
		}

		// re-slice dataUser
		index := id - 1
		dataUser = append(dataUser[:index], dataUser[index+1:]...)

		return c.JSON(fiber.Map{
			"success":     true,
			"status_code": fiber.StatusOK,
			"message":     "delete success",
		})
	}
}
