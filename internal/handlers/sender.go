package handlers

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/kenueyy/smtp-service/internal/models"
	"github.com/kenueyy/smtp-service/internal/templates"

	"github.com/gin-gonic/gin"
	"github.com/kenueyy/smtp-service/internal/config"
	"gopkg.in/gomail.v2"
)

var (
	cfg    = config.Load()
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func SendCode(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := SendAuthCode(msg.Email, msg.Code); err != nil {
		logger.Error("send code failed", "email", msg.Email, "error", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

func SendSubscriptionNotificationHandler(c *gin.Context) {
	var msg models.SubscriptionNotificationMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := SendSubscriptionEmail(msg.Email, msg.Type, msg.ExpireDate, msg.RenewalURL); err != nil {
		logger.Error("send subscription notification failed",
			"email", msg.Email,
			"type", msg.Type,
			"error", err,
		)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("subscription notification sent",
		"email", msg.Email,
		"type", msg.Type,
	)
	c.JSON(200, gin.H{"message": "ok"})
}

func SendSupportTicketHandler(c *gin.Context) {
	var msg models.SupportTicketMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := SendSupportTicketEmail(msg.Description, msg.ContactMethod, msg.Contact, msg.ToEmail); err != nil {
		logger.Error("send support ticket email failed",
			"contact_method", msg.ContactMethod,
			"error", err,
		)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("support ticket email sent",
		"contact_method", msg.ContactMethod,
		"to_email", msg.ToEmail,
	)
	c.JSON(200, gin.H{"message": "ok"})
}

func SendSupportTicketEmail(description, contactMethod, contact, toEmail string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Обращение в поддержку — "+contactMethod)
	m.SetBody("text/html", templates.SupportTicketEmail(description, contactMethod, contact))

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.From, cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send support ticket error: %w", err)
	}

	return nil
}

func SendAuthCode(toEmail, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Код подтверждения")
	m.SetBody("text/html", templates.AuthCodeEmail(code))

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.From, cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send error: %w", err)
	}

	return nil
}

func SendSubscriptionEmail(toEmail, notifType string, expireDate interface{}, renewalURL string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", toEmail)

	var subject, body string

	switch notifType {
	case "expiring_soon":
		subject = "Подписка заканчивается через 3 дня"
		ed, ok := expireDate.(time.Time)
		if !ok {
			return fmt.Errorf("invalid expireDate type for expiring_soon")
		}
		body = templates.SubscriptionExpiringSoon(ed, renewalURL)
	case "expired":
		subject = "Подписка уже закончилась"
		ed, ok := expireDate.(time.Time)
		if !ok {
			return fmt.Errorf("invalid expireDate type for expired")
		}
		body = templates.SubscriptionExpired(ed, renewalURL)
	default:
		return fmt.Errorf("unknown notification type: %s", notifType)
	}

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.From, cfg.Password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("send error: %w", err)
	}

	return nil
}


