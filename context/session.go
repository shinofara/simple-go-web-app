package context

import (
	"github.com/shinofara/simple-go-web-app/session"		
	"fmt"
	"context"
)

// CtxKeySession コンテキスト内にSessionを管理するKey
const CtxKeySession = "SESSION"

// SetSessionStore sets sesion store to context.
func SetSessionStore(ctx context.Context, store *session.Store) context.Context {
	return context.WithValue(ctx, CtxKeySession, store)
}

// GetSessionStore sets sesion store from context.
func GetSessionStore(ctx context.Context) (*session.Store, error) {
	sess, ok := ctx.Value(CtxKeySession).(*session.Store)
	if ok {
		return sess, nil
	}

	return nil, fmt.Errorf("Failed to get SESSION from context")
}

// MustGetSessionStore 確実にsession storeをsyutoku
func MustGetSessionStore(ctx context.Context) *session.Store {
	sess, err := GetSessionStore(ctx)
	if err != nil {
		panic(err)
	}
	return sess
}
