package context

import (
	gorp	"gopkg.in/gorp.v1"
	"fmt"
	"context"
)

// SetDB sets db connection to context.
func SetDB(ctx context.Context, name string, db *gorp.DbMap) context.Context {
	return context.WithValue(ctx, fmt.Sprintf("DB%s", name), db)
}

// GetDB gets db connection from context.
func GetDB(ctx context.Context, name string) (*gorp.DbMap, error) {
	db, ok := ctx.Value(fmt.Sprintf("DB%s", name)).(*gorp.DbMap)
	if ok {
		return db, nil
	}

	return nil, fmt.Errorf("Failed to get DB from context")
}

// MustGetDB 確実にDBコネクションを取得
func MustGetDB(ctx context.Context, name string) *gorp.DbMap {
	db, err := GetDB(ctx, name)
	if err != nil {
		panic(err)
	}
	return db
}
