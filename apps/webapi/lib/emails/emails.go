package emails

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/smtp"
	"net/url"

	"provar.se/webapi/lib/config"
)

var (
	cachedSender *EmailSender
)

// EmailSender is a helper struct for sending email messages
type EmailSender struct {
	emailServer *url.URL
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
func Setup(emailConfig config.EmailConfig) error {
	if cachedSender != nil {
		return nil
	}
	cachedSender = &EmailSender{
		emailServer: emailConfig.EmailServer,
		fromAddress: emailConfig.EmailFrom,
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
	user := s.emailServer.User.Username()
	pass, _ := s.emailServer.User.Password()
	recv := []string{options.RecvAddress}
	auth := smtp.PlainAuth("", user, pass, s.emailServer.Host)
	err = smtp.SendMail(s.emailServer.Host+":"+s.emailServer.Port(), auth, s.fromAddress, recv, message)
	if err != nil {
		return err
	}
	return nil
}

// createMessage creates a multipart email message with the given parameters
func (s *EmailSender) createMessage(toAddress, subject, contentText, contentHTML string) ([]byte, error) {
	msg := bytes.NewBuffer(nil)
	multipartWriter := multipart.NewWriter(msg)
	// Add text part
	textPart, err := multipartWriter.CreatePart(map[string][]string{"Content-Type": {"text/plain; charset=UTF-8"}})
	if err != nil {
		return nil, err
	}
	_, err = textPart.Write([]byte(contentText))
	if err != nil {
		return nil, err
	}
	// Add HTML part
	htmlPart, err := multipartWriter.CreatePart(map[string][]string{"Content-Type": {"text/html; charset=UTF-8"}})
	if err != nil {
		return nil, err
	}
	_, err = htmlPart.Write([]byte(contentHTML))
	if err != nil {
		return nil, err
	}
	// Close multipart writer
	err = multipartWriter.Close()
	if err != nil {
		return nil, err
	}
	// Prepare email message
	headers := make(map[string]string)
	headers["From"] = s.fromAddress
	headers["To"] = toAddress
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "multipart/alternative; boundary=" + multipartWriter.Boundary()
	// Write email headers
	for key, value := range headers {
		_, err := msg.WriteString(key + ": " + value + "\r\n")
		if err != nil {
			return nil, err
		}
	}
	_, err = msg.WriteString("\r\n")
	if err != nil {
		return nil, err
	}
	return msg.Bytes(), nil
}
