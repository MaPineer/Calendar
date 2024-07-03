package models

import "log"

type Notifier interface {
	Send(to, content string) error
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(to, content string) error {
	// Mock implementation, replace with actual email sending logic
	log.Printf("Sending email to %s: %s", to, content)
	return nil
}

type MessageNotifier struct{}

func (s MessageNotifier) Send(to, content string) error {
	log.Printf("Sending Message to %s: %s", to, content)
	return nil
}
