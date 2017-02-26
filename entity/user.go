// Package entity データストアに関するEntityを定義
package entity

// User ユーザ情報を保持
type User struct {
	ID int `db:"id, primarykey, autoincrement"`
	Name string `db:"name"`
}
