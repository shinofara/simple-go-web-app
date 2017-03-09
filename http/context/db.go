package context

import (
	gorp	"gopkg.in/gorp.v1"
	"fmt"
	"context"
)

const ctxKeyDB = contextKey("DB")

// SetDB sets db connection to context.
func SetDB(ctx context.Context, db *gorp.DbMap) context.Context {
	return context.WithValue(ctx, ctxKeyDB, db)
}

// GetDB gets db connection from context.
func GetDB(ctx context.Context) (*gorp.DbMap, error) {
	db, ok := ctx.Value(ctxKeyDB).(*gorp.DbMap)
	if ok {
		return db, nil
	}

	return nil, fmt.Errorf("Failed to get DB from context")
}

// MustGetDB 確実にDBコネクションを取得
func MustGetDB(ctx context.Context) *gorp.DbMap {
	db, err := GetDB(ctx)
	if err != nil {
		panic(err)
	}
	return db
}
