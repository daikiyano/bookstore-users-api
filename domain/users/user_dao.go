package users

import (
	"bookstore-users-api/datasources/mysql/users_db"
	"bookstore-users-api/utils/date_utils"
	"bookstore-users-api/utils/errors"
	"fmt"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?"

)

var (
	usersDB = make(map[int64]*User)
)
func something() {
	user := User{}

	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)


}

func (user *User) Get() *errors.RestErr {
	stmt,err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	//必ずcloseする
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	fmt.Println("QueryRow",result)
	if err := result.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated); err != nil {
		if strings.Contains(err.Error(),errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found",user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s",user.Id,err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	println("Resecdc",user)
	stmt,err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	//必ずcloseする
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()

	insertResult,err :=  stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(),indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("emails %s already exists",user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s",err.Error()))
	}
	//result,err := users_db.Client.Exec(queryInsertUser,user.FirstName,user.LastName,user.Email,user.DateCreated)
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s",err.Error()))
	}
	user.Id = userId
	return nil
	//
	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("email %s already registerd",user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("user already exists",user.Id))
	//}
	//user.DateCreated = date_utils.GetNowString()
	//
	//usersDB[user.Id] = user
	//return nil
}