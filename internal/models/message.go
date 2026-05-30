package models

import "time"

type Message struct {
	Code  string `json:"Code" binding:"required"`
	Email string `json:"Email" binding:"required,email"`
}

type SubscriptionNotificationMsg struct {
	Email      string    `json:"email" binding:"required,email"`
	Type       string    `json:"type" binding:"required"` // "expiring_soon" or "expired"
	ExpireDate time.Time `json:"expire_date" binding:"required"`
	RenewalURL string   `json:"renewal_url" binding:"required"`
}
