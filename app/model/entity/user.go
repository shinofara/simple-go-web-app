// Package entity データストアに関するEntityを定義
package entity

import "encoding/base64"

// User ユーザ情報を保持
type User struct {
	ID int `db:"id, primarykey, autoincrement"`
	Name string `db:"name"`
	Pass string `db:"-"`
	EncryptPass string `db:"pass"`
}

// NewUser creates a User
func NewUser(name, pass string) *User {
	return &User{
		Name: name,
		Pass: pass,
		EncryptPass: encrypt(pass),
	}
}

// Encrypt 元となる文字列を暗号化する
func encrypt(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
