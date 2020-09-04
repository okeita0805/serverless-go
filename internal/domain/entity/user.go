package entity

// User はユーザ
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Users はUserのスライス
type Users []User
