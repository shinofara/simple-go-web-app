package context

import (
	"github.com/shinofara/simple-go-web-app/session"		
	"fmt"
	"context"
)

const CtxKeySession = "SESSION"

func SetSession(ctx context.Context, sess *session.Session) context.Context {
	return context.WithValue(ctx, CtxKeySession, sess)
}

func GetSession(ctx context.Context) (*session.Session, error) {
	sess, ok := ctx.Value(CtxKeySession).(*session.Session)
	if ok {
		return sess, nil
	}

	return nil, fmt.Errorf("Failed to get SESSION from context")
}

func MustGetSession(ctx context.Context) *session.Session {
	sess, err := GetSession(ctx)
	if err != nil {


		
		panic(err)
	}
	return sess
}
