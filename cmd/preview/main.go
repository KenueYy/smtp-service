package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kenueyy/smtp-service/internal/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./cmd/preview expiring_soon|expired")
		os.Exit(1)
	}

	notifType := os.Args[1]
	email := "kenueyy@gmail.com"
	renewalURL := "https://nevpn.shop/account"

	now := time.Now()

	var expireDate time.Time
	switch notifType {
	case "expiring_soon":
		expireDate = now.Add(3 * 24 * time.Hour)
		fmt.Println("Sending expiring_soon preview...")
	case "expired":
		expireDate = now.Add(-1 * time.Hour)
		fmt.Println("Sending expired preview...")
	default:
		fmt.Println("unknown type, use: expiring_soon or expired")
		os.Exit(1)
	}

	fmt.Printf("To: %s\nType: %s\nExpireDate: %s\n\n", email, notifType, expireDate.Format("02.01.2006 15:04"))

	if err := handlers.SendSubscriptionEmail(email, notifType, expireDate, renewalURL, 3); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("OK — письмо отправлено, проверь почту.")
}
