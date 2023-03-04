package controller

import (
	"AccountSell/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
}

func UserList(c *gin.Context) {
	UserList, err := database.GetUserList()
	if err != nil {
		data := gin.H{
			"status":  1,
			"message": err.Error(),
		}
		c.JSON(http.StatusOK, data)
		return
	}
	data := gin.H{
		"status":   0,
		"userList": UserList,
	}
	c.JSON(http.StatusOK, data)
}

func AddUser(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	user, err := database.CheckUserName(form.UserName)
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(user.UserName) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "用户名已存在",
		})
		return
	}
	users := &database.User{
		UserName: form.UserName,
	}
	err = users.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "添加成功",
		"data":    users,
	})
}
