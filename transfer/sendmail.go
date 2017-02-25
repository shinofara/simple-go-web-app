package transfer

import (
	"fmt"
	"net/smtp"
	"github.com/shinofara/simple-go-web-app/config"
	"github.com/shinofara/simple-go-web-app/context"		
)

type Mailer struct {
	ctx context.Context
}

func NewMailer(ctx context.Context) *Mailer {
	return &Mailer{
		ctx: ctx,
	}
}

func (m *Mailer) Send(from string, to, cc []string, subject, body string) error {
	logger := context.MustGetLogger(m.ctx)
	smtpCfg := config.GetSMTP()
	logger.Info(fmt.Sprintf("%+v", smtpCfg))

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
