package models

import (
	"github.com/Abhishek-Mali-Simform/SendMailGolang/errors"
)

type MailGun struct {
	email         Email
	domainName    string
	privateAPIKey string
}

func (email *MailGun) MailGun(sender, subject, body, recipient string) (err error) {
	if len(email.GetPrivateAPIKey()) != 0 {
		email.setEmail(sender, subject, body, recipient)
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
func (email *MailGun) setEmail(sender, subject, body, recipient string) {
	if len(sender) != 0 && len(subject) != 0 && len(body) != 0 && len(recipient) != 0 {
		email.email.Email(sender, subject, body, recipient)
	}
}

func (email *MailGun) GetEmail() map[string]string {
	var eMail = map[string]string{}
	eMail["sender"] = email.email.GetSender()
	eMail["recipient"] = email.email.GetRecipient()
	eMail["body"] = email.email.GetBody()
	eMail["subject"] = email.email.GetSubject()
	return eMail
}

func (email *MailGun) GetDomainName() string {
	return email.domainName
}

func (email *MailGun) GetPrivateAPIKey() string {
	return email.privateAPIKey
}
