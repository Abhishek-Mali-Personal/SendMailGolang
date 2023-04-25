package models

import (
	"go.mod/errors"
)

type GridEmail struct {
	fromEmail        string
	fromName         string
	subject          string
	toName           string
	toEmail          string
	sendGridAPIKey   string
	plainTextContent string
	htmlContent      string
}

func (email *GridEmail) Email(fromName, fromEmail, subject, toName, toEmail, contentType, content string) (err error) {
	if len(email.sendGridAPIKey) != 0 {
		email.SetFromName(fromName)
		email.SetFromEmail(fromEmail)
		email.setSubject(subject)
		email.setToName(toName)
		email.setToEmail(toEmail)
		switch contentType {
		case "plain-text-content":
			email.setPlainTextContent(content)
		case "html-content":
			email.setHtmlContent(content)
		}
		err = nil
	} else {
		err = errors.EmptySendGridAPIKeyError
	}
	return
}

func (email *GridEmail) setPlainTextContent(plainTextContent string) {
	if len(plainTextContent) != 0 {
		email.plainTextContent = plainTextContent
	}
}

func (email *GridEmail) setHtmlContent(htmlContent string) {
	if len(htmlContent) != 0 {
		email.htmlContent = htmlContent
	}
}

func (email *GridEmail) SetFromEmail(fromEmail string) {
	if len(fromEmail) != 0 {
		email.fromEmail = fromEmail
	}
}

func (email *GridEmail) SetFromName(fromName string) {
	if len(fromName) != 0 {
		email.fromName = fromName
	} else if len(email.GetFromName()) == 0 {
		email.fromName = "Unknown"
	}
}

func (email *GridEmail) setToEmail(toEmail string) {
	if len(toEmail) != 0 {
		email.toEmail = toEmail
	}
}

func (email *GridEmail) setToName(toName string) {
	if len(toName) != 0 {
		email.toName = toName
	} else {
		email.toName = "Unknown"
	}
}

func (email *GridEmail) setSubject(subject string) {
	if len(subject) != 0 {
		email.subject = subject
	} else {
		email.subject = "Welcome to Grid GridEmail Functionality Thanks for using it."
	}
}

func (email *GridEmail) SetSendGridAPIKey(sendGridAPIKey string) (err error) {
	if len(sendGridAPIKey) != 0 {
		email.sendGridAPIKey = sendGridAPIKey
		err = nil
	} else {
		err = errors.EmptySendGridAPIKeyError
	}
	return
}

func (email *GridEmail) GetSendGridAPIKey() string {
	return email.sendGridAPIKey
}

func (email *GridEmail) GetSubject() string {
	return email.subject
}

func (email *GridEmail) GetFromName() string {
	return email.fromName
}

func (email *GridEmail) GetFromEmail() string {
	return email.fromEmail
}

func (email *GridEmail) GetToName() string {
	return email.toName
}

func (email *GridEmail) GetToEmail() string {
	return email.toEmail
}

func (email *GridEmail) GetPlainTextContent() string {
	return email.plainTextContent
}

func (email *GridEmail) GetHtmlContent() string {
	return email.htmlContent
}
