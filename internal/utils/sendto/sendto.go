package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/albertbui010/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func getSMTPConfig() (host, port, username, password string) {
	smtp := global.Config.SMTP
	return smtp.SMTPHost, smtp.SMTPPort, smtp.SMTPUsername, smtp.SMTPPassword
}

// Send text
func SendTextEmailOTP(to []string, from string, otp string) error {
	SMTPHost, SMTPPort, SMTPUsername, SMTPPassword := getSMTPConfig()

	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	}

	messageMail := BuildMessage(contentEmail)

	// Send SMTP
	authtion := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, authtion, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send faild::", zap.Error(err))
		return err
	}

	return nil
}

// Send template
func SendTemplateEmailOTP(to []string, from string, templateName string, dataTemplate map[string]interface{}) error {
	htmlBody, err := getMailTemplate(templateName, dataTemplate)
	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(tempalateName string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(tempalateName).ParseFiles("template-email/" + tempalateName))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	SMTPHost, SMTPPort, SMTPUsername, SMTPPassword := getSMTPConfig()

	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)
	authtion := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, authtion, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send faild::", zap.Error(err))
		return err
	}

	return nil
}
