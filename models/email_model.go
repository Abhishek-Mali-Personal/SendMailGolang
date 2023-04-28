package models

import "github.com/Abhishek-Mali-Simform/SendMailGolang/errors"

type Email struct {
	sender    string
	subject   string
	body      string
	recipient string
	apiKey    string
}

func (email *Email) Email(sender, subject, body, recipient string) {
	email.setSender(sender)
	email.setSubject(subject)
	email.setBody(body)
	email.setRecipient(recipient)
}

func (email *Email) setRecipient(recipient string) {
	if recipient != "" {
		email.recipient = recipient
	}
}

func (email *Email) setBody(body string) {
	if body != "" {
		email.body = body
	}
}

func (email *Email) setSubject(subject string) {
	if subject != "" {
		email.subject = subject
	}
}

func (email *Email) setSender(sender string) {
	if sender != "" {
		email.sender = sender
	}
}

func (email *Email) SetAPIKey(apiKey string) (err error) {
	if apiKey != "" {
		email.apiKey = apiKey
	} else {
		err = errors.EmptyAPIKeyError
	}
	return
}

func (email *Email) GetAPIKey() string {
	return email.apiKey
}

func (email *Email) GetSender() string {
	return email.sender
}

func (email *Email) GetSubject() string {
	return email.subject
}

func (email *Email) GetBody() string {
	return email.body
}

func (email *Email) GetRecipient() string {
	return email.recipient
}
