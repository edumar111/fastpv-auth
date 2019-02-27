package service

import (
	"fmt"
	"github.com/edumar111/fastpv-auth/tools"
	"github.com/edumar111/fastpv-auth/user/db"
	"github.com/edumar111/fastpv-auth/user/model"
	"time"
)

// CreateUser crear usuario
func CreateUser(user model.User ) (model.User, error){

	var datetime = time.Now().Local()
	time.LoadLocation("America/Lima")
	datetime.Format(time.RFC3339)
	password := user.Password

	hashPassword ,err := tools.HashPassword(password)
	if err != nil{

		return user,err
	}

	user.Password = hashPassword
	user.CreatedDate = datetime.String()
	user.UpdatedDate = datetime.String()

	fmt.Println(datetime)

	fmt.Println("CreateUser", user)

	id :=db.InsertUser (user)
	user.ID=id
	user.Password=""
	return user,nil
}
//GetUser
func GetUser(iduser int64) model.User {

	return db.FindOneUser(iduser)

}

func Authenticate(username string, password string  ) bool {
	user :=db.FindUsername(username)
	return user.Username == username && tools.CheckHashAndPassword(user.Password,password)
}