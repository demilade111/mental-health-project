package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// SendEmail sends an email with the given subject and body to the recipient
func SendEmail(to, subject, body string,) {
	// Load SMTP configuration from environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASSWORD")

	// Validate configuration
	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" {
		log.Println("Error: Missing SMTP configuration")
		return
	}

	// Build the email message
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// Configure the SMTP authentication
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Send the email
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort), // SMTP server address
		auth,
		smtpUser,        // Sender email
		[]string{to},    // Recipient email
		[]byte(message), // Email body
	)

	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", to, err)
		return
	}

	log.Printf("Email sent successfully to %s\n", to)
}

