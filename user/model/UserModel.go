package model


//User to person
type User struct {
	ID          int64  `json:"id_user"`
	Username    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CreatedUser int64  `json:"created_user"`
	CreatedDate string `json:"created_date"`
	UpdatedUser int64  `json:"updated_user"`
	UpdatedDate string `json:"updated_date"`
	Origin      string `json:"origin"`
	Status      string `json:"status"`
	Image       string `json:"image"`
	FisrtName   string `json:"fisrt_name"`
	LastName    string `json:"last_name"`
}