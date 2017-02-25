package context

import (
	"github.com/shinofara/simple-go-web-app/session"		
	"fmt"
	"context"
)

const CtxKeySession = "SESSION"

func SetSessionStore(ctx context.Context, store *session.SessionStore) context.Context {
	return context.WithValue(ctx, CtxKeySession, store)
}

func GetSessionStore(ctx context.Context) (*session.SessionStore, error) {
	sess, ok := ctx.Value(CtxKeySession).(*session.SessionStore)
	if ok {
		return sess, nil
	}

	return nil, fmt.Errorf("Failed to get SESSION from context")
}

func MustGetSessionStore(ctx context.Context) *session.SessionStore {
	sess, err := GetSessionStore(ctx)
	if err != nil {


		
		panic(err)
	}
	return sess
}
