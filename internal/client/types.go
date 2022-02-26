package client

import "net/mail"

type Message struct {
	Sender     mail.Address
	Recipients []mail.Address
	Subject    string
	TextPart   string
	HtmlPart   string
}
