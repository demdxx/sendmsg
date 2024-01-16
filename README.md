# Message sender abstraction

## Usage

```sh
go get github.com/demdxx/sendmsg
```

```go
package main

import (
  "fmt"
  "log"

  "github.com/demdxx/sendmsg"
  "github.com/demdxx/sendmsg/email"
  "github.com/demdxx/sendmsg/template"
)

func main() {
  smtpSender, err := email.New(
    email.WithConfig(email.Config{
      URL:         "smtp.gmail.com",
      FromAddress: "hello@mail.com",
      FromName:    "Mailer",
      Password:    "password",
      Port:        587,
    }),
    // Default vars for all messages
    email.WithVars(map[string]any{
      "company": "My Awesome Company",
    }),
  )

  if err != nil {
    log.Fatal(err)
  }

  // ...
  templateStorage := template.NewDefaultStorage(registerTmpl, resetPasswordTmpl)
  messanger := sendmsg.NewDefaultMessanger(templateStorage).
                       RegisterSender("email", smtpSender)

  // ...
  err := messanger.Send(ctx,
    sendmsg.WithTemplate("register"),
    sendmsg.WithSender("email"),
    sendmsg.WithVars(map[string]any{
      "name": "Mr. Smith", 
      "email": "smith@mail.com",
    }),
  )
}
```

## TODO

- [ ] Add Telegram sender implementation
- [ ] Add Slack sender implementation
- [ ] Add WhatsApp sender implementation
- [ ] Add Viber sender implementation
- [ ] Add Facebook sender implementation
- [ ] Add Equipos sender implementation
- [ ] Add Line sender implementation
- [ ] Add WeChat sender implementation
- [ ] Add Skype sender implementation
- [ ] Add Signal sender implementation
- [ ] Add Discord sender implementation
- [ ] Add Snapchat sender implementation
- [ ] Add SMS sender implementation
- [X] Add EMail sender implementation
