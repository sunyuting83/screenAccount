package controller

import (
	"AccountSell/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AccountList(c *gin.Context) {
	var page string = c.DefaultQuery("page", "0")
	var UserID string = c.DefaultQuery("UserID", "0")
	var Limit string = c.DefaultQuery("limit", "10")
	pageInt, _ := strconv.Atoi(page)
	LimitInt, _ := strconv.Atoi(Limit)
	var account *database.Accounts
	count, err := account.GetCount(UserID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	dataList, err := database.GetAccountList(pageInt, UserID, LimitInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	Data := gin.H{
		"status": 0,
		"data":   dataList,
		"total":  count,
	}
	c.JSON(http.StatusOK, Data)
}
