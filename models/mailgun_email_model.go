package models

import "github.com/Abhishek-Mali-Simform/SendMailGolang/errors"

type MailGun struct {
	sender        string
	subject       string
	body          string
	recipient     string
	domainName    string
	privateAPIKey string
}

func (email *MailGun) MailGun(sender, subject, body, recipient string) (err error) {
	if len(email.GetPrivateAPIKey()) != 0 {
		email.setSender(sender)
		email.setSubject(subject)
		email.setBody(body)
		email.setRecipient(recipient)
		err = nil
	} else {
		err = errors.EmptyAPIKeyError
	}
	return
}

func (email *MailGun) SetDomainName(domainName string) {
	if len(domainName) != 0 {
		email.domainName = domainName
	}
}

func (email *MailGun) SetPrivateAPIKey(privateAPIKey string) {
	if len(privateAPIKey) != 0 {
		email.privateAPIKey = privateAPIKey
	}
}

func (email *MailGun) setRecipient(recipient string) {
	if len(recipient) != 0 {
		email.recipient = recipient
	}
}

func (email *MailGun) setBody(body string) {
	if len(body) != 0 {
		email.body = body
	}
}

func (email *MailGun) setSubject(subject string) {
	if len(subject) != 0 {
		email.subject = subject
	}
}

func (email *MailGun) setSender(sender string) {
	if len(sender) != 0 {
		email.sender = sender
	}
}
func (email *MailGun) GetDomainName() string {
	return email.domainName
}

func (email *MailGun) GetPrivateAPIKey() string {
	return email.privateAPIKey
}

func (email *MailGun) GetSender() string {
	return email.sender
}

func (email *MailGun) GetSubject() string {
	return email.subject
}

func (email *MailGun) GetBody() string {
	return email.body
}

func (email *MailGun) GetRecipient() string {
	return email.recipient
}
