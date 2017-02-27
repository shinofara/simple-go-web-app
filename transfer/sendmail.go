package transfer

import (
	"fmt"
	"net/smtp"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/context"
)

// Mailer メール送信に関わる処理を定義
type Mailer struct {
	ctx context.Context
}

// NewMailer creates a Mailer
func NewMailer(ctx context.Context) *Mailer {
	return &Mailer{
		ctx: ctx,
	}
}

// Send メール送信
func (m *Mailer) Send(from string, to, cc []string, subject, body string) error {
	smtpCfg := config.GetSMTP()

	smtpServer := fmt.Sprintf("%s:%d", smtpCfg.Host, smtpCfg.Port)

	msg := "" +
		"From:" + from + "\r\n" +
		"To:" + to[0] + "\r\n" +
		"Subject:SMTP Test\r\n" +
		"\r\n" +
		body

	//auth := smtp.PlainAuth("", user, pass, "smtp.example.com")
	//auth := smtp.CRAMMD5Auth(user, pass)
	return smtp.SendMail(smtpServer, nil, from, to, []byte(msg))
}
