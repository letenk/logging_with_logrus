package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

type UserAction struct {
	UserID    int
	Action    string
	Timestamp time.Time
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)

	action := UserAction{
		UserID:    12345,
		Action:    "checkout",
		Timestamp: time.Now(),
	}

	// Print log
	log.WithFields(logrus.Fields{
		"user_id":    action.UserID,
		"action":     action.Action,
		"timestamp":  time.Now().Format(time.RFC3339),
		"session_id": generateSessionID(),
		"module":     "payment_processor",
		"ip_address": "192.168.1.100",
	}).Error("Payment failed")

}

func generateSessionID() string {
	return "sess_abc123"
}
