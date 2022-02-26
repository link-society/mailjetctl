package client

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v3"
	"golang.org/x/xerrors"
)

func SendMail(msg *Message) error {
	publicKey := os.Getenv(MJ_APIKEY_PUBLIC_VAR)
	secretKey := os.Getenv(MJ_APIKEY_PRIVATE_VAR)

	if len(publicKey) == 0 {
		return xerrors.Errorf("Missing environment variable: %s", MJ_APIKEY_PUBLIC_VAR)
	}

	if len(secretKey) == 0 {
		return xerrors.Errorf("Missing environment variable: %s", MJ_APIKEY_PRIVATE_VAR)
	}

	mj := mailjet.NewMailjetClient(publicKey, secretKey)
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: msg.Sender.Address,
					Name:  msg.Sender.Name,
				},
				To:       toMailjetRecipients(msg.Recipients),
				Subject:  msg.Subject,
				TextPart: msg.TextPart,
				HTMLPart: msg.HtmlPart,
			},
		},
	}

	res, err := mj.SendMailV31(&messages)
	if err != nil {
		return xerrors.Errorf("Unable to send mail: %w", err)
	}

	fmt.Println(res)
	return nil
}

func toMailjetRecipients(recipients []mail.Address) *mailjet.RecipientsV31 {
	mjRecipients := make(mailjet.RecipientsV31, len(recipients))

	for i, recipient := range recipients {
		mjRecipients[i] = mailjet.RecipientV31{
			Email: recipient.Address,
			Name:  recipient.Name,
		}
	}

	return &mjRecipients
}
