package users

import (
	"bookstore-users-api/domain/users"
	"bookstore-users-api/services"
	"bookstore-users-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userIdParams string)(int64,*errors.RestErr) {
	userId,userErr := strconv.ParseInt(userIdParams,10,64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id should be a number")
	}
	return userId,nil
}

func Create(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	//{0    }
	//リクエストBodyの読み込み
	//json.UnmarshalでUser構造体に格納
	//指定されたバインディングエンジンを使用して、渡された構造体ポインターをバインドする。
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("heyyyyyyyyyyyyy")
		fmt.Println(err)
		//無効なJSONの場合、エラー返し
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated,result.Marshall(c.GetHeader("X-Public") == "true"))
	}



func Get(c *gin.Context) {
	// int型に変換し、userIdを取得
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id should be anumber")
		c.JSON(err.Status,err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK,user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id should be anumber")
		c.JSON(err.Status,err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//無効なJSONの場合、エラー返し
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result,err := services.UpdateUser(isPartial,user)
	if err != nil {
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId,idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status,idErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,map[string]string{"status":"deleted"})


}

func Search(c *gin.Context) {
	status := c.Query("status")
	users,err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	result := make([]interface{,len(users)})
	for index,user := range users {
		result[index] = user.Marshall(c.GetHeader("X-Public") == "true")
	}
	c.JSON(http.StatusOK,result)
}


//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented,"implement me!")
//}
