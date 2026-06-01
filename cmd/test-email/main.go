package main

import (
	"fmt"
	"os"
	"time"
	"gopkg.in/gomail.v2"
	"github.com/kenueyy/smtp-service/internal/templates"
)

func main() {
	from := "no-reply@nevpn.shop"
	password := "yy4xOp10Jp2s"
	host := "mail.hosting.reg.ru"
	port := 465
	to := "kenueyy@gmail.com"
	renewalURL := "https://nevpn.shop/account"
	daysLeft := 0 // "today"

	// 20 hours from now
	expireDate := time.Now().Add(20 * time.Hour)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Подписка заканчивается сегодня")
	m.SetBody("text/html", templates.SubscriptionExpiringSoon(expireDate, renewalURL, daysLeft))

	d := gomail.NewDialer(host, port, from, password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
	fmt.Println("OK — письмо отправлено (daysLeft=0, сегодня)")
}
