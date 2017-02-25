package transfer

import (
	"github.com/shinofara/simple-go-web-app/context"	
)

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
