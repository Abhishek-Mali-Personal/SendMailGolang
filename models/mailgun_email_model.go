package models

import "github.com/Abhishek-Mali-Simform/SendMailGolang/errors"

type MailGun struct {
	Email
	domainName string
}

func (email *MailGun) MailGun(sender, subject, body, recipient string) (err error) {
	if len(email.GetAPIKey()) != 0 {
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

func (email *MailGun) setEmail(sender, subject, body, recipient string) {
	if len(sender) != 0 && len(subject) != 0 && len(body) != 0 && len(recipient) != 0 {
		email.Email.Email(sender, subject, body, recipient)
	}
}

func (email *MailGun) GetEmail() map[string]string {
	var eMail = map[string]string{}
	eMail["sender"] = email.GetSender()
	eMail["recipient"] = email.GetRecipient()
	eMail["body"] = email.GetBody()
	eMail["subject"] = email.GetSubject()
	return eMail
}

func (email *MailGun) GetDomainName() string {
	return email.domainName
}
