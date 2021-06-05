package sendgrid

import (
	"fmt"
	"log"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailSenderProps struct {
	from        *mail.Email
	subject     string
	to          *mail.Email
	content     string
	htmlContent string
}

type EmailSender struct {
	client *sendgrid.Client
	props  EmailSenderProps
}

func (e *EmailSender) Send() {
	message := mail.NewSingleEmail(e.props.from, e.props.subject, e.props.to, e.props.content, e.props.htmlContent)
	response, err := e.client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func NewSendEmail(props EmailSenderProps) EmailSender {
	return EmailSender{
		client: sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
		props:  props,
	}
}
