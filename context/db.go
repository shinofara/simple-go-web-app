package context

import (
	gorp	"gopkg.in/gorp.v1"
	"fmt"
	"context"
)

func SetDB(ctx context.Context, db *gorp.DbMap) context.Context {
	return context.WithValue(ctx, "DB", db)
}

func GetDB(ctx context.Context) (*gorp.DbMap, error) {
	db, ok := ctx.Value("DB").(*gorp.DbMap)
	if ok {
		return db, nil
	}

	return nil, fmt.Errorf("Failed to get DB from context")
}

func MustGetDB(ctx context.Context) *gorp.DbMap {
	db, err := GetDB(ctx)
	if err != nil {
		panic(err)
	}
	return db
}
