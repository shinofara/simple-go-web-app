package context

import (
	gorp	"gopkg.in/gorp.v1"
	"fmt"
	"context"
)

func SetDB(ctx context.Context, name string, db *gorp.DbMap) context.Context {
	return context.WithValue(ctx, fmt.Sprintf("DB%s", name), db)
}

func GetDB(ctx context.Context, name string) (*gorp.DbMap, error) {
	db, ok := ctx.Value(fmt.Sprintf("DB%s", name)).(*gorp.DbMap)
	if ok {
		return db, nil
	}

	return nil, fmt.Errorf("Failed to get DB from context")
}

func MustGetDB(ctx context.Context, name string) *gorp.DbMap {
	db, err := GetDB(ctx, name)
	if err != nil {
		panic(err)
	}
	return db
}
