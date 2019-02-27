package db

import (
	"fmt"
	"github.com/edumar111/fastpv-auth/user/model"

	tools "github.com/edumar111/fastpv-auth/tools"
	"log"
)

//InsertUser insert User to DB
func InsertUser(user model.User) int64{
	fmt.Println("InsertUser")
	db := tools.DBConn()


	var stmt, err = db.Prepare("INSERT INTO tb_user(user_name, password," +
		" email,created_user,created_date,updated_user,updated_date, origin, status,image," +
		"fisrt_name,last_name) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	fmt.Println("stmt")
	if err != nil {
		log.Println(err.Error())
	}

	res, err :=stmt.Exec(user.Username,
		user.Password, user.Email,
		user.CreatedUser,user.CreatedDate,
		user.UpdatedUser, user.UpdatedDate,user.Origin, user.Status, user.Image, user.FisrtName, user.LastName)


	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(id)
	defer db.Close()
	return id
}

//FindOneUser get user for (iduser)
func FindOneUser(iduser int64)model.User {
	db := tools.DBConn()

	selDB, err := db.Query("SELECT id_user,user_name,email, origin,status, image,fisrt_name, last_name," +
		"updated_user,updated_date  FROM tb_user WHERE id_user=?", iduser)
	if err != nil {
		log.Println(err.Error())
	}
	user := model.User{}
	for selDB.Next() {

		err = selDB.Scan(&user.ID,&user.Username, &user.Email, &user.Origin, &user.Status,
			&user.Image, &user.FisrtName,&user.LastName, &user.UpdatedUser, &user.UpdatedDate)
		if err != nil {
			log.Println(err.Error())
		}

	}
	fmt.Println(user)
	defer db.Close()
	return  user
}

//FindUsername
func  FindUsername(username string  ) model.User{
	db := tools.DBConn()
	selDB, err := db.Query("SELECT id_user,user_name,password,email, origin,status, image, fisrt_name, last_name FROM tb_user WHERE user_name=?", username)
	if err != nil {
		log.Println(err.Error())
	}
	user := model.User{}
	for selDB.Next() {

		err = selDB.Scan(&user.ID,&user.Username,&user.Password, &user.Email, &user.Origin, &user.Status,
			&user.Image, &user.FisrtName,&user.LastName)
		if err != nil {
			log.Println(err.Error())
		}

	}
	defer db.Close()
	return  user
}

//UpdateUser update user
func UpdateUser(user model.User){
	db := tools.DBConn()


	insForm, err := db.Prepare("UPDATE tb_user SET user_name=?, email=?, fisrt_name=?,last_name," +
		"updated_user, updated_date WHERE id_user=?")
	if err != nil {
		log.Println(err.Error())
	}
	insForm.Exec(user.Username, user.Email, user.FisrtName, user.LastName, user.UpdatedUser, user.UpdatedDate)
	log.Println("UPDATE: Name: " + user.Username + " | Email: " + user.Email)

	defer db.Close()

}
