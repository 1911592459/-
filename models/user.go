package models

type User struct {
	Id   int `json:"userId" binding:"required"` //用户id
	Step int `json:"step" binding:"required"`   //用户骰子数

}

func NewUser() *User {
	return &User{
		0,
		0,
	}
}
