package emails

import (
	"errors"
	"log"
	"strconv"

	"github.com/wneessen/go-mail"
	"provar.se/webapi/lib/config"
)

var (
	cachedSender *EmailSender
)

// EmailSender is a helper struct for sending email messages
type EmailSender struct {
	mailClient  *mail.Client
	fromAddress string
}

// EmailTemplate represents a template for an email message
type EmailTemplate interface {
	SubjectText(map[string]string) string
	ContentText(map[string]string) string
	ContentHTML(map[string]string) string
}

// EmailOptions represents the options for sending an email
type EmailOptions struct {
	RecvAddress string
	SubjectData map[string]string
	ContentData map[string]string
}

// Setup stores the SMTP URI and the email address to send emails from.
// The smtp server uri string is expected to be in the following format:
//
//	smtp://<username>:<password>@<host>:<port>
func Setup() error {
	if cachedSender != nil {
		return nil
	}
	cfg := config.Get()
	smtpPort, err := strconv.Atoi(cfg.Email.EmailServer.Port())
	if err != nil {
		return err
	}
	// Note: we assume we will always use a password for authentication
	smtpPass, ok := cfg.Email.EmailServer.User.Password()
	if !ok {
		return errors.New("SMTP password not set")
	}
	mailClient, err := mail.NewClient(
		cfg.Email.EmailServer.Hostname(),
		mail.WithTLSPortPolicy(mail.TLSOpportunistic),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Email.EmailServer.User.Username()),
		mail.WithPassword(smtpPass),
		mail.WithPort(smtpPort),
	)
	if err != nil {
		return err
	}
	cachedSender = &EmailSender{
		fromAddress: cfg.Email.EmailFrom,
		mailClient:  mailClient,
	}
	return nil
}

// Sender sends a magic link to the user's email.
func Sender() *EmailSender {
	if cachedSender == nil {
		log.Fatalln("Cannot send emails before setting up the email sender")
	}
	return cachedSender
}

// Send sends an email to the given email address with the given options.
func (s *EmailSender) Send(template EmailTemplate, options *EmailOptions) error {
	subject := template.SubjectText(options.SubjectData)
	contentText := template.ContentText(options.ContentData)
	contentHTML := template.ContentHTML(options.ContentData)
	message, err := s.createMessage(options.RecvAddress, subject, contentText, contentHTML)
	if err != nil {
		return err
	}
	return s.mailClient.DialAndSend(message)
}

// createMessage creates a multipart email message with the given parameters
func (s *EmailSender) createMessage(toAddress, subject, contentText, contentHTML string) (*mail.Msg, error) {
	m := mail.NewMsg()
	if err := m.From(s.fromAddress); err != nil {
		return nil, err
	}
	if err := m.To(toAddress); err != nil {
		return nil, err
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, contentText)
	m.SetBodyString(mail.TypeTextHTML, contentHTML)
	return m, nil
}
