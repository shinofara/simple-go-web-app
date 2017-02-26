package context

import (
	gorp	"gopkg.in/gorp.v1"
	"fmt"
	"context"
)

// SetDB sets db connection to context.
func SetDB(ctx context.Context, name string, db *gorp.DbMap) context.Context {
	return context.WithValue(ctx, generateDBContextKey(name), db)
}

// GetDB gets db connection from context.
func GetDB(ctx context.Context, name string) (*gorp.DbMap, error) {
	db, ok := ctx.Value(generateDBContextKey(name)).(*gorp.DbMap)
	if ok {
		return db, nil
	}

	return nil, fmt.Errorf("Failed to get DB from context")
}

// generateDBContextKey db接続情報をcontextに格納する際のkeyを生成
func generateDBContextKey(name string) contextKey {
	return contextKey(fmt.Sprintf("DB_%s", name))
}

// MustGetDB 確実にDBコネクションを取得
func MustGetDB(ctx context.Context, name string) *gorp.DbMap {
	db, err := GetDB(ctx, name)
	if err != nil {
		panic(err)
	}
	return db
}
