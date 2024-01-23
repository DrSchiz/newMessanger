package functions

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(code string, email string) {
	username := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")

	subject := "Verification code"
	body := "Here's your verification code from messanger:\n" + code

	from := os.Getenv("EMAIL_USERNAME")
	to := []string{
		email,
	}

	message := fmt.Sprintf("From: %s\r\n", from)
	message += fmt.Sprintf("To: %s\r\n", to)
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	auth := smtp.PlainAuth("", username, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, []byte(message))
	if err != nil {
		log.Println(err.Error())
		return
	}
}
