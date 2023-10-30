package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	APP_PORT  = ":4141"
	EMAIL_URL = "http://localhost:5555"
)

type SendEmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

type SendEmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func main() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.Post("/send", sendEmail)

	router.Listen(APP_PORT)
}

func sendEmail(c *fiber.Ctx) error {
	var req = SendEmailRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	reqBody, _ := json.Marshal(req)

	resp, err := http.Post(fmt.Sprintf("%s/send", EMAIL_URL), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	if resp.StatusCode != 200 {
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   string(respBody),
		})
	}

	var emailResp SendEmailResponse
	err = json.Unmarshal(respBody, &emailResp)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	if !emailResp.Success {
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"success": false,
			"message": emailResp.Message,
			"error":   emailResp.Error,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "SUCCESS",
	})
}
