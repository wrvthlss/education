package main

import "fmt"

// Notifier is an interface for sending notification.
type Notifier interface {
	Notify(message string)
}

// EmailNotifier Implements Notifier with email notification.
type EmailNotifier struct{}

func (n EmailNotifier) Notify(message string) {
	fmt.Println("sending email: ", message)
}

// LogNotifier implements Notifier with log notifications.
type LogNotifier struct{}

func (n LogNotifier) Notify(message string) {
	fmt.Println("Logging: ", message)
}

func sendNotification(notifier Notifier, message string) {
	notifier.Notify(message)
}

func main() {

	emailNotifier := EmailNotifier{}
	logNotifier := LogNotifier{}

	sendNotification(logNotifier, "[+] Notify log message...")
	sendNotification(emailNotifier, "[!] Notify email message...")

}
