package users

import (
	"bookstore-users-api/datasources/mysql/users_db"
	"bookstore-users-api/utils/date_utils"
	"bookstore-users-api/utils/errors"
	"bookstore-users-api/utils/mysql_utils"
	"fmt"
)

const (
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created,status,password) VALUES(?, ?, ?, ?,?,?);"
	queryGetUser = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status=?;"

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
	if getErr := result.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
		//if strings.Contains(getErr.Error(),errorNoRows) {
		//	return errors.NewNotFoundError(
		//		fmt.Sprintf("user %d not found",user.Id))
		//}
		//sqlErr,ok := getErr.(*mysql.MySQLError)
		//if !ok {
		//	return errors.NewInternalServerError(
		//		fmt.Sprintf("error when trying to get user: %s",getErr.Error()))
		//}
		//fmt.Println(sqlErr)
		//
		//fmt.Println(sqlErr)
		//return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s",user.Id,getErr.Error()))
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

	insertResult,saveErr :=  stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated,user.Status,user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
		//sqlErr,ok := saveErr.(*mysql.MySQLError)
		//if !ok {
		//	return errors.NewInternalServerError(
		//		fmt.Sprintf("error when trying to save user: %s",saveErr.Error()))
		//}
		//fmt.Println(sqlErr.Number)
		//fmt.Println(sqlErr.Message)
		//switch sqlErr.Number{
		//case 1062:
		//	return errors.NewBadRequestError(
		//			fmt.Sprintf("emails %s already exists",user.Email))
		//}
		//
		//return errors.NewInternalServerError(
		//	fmt.Sprintf("error when trying to save user: %s",err.Error()))
	}
	//result,err := users_db.Client.Exec(queryInsertUser,user.FirstName,user.LastName,user.Email,user.DateCreated)
	userId, err := insertResult.LastInsertId()
	if err != nil {
		mysql_utils.ParseError(saveErr)
		//return errors.NewInternalServerError(
		//	fmt.Sprintf("error when trying to save user: %s",err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt,err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_,err = stmt.Exec(user.FirstName,user.LastName,user.Email,user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete()*errors.RestErr {
	stmt,err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _,err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User,*errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil,errors.NewInternalServerError(err.Error())

	}
		rows,err := stmt.Query(status)
	defer rows.Close()
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
	defer stmt.Close()

		results := make([]User,0)
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); err != nil {
				return nil,mysql_utils.ParseError(err)
			}
			results = append(results,user)
		}
		if len(results) == 0 {
			return nil,errors.NewNotFoundError(fmt.Sprintf("no users matching status %s",status))
		}
		return results,nil
}