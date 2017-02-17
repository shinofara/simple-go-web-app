package entity

type User struct {
	ID int `db:"id, primarykey, autoincrement"`
	Name string `db:"name"`
}
