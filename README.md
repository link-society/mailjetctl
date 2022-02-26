# mailjetctl

Small binary used to send mail via [Mailjet](https://mailjet.com).

## Usage

You will need to configure your API key with the following environment variables:

```bash
export MJ_APIKEY_PUBLIC="my public key"
export MJ_APIKEY_PRIVATE="my secret key"
```

Then, you can use the following options:

```
usage: mailjetctl [--help] [--from FROM] [--recipient RECIPIENT [RECIPIENT ...]] [--recipients-file RECIPIENTS-FILE] [--subject SUBJECT] [--subject-file SUBJECT-FILE] [--text-file TEXT-FILE] [--html-file HTML-FILE]

Command Line Interface for Mailjet

optional arguments:
  --help, -h
  --from FROM, -f FROM                                   Sender email address
  --recipient RECIPIENT, -r RECIPIENT                    Recipient email address
  --recipients-file RECIPIENTS-FILE, -R RECIPIENTS-FILE  Path to file containing recipients (one per line)
  --subject SUBJECT, -s SUBJECT                          Email subject
  --subject-file SUBJECT-FILE, -S SUBJECT-FILE           Path to file containing the email subject
  --text-file TEXT-FILE, -T TEXT-FILE                    Path to file containing the email text part
  --html-file HTML-FILE, -H HTML-FILE                    Path to file containing the email HTML part
```

## License

This project is released under the terms of the [MIT License](./LICENSE.txt).
