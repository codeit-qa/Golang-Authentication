package helpers

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func HandleEmailService(email string, code int) {

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
	body := "Email from GO-Authentication \nVerification code is : " + strconv.Itoa(code)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
