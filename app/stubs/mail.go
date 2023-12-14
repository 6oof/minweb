package stubs

import (
	"fmt"
	"net/smtp"
	"strings"
)

// SMTPConfig represents the SMTP configuration parameters.
type SMTPConfig struct {
	Host             string
	Port             string
	Username         string
	Password         string
	SecurityStrategy string
}

// EmailMessage represents the details of an email message.
type EmailMessage struct {
	To          string
	Subject     string
	Body        string
	Attachments []string
	CC          []string
}

// GetSMTPConfig retrieves SMTP configuration from environment variables.
// If variables are not set, it returns an error.
func GetSMTPConfig() (SMTPConfig, error) {
	host := GetEnvVarOrPanic("SMTP_HOST")
	port := GetEnvVarOrPanic("SMTP_PORT")
	username := GetEnvVarOrPanic("SMTP_USERNAME")
	password := GetEnvVarOrPanic("SMTP_PASSWORD")
	securityStrategy := GetEnvVarOrPanic("SMTP_SECURITY_STRATEGY")

	return SMTPConfig{
		Host:             host,
		Port:             port,
		Username:         username,
		Password:         password,
		SecurityStrategy: securityStrategy,
	}, nil
}

// SendEmail sends an email using the provided SMTP configuration.
func SendEmail(message EmailMessage) error {
	config, err := GetSMTPConfig()
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	var serverAddr string
	if config.SecurityStrategy == "tls" {
		serverAddr = fmt.Sprintf("%s:%s", config.Host, config.Port)
	} else {
		serverAddr = fmt.Sprintf("%s:%s", config.Host, "587") // Default to 587 for non-secure connections
	}

	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", message.To, message.Subject, message.Body)

	err = smtp.SendMail(serverAddr, auth, config.Username, append(strings.Split(message.To, ","), message.CC...), []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
