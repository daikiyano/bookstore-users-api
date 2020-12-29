package users

import (
	"bookstore-users-api/domain/users"
	"bookstore-users-api/services"
	//"bookstore-users-api/domain/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	//{0    }
	//リクエストBodyの読み込み
	//json.UnmarshalでUser構造体に格納
	//指定されたバインディングエンジンを使用して、渡された構造体ポインターをバインドする。
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		// TODO: return bad request
	}
	fmt.Println(user)
	//{123 daiki   }
	//{"id":123,"first_name": "daiki"}
	result,saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented,"implement me!")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented,"implement me!")
//}
