package SendMail

import (
	"context"
	"github.com/Abhishek-Mali-Simform/SendMailGolang/models"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

func ByMailGun(email models.MailGun) (msg, id string, err error) {
	mg := mailgun.NewMailgun(email.GetDomainName(), email.GetPrivateAPIKey())
	sender := email.GetSender()
	subject := email.GetSubject()
	body := email.GetBody()
	recipient := email.GetRecipient()
	message := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return mg.Send(ctx, message)
}
