package SendMail

import (
	"context"
	"github.com/Abhishek-Mali-Simform/SendMailGolang/models"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

func ByMailGun(email models.MailGun) (msg, id string, err error) {
	mg := mailgun.NewMailgun(email.GetDomainName(), email.GetAPIKey())
	eMail := email.GetEmail()
	sender := eMail["sender"]
	subject := eMail["subject"]
	body := eMail["body"]
	recipient := eMail["recipient"]
	message := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return mg.Send(ctx, message)
}
