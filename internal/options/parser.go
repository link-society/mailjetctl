package options

import (
	"net/mail"
	"os"
	"strings"

	"github.com/hellflame/argparse"
	"golang.org/x/xerrors"

	"github.com/link-society/mailjetctl/internal/client"
)

func ParseCliArgs(args []string) (*client.Message, error) {
	msg := &client.Message{}

	parser := argparse.NewParser(
		PROGRAM_NAME, PROGRAM_DESC,
		&argparse.ParserConfig{
			DisableHelp:            true,
			DisableDefaultShowHelp: true,
		},
	)

	help := parser.Flag("h", "help", nil)

	parser.String(
		"f", "from",
		&argparse.Option{
			Positional: false,
			Required:   true,
			Help:       "Sender email address",
			Validate: func(arg string) error {
				email, err := mail.ParseAddress(arg)

				if err != nil {
					return xerrors.Errorf("Unable to parse sender address: %w", err)
				}

				msg.Sender = *email
				return nil
			},
		},
	)

	parser.Strings(
		"r", "recipient",
		&argparse.Option{
			Positional: false,
			Required:   false,
			Help:       "Recipient email address",
			Validate: func(arg string) error {
				email, err := mail.ParseAddress(arg)

				if err != nil {
					return xerrors.Errorf("Unable to parse recipient address: %w", err)
				}

				msg.Recipients = append(msg.Recipients, *email)
				return nil
			},
		},
	)

	parser.String(
		"R", "recipients-file",
		&argparse.Option{
			Positional: false,
			Required:   false,
			Help:       "Path to file containing recipients (one per line)",
			Validate: func(arg string) error {
				data, err := os.ReadFile(arg)

				if err != nil {
					return xerrors.Errorf("Unable to read recipient file: %w", err)
				}

				addresses := strings.Split(string(data), "\n")

				for _, line := range addresses {
					address := strings.TrimSpace(line)

					if len(address) > 0 {
						email, err := mail.ParseAddress(address)

						if err != nil {
							return xerrors.Errorf("Unable to parse recipient address: %w", err)
						}

						msg.Recipients = append(msg.Recipients, *email)
					}
				}

				return nil
			},
		},
	)

	parser.String(
		"s", "subject",
		&argparse.Option{
			Positional: false,
			Required:   false,
			Help:       "Email subject",
			Validate: func(arg string) error {
				msg.Subject = arg
				return nil
			},
		},
	)

	parser.String(
		"S", "subject-file",
		&argparse.Option{
			Positional: false,
			Required:   false,
			Help:       "Path to file containing the email subject",
			Validate: func(arg string) error {
				data, err := os.ReadFile(arg)

				if err != nil {
					return xerrors.Errorf("Unable to read subject file: %w", err)
				}

				msg.Subject = string(data)
				return nil
			},
		},
	)

	parser.String(
		"T", "text-file",
		&argparse.Option{
			Positional: false,
			Required:   false,
			Help:       "Path to file containing the email text part",
			Validate: func(arg string) error {
				data, err := os.ReadFile(arg)

				if err != nil {
					return xerrors.Errorf("Unable to read text part file: %w", err)
				}

				msg.TextPart = string(data)
				return nil
			},
		},
	)

	parser.String(
		"H", "html-file",
		&argparse.Option{
			Positional: false,
			Required:   false,
			Help:       "Path to file containing the email HTML part",
			Validate: func(arg string) error {
				data, err := os.ReadFile(arg)

				if err != nil {
					return xerrors.Errorf("Unable to read text part file: %w", err)
				}

				msg.HtmlPart = string(data)
				return nil
			},
		},
	)

	err := parser.Parse(args)

	if *help {
		parser.PrintHelp()
		os.Exit(0)
	}

	if err != nil {
		parser.PrintHelp()
		return nil, err
	}

	if len(msg.Recipients) == 0 {
		return nil, xerrors.New("No recipients configured")
	}

	if len(msg.Subject) == 0 {
		return nil, xerrors.New("Subject is empty")
	}

	if len(msg.TextPart) == 0 && len(msg.HtmlPart) == 0 {
		return nil, xerrors.New("Both text part and HTML part are empty")
	}

	return msg, nil
}
