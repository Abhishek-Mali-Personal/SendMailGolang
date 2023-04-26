package SendMail

import (
	"github.com/Abhishek-Mali-Simform/SendMailGolang/errors"
	"github.com/Abhishek-Mali-Simform/SendMailGolang/models"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ByGrid(email models.GridEmail) (response *rest.Response, err error) {
	if len(email.GetSendGridAPIKey()) != 0 {
		eMail := email.GetEmail()
		from := mail.NewEmail(email.GetSenderName(), eMail["sender"])
		subject := eMail["subject"]
		to := mail.NewEmail(email.GetRecipientName(), eMail["recipient"])
		var (
			plainTextContent string
			htmlContent      string
		)
		switch email.GetContentType() {
		case "plain-text-content":
			plainTextContent = eMail["body"]
		case "html-content":
			htmlContent = eMail["body"]
		}
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(email.GetSendGridAPIKey())
		response, err = client.Send(message)
	} else {
		err = errors.EmptyAPIKeyError
	}
	return
}
