package model
type UserLogin struct {
	ID     int64 `json:"id_user" form:"-"`
	Username string `json:"username" form:"user_name"`
	Password string `json:"password" form:"password"`
}