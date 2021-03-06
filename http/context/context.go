// Package context リクエストを受けてからレスポンスを返すまでに一貫して維持したい情報を格納
package context

import "context"

// contextKey key
type contextKey string

// Context wraps original context.
type Context context.Context
