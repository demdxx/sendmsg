package email

import (
	"fmt"

	"github.com/ainsleyclark/go-mail/drivers"
	"github.com/ainsleyclark/go-mail/mail"

	"github.com/demdxx/sendmsg/sender/email/customsmtp"
)

// ErrUnknownDriver is returned when the driver is not registered
var ErrUnknownDriver = fmt.Errorf("unknown driver")

var driverList = map[string]func(mail.Config) (mail.Mailer, error){
	"mailgun":   drivers.NewMailgun,
	"sparkpost": drivers.NewSparkPost,
	"smtp":      customsmtp.NewSMTP,
	"sendgrid":  drivers.NewSendGrid,
	"postal":    drivers.NewPostal,
	"postmark":  drivers.NewPostmark,
}
