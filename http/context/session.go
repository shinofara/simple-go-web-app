package context

import (
	"github.com/shinofara/simple-go-web-app/http/session/core"
	"fmt"
	"context"
)

// CtxKeySession コンテキスト内にSessionを管理するKey
const CtxKeySession = contextKey("SESSION")

// SetSessionStore sets sesion store to context.
func SetSessionStore(ctx context.Context, store *core.Store) context.Context {
	return context.WithValue(ctx, CtxKeySession, store)
}

// GetSessionStore sets sesion store from context.
func GetSessionStore(ctx context.Context) (*core.Store, error) {
	sess, ok := ctx.Value(CtxKeySession).(*core.Store)
	if ok {
		return sess, nil
	}

	return nil, fmt.Errorf("Failed to get SESSION from context")
}

// MustGetSessionStore 確実にsession storeをsyutoku
func MustGetSessionStore(ctx context.Context) *core.Store {
	sess, err := GetSessionStore(ctx)
	if err != nil {
		panic(err)
	}
	return sess
}
