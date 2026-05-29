package handlers

import (
	"fmt"

	"github.com/kenueyy/smtp-service/internal/models"
	"github.com/kenueyy/smtp-service/internal/templates"

	"github.com/gin-gonic/gin"
	"github.com/kenueyy/smtp-service/internal/config"
	"gopkg.in/gomail.v2"
)

var cfg = config.Load()

func SendCode(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	SendAuthCode(msg.Email, msg.Code)
}

func SendAuthCode(toEmail, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Код подтверждения")
	m.SetBody("text/html", templates.AuthCodeEmail(code))

	d := gomail.NewDialer("mail.hosting.reg.ru", 465, cfg.From, cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send error: %w", err)
	}

	return nil
}


