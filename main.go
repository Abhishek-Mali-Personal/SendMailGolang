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
	var email models.GridEmail
	errorSettingKey := email.SetSendGridAPIKey(os.Getenv("SENDGRID_API_KEY"))
	if errorSettingKey != nil {
		log.Println(errorSettingKey)
	}
	email.SetFromName(os.Getenv("FROM_NAME"))
	email.SetFromEmail(os.Getenv("FROM_EMAIL"))
	errorSettingEmail := email.Email(
		"",
		"",
		"Test Purpose",
		"Abhishek Mali",
		"abhishek.m@simformsolutions.com",
		"plain-text-content",
		"<h1>Hello World</h1>",
	)
	if errorSettingEmail != nil {
		log.Println(errorSettingEmail)
	}
	response, sendMailError := SendMail.ByGrid(email)
	if sendMailError != nil {
		log.Println(sendMailError)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
