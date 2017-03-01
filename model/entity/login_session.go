package entity

import "time"

// LoginSession ログインに必要なセッション情報を保持
type LoginSession struct {
	LastLoginDate time.Time
}
