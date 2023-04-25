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
		from := mail.NewEmail(email.GetFromName(), email.GetFromEmail())
		subject := email.GetSubject()
		to := mail.NewEmail(email.GetToName(), email.GetToEmail())
		plainTextContent := email.GetPlainTextContent()
		htmlContent := email.GetHtmlContent()
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(email.GetSendGridAPIKey())
		response, err = client.Send(message)
	} else {
		err = errors.EmptySendGridAPIKeyError
	}
	return
}
