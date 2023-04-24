package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/densmart/users-manager/internal/logger"
	"github.com/spf13/viper"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"strings"
	"text/template"
)

type Mailer struct {
	Auth smtp.Auth
	Addr string
	From string
}

func NewMailer() *Mailer {
	auth := smtp.PlainAuth("", viper.GetString("mail.user"), viper.GetString("mail.password"),
		viper.GetString("mail.host"))

	return &Mailer{
		Addr: viper.GetString("mail.host") + ":" + viper.GetString("mail.port"),
		From: viper.GetString("mail.from"),
		Auth: auth,
	}
}

func (m *Mailer) Send(message *MailMessage) error {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("Subject: SmartBooking :: %s\n", message.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(message.To, ",")))

	if len(message.Copy) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(message.Copy, ",")))
	}

	buf.WriteString("MIME-version: 1.0;\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()

	if len(message.Attachments) > 0 {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s;\n\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	}

	buf.WriteString("Content-Type: text/html; charset=\"UTF-8\";\n\n")
	buf.WriteString(message.Body)

	if len(message.Attachments) > 0 {
		for k, v := range message.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	if err := smtp.SendMail(m.Addr, m.Auth, m.From, message.To, buf.Bytes()); err != nil {
		return err
	}
	return nil
}

type MailMessage struct {
	To          []string
	Copy        []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func (m *MailMessage) Attach(filename string, body []byte) {
	m.Attachments[filename] = body
}

func ParseMailTemplate(name string, params map[string]interface{}) (string, error) {
	templateDir := viper.GetString("app.template_dir")
	files := []string{
		templateDir + name + ".html",
		templateDir + "header.html",
		templateDir + "footer.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		logger.Errorf("failed parsing template files: %s", err.Error())
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, params); err != nil {
		return "", err
	}
	return buf.String(), err
}
