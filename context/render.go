package context

import (
	"fmt"
	"context"
	"github.com/shinofara/simple-go-web-app/render"
)

func SetRender(ctx context.Context, re *render.Render) context.Context {
	return context.WithValue(ctx, "RENDER", re)
}

func GetRender(ctx context.Context) (*render.Render, error) {
	l, ok := ctx.Value("RENDER").(*render.Render)
	if ok {
		return l, nil
	}

	return nil, fmt.Errorf("Failed to get RENDER from context")
}

func MustGetRender(ctx context.Context) *render.Render {
	l, err := GetRender(ctx)
	if err != nil {
		panic(err)
	}
	return l
}
