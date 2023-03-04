package controller

import (
	"AccountSell/database"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Node node
type NodeList struct {
	UserID   string `form:"userid" json:"userid" xml:"userid"  binding:"required"`
	Data     string `form:"data" json:"data" xml:"data"  binding:"required"`
	SplitStr string `form:"splitstr" json:"splitstr" xml:"splitstr"  binding:"required"`
}

func PostAccount(c *gin.Context) {
	var form NodeList
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(form.UserID) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't node",
		})
		return
	}

	UserID, _ := strconv.Atoi(form.UserID)

	linSplit := "\r\n"
	if !strings.Contains(form.Data, "\r") {
		linSplit = "\n"
	}
	if !strings.Contains(form.Data, "\n") {
		linSplit = "\r"
	}

	itemSplit := makeSplitStr(form.SplitStr)

	data := strings.Split(form.Data, linSplit)
	hasPhone, index := isPhone(data[0], itemSplit)

	var account []database.Accounts

	for _, item := range data {
		itemS := strings.Split(item, itemSplit)
		if len(item) != 0 {
			PriceLen := len(itemS) - 1
			Price, _ := strconv.Atoi(itemS[PriceLen])
			if hasPhone {
				Gold, _ := strconv.Atoi(itemS[3])
				account = append(account, database.Accounts{
					UserID:   uint(UserID),
					Account:  itemS[0],
					Phone:    itemS[index],
					Password: itemS[2],
					Gold:     Gold,
					Price:    Price,
				})
			} else {
				Gold, _ := strconv.Atoi(itemS[2])
				account = append(account, database.Accounts{
					UserID:  uint(UserID),
					Account: itemS[0],
					Gold:    Gold,
					Price:   Price,
				})
			}
		}
	}
	database.AccountBatches(account)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "导入成功",
	})
}

func makeSplitStr(s string) string {
	var x string = "\t"
	switch s {
	case "0":
		x = "\t"
	case "1":
		x = "----"
	case "2":
		x = " "
	default:
		x = "\t"
	}
	return x
}
