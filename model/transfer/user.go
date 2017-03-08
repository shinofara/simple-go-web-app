package transfer

import (
	"github.com/shinofara/simple-go-web-app/http/context"
)

// SendActivationEmail アクティベーション要求メールを送信
func SendActivationEmail(ctx context.Context) error {
	m := NewMailer(ctx)
	return m.Send(
		"from@example.com",
		[]string{"to@example.com"},
		[]string{},
		"title",
		"body",
	)
}
