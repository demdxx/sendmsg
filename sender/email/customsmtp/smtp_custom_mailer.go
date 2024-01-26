package customsmtp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"mime/quotedprintable"
	"net/http"
	"net/smtp"
	"net/textproto"
	"strconv"
	"strings"

	"github.com/ainsleyclark/go-mail/mail"
)

type smtpClient struct {
	cfg  mail.Config
	send smtpSendFunc
}

// smtpSendFunc defines the function for ending
// SMTP mail.Transmissions.
type smtpSendFunc func(addr string, a smtp.Auth, from string, to []string, msg []byte) error

// NewSMTP creates a new smtp client. Configuration
// is validated before initialisation.
func NewSMTP(cfg mail.Config) (mail.Mailer, error) {
	if cfg.URL == "" {
		return nil, errors.New("driver requires a url")
	}
	if cfg.FromAddress == "" {
		return nil, errors.New("driver requires from address")
	}
	if cfg.FromName == "" {
		return nil, errors.New("driver requires from name")
	}
	if cfg.Password == "" {
		return nil, errors.New("driver requires a password")
	}
	return &smtpClient{
		cfg:  cfg,
		send: smtp.SendMail,
	}, nil
}

// Send mail via plain SMTP. mail.Transmissions are validated
// before sending and attachments are added. Returns
// an error upon failure.
func (m *smtpClient) Send(t *mail.Transmission) (mail.Response, error) {
	err := t.Validate()
	if err != nil {
		return mail.Response{}, err
	}

	auth := smtp.PlainAuth("", m.cfg.FromAddress, m.cfg.Password, m.cfg.URL)
	mailData, err := m.bytes(t)
	if err != nil {
		return mail.Response{}, err
	}
	err = m.send(m.cfg.URL+":"+strconv.Itoa(m.cfg.Port), auth, m.cfg.FromAddress, m.getTo(t), mailData)
	if err != nil {
		return mail.Response{}, err
	}

	return mail.Response{
		StatusCode: http.StatusOK,
		Message:    "Email sent successfully",
	}, nil
}

func (m *smtpClient) getTo(t *mail.Transmission) []string {
	var to []string
	to = append(t.Recipients, t.CC...)
	to = append(to, t.BCC...)
	return to
}

func (m *smtpClient) bytes(t *mail.Transmission) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	body.WriteString("MIME-Version: 1.0\n")
	body.WriteString(fmt.Sprintf("Subject: %s\n", t.Subject))
	body.WriteString(fmt.Sprintf("To: %s\n", strings.Join(t.Recipients, ",")))
	if t.HasCC() {
		body.WriteString(fmt.Sprintf("CC: %s\n", strings.Join(t.CC, ",")))
	}
	body.WriteString("Content-Type: multipart/alternative; boundary=" + writer.Boundary() + "\n\n")

	var (
		part io.Writer
		err  error
	)
	if t.PlainText != "" {
		part, err = createQuoteTypePart(writer, "text/plain")
		if err != nil {
			return nil, err
		}
		_, _ = part.Write([]byte(t.PlainText))
	}

	if t.HTML != "" {
		part, err = createQuoteTypePart(writer, "text/html")
		if err != nil {
			return nil, err
		}
		_, _ = part.Write([]byte(t.HTML))
	}

	for _, attachment := range t.Attachments {
		part, err = writer.CreatePart(textproto.MIMEHeader{
			"Content-Type":              {attachment.Mime()},
			"Content-Transfer-Encoding": {"base64"},
			"Content-Disposition":       {"attachment; filename=" + attachment.Filename},
		})
		if err != nil {
			return nil, err
		}
		_, _ = part.Write([]byte(attachment.B64()))
	}

	_ = writer.Close()

	return body.Bytes(), nil
}

// https://github.com/domodwyer/mailyak/blob/master/attachments.go#L142
func createQuoteTypePart(writer *multipart.Writer, contentType string) (part io.Writer, err error) {
	header := textproto.MIMEHeader{
		"Content-Type":              []string{contentType},
		"Content-Transfer-Encoding": []string{"quoted-printable"},
	}

	part, err = writer.CreatePart(header)
	if err != nil {
		return nil, err
	}
	part = quotedprintable.NewWriter(part)
	return part, err
}
