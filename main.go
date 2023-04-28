package main

import (
	"fmt"
	"github.com/Abhishek-Mali-Simform/SendMailGolang/SendMail"
	"github.com/Abhishek-Mali-Simform/SendMailGolang/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal("Error loading .env file", envError)
	}
}

func main() {
	//GridEmailExample()
	MailGunByExample()
}

func GridEmailExample() {
	var email models.GridEmail
	errorSettingKey := email.SetAPIKey(os.Getenv("SENDGRID_API_KEY"))
	if errorSettingKey != nil {
		log.Println(errorSettingKey)
	}
	errorSettingEmail := email.GridEmail(
		os.Getenv("FROM_NAME"),
		os.Getenv("FROM_EMAIL"),
		"Test Purpose",
		"Abhishek Mali",
		"abhishek.m@simformsolutions.com",
		"plain-text-content",
		"<h1>Hello World</h1>",
	)
	if errorSettingEmail != nil {
		log.Println(errorSettingEmail)
	}
	response, sendMailError := SendMail.BySendGrid(email)
	if sendMailError != nil {
		log.Println(sendMailError)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func MailGunByExample() {
	var email models.MailGun
	createAPIErr := email.SetAPIKey(os.Getenv("PRIVATE_KEY"))
	if createAPIErr != nil {
		log.Println(createAPIErr)
	}
	email.SetDomainName(os.Getenv("DOMAIN_NAME"))
	errorSettingEmail := email.MailGun("abhishek.m@simformsolutions.com", "Congratulations ", "<h1>This is the test</h1>", "kishan.m@simformsolutions.com")
	if errorSettingEmail != nil {
		log.Println(errorSettingEmail)
	}
	resp, id, err := SendMail.ByMailGun(email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
