package dto

type ChangeUserPassword struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
