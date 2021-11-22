package helpers

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func HandleEmailService(email string) {

	godotenv.Load(".env")

	sender := os.Getenv("email")
	Senderauth := os.Getenv("password")

	to := []string{email}

	from := sender
	password := Senderauth

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Our Golang Email\n"
	body := `Email from GO-Authentication ${email}`
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
