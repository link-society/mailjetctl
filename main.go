package main

import (
	"fmt"
	"os"

	"github.com/link-society/mailjetctl/internal/client"
	"github.com/link-society/mailjetctl/internal/options"
)

func main() {
	msg, err := options.ParseCliArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = client.SendMail(msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
