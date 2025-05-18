package model

type User struct {
	RoleId   int    `json:"role_id"`
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}
