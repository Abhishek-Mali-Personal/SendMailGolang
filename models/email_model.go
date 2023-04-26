package models

type Email struct {
	sender    string
	subject   string
	body      string
	recipient string
}

func (email *Email) Email(sender, subject, body, recipient string) {
	email.setSender(sender)
	email.setSubject(subject)
	email.setBody(body)
	email.setRecipient(recipient)
}

func (email *Email) setRecipient(recipient string) {
	if len(recipient) != 0 {
		email.recipient = recipient
	}
}

func (email *Email) setBody(body string) {
	if len(body) != 0 {
		email.body = body
	}
}

func (email *Email) setSubject(subject string) {
	if len(subject) != 0 {
		email.subject = subject
	}
}

func (email *Email) setSender(sender string) {
	if len(sender) != 0 {
		email.sender = sender
	}
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
