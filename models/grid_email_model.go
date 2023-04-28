package models

import (
	"github.com/Abhishek-Mali-Simform/SendMailGolang/errors"
)

type GridEmail struct {
	Email
	senderName    string
	recipientName string
	contentType   string
}

func (email *GridEmail) GridEmail(senderName, senderEmail, subject, recipientName, recipientEmail, contentType, body string) (err error) {
	if email.GetAPIKey() != "" {
		email.SetSenderName(senderName)
		email.setRecipientName(recipientName)
		email.setEmail(senderEmail, subject, body, recipientEmail)
		email.setContentType(contentType)
		err = nil
	} else {
		err = errors.EmptyAPIKeyError
	}
	return
}

func (email *GridEmail) SetSenderName(senderName string) {
	if senderName != "" {
		email.senderName = senderName
	} else if email.GetSenderName() == "" {
		email.senderName = "Unknown"
	}
}

func (email *GridEmail) setRecipientName(recipientName string) {
	if recipientName != "" {
		email.recipientName = recipientName
	} else {
		email.recipientName = "Unknown"
	}
}

func (email *GridEmail) setContentType(contentType string) {
	if contentType != "" {
		email.contentType = contentType
	}
}

func (email *GridEmail) setEmail(sender, subject, body, recipient string) {
	if sender != "" && subject != "" && body != "" && recipient != "" {
		email.Email.Email(sender, subject, body, recipient)
	}
}

func (email *GridEmail) GetEmail() map[string]string {
	var eMail = map[string]string{}
	eMail["sender"] = email.GetSender()
	eMail["recipient"] = email.GetRecipient()
	eMail["body"] = email.GetBody()
	eMail["subject"] = email.GetSubject()
	return eMail
}

func (email *GridEmail) GetContentType() string {
	return email.contentType
}

func (email *GridEmail) GetSenderName() string {
	return email.senderName
}

func (email *GridEmail) GetRecipientName() string {
	return email.recipientName
}
