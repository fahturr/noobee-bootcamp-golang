package main

import (
	"log"
	"mail_campaign/mail/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

type SendEmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

func main() {
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	app := fiber.New()

	app.Post("/send", sendEmail)
	app.Listen(config.Cfg.App.Port)
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

	err = sendMailGoMail(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ERR BAD REQUEST",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "SUCCESS",
	})
}

func sendMailGoMail(req SendEmailRequest) (err error) {
	mailer := gomail.NewMessage()

	// Set Header
	mailer.SetHeader("From", req.From)
	mailer.SetHeader("To", req.To...)

	// Set Content
	mailer.SetHeader("Subject", req.Subject)
	mailer.SetBody("text/html", req.Message)

	dialer := gomail.NewDialer(
		config.Cfg.Mail.Host,
		config.Cfg.Mail.Port,
		config.Cfg.Mail.Email,
		config.Cfg.Mail.Password,
	)

	err = dialer.DialAndSend(mailer)
	return
}
