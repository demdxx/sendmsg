package email

import (
	"github.com/ainsleyclark/go-mail/mail"
	"github.com/pkg/errors"
)

// ErrUnknownDriver is returned when the driver is not registered
var ErrMailNotConfigured = errors.New("mailer not configured")

// Config for mailer
type Config = mail.Config

// Mailer is the interface that wraps the basic Send method
type Mailer = mail.Mailer

// Options for email sender
type Options struct {
	mailer  Mailer
	cfg     *Config
	driver  string
	vars    map[string]any
	headers map[string]string
}

// Option for email sender
type Option func(*Options)

// WithMailer sets mailer
func WithMailer(mailer Mailer) Option {
	return func(o *Options) {
		o.mailer = mailer
		if o.cfg != nil && o.mailer != nil {
			panic("cannot use WithConfig and WithMailer together")
		}
	}
}

// WithConfig sets mailer config
func WithConfig(driver string, cfg *Config) Option {
	return func(o *Options) {
		o.cfg = cfg
		o.driver = driver
		if o.cfg != nil && o.mailer != nil {
			panic("cannot use WithConfig and WithMailer together")
		}
	}
}

// WithVars sets default vars
func WithVars(vars map[string]any) Option {
	return func(o *Options) {
		o.vars = vars
	}
}

// WithHeaders sets default headers
func WithHeaders(headers map[string]string) Option {
	return func(o *Options) {
		o.headers = headers
	}
}

// Mailer returns mailer
func (o *Options) Mailer() (mail.Mailer, error) {
	if o.mailer != nil {
		return o.mailer, nil
	}
	if o.cfg == nil {
		return nil, ErrMailNotConfigured
	}
	curDriver := "smtp"
	if o.driver != "" {
		curDriver = o.driver
	}
	connector, ok := driverList[curDriver]
	if !ok {
		return nil, errors.Wrap(ErrUnknownDriver, curDriver)
	}
	return connector(*o.cfg)
}
