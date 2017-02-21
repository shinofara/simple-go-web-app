package transfer

import "net/smtp"

func SendActivationEmail() error {
	// 適宜変更してください
	smtpServer := "localhost:1025"
	mailAddr := "test@example.com"
	receivers := []string{mailAddr}

	msg := "" +
		"From:" + mailAddr + "\r\n" +
		"To:" + mailAddr + "\r\n" +
		"Cc:" + mailAddr + "\r\n" +
		"Subject:SMTP Test\r\n" +
		"\r\n" +
		"This is a test mail."

	//auth := smtp.PlainAuth("", user, pass, "smtp.example.com")
	//auth := smtp.CRAMMD5Auth(user, pass)
	return smtp.SendMail(smtpServer, nil, mailAddr, receivers, []byte(msg))
}
