package controller

import (
	"AccountSell/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Node node
type NodeFind struct {
	Data     string `form:"data" json:"data" xml:"data"  binding:"required"`
	SplitStr string `form:"splitstr" json:"splitstr" xml:"splitstr"  binding:"required"`
}

func FindAccount(c *gin.Context) {
	var form NodeFind
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	linSplit := "\r\n"
	if !strings.Contains(form.Data, "\r") {
		linSplit = "\n"
	}
	if !strings.Contains(form.Data, "\n") {
		linSplit = "\r"
	}

	itemSplit := makeSplitStr(form.SplitStr)

	data := strings.Split(form.Data, linSplit)

	var account []string
	hasPhone, index := isPhone(data[0], itemSplit)
	if hasPhone {
		for _, item := range data {
			if len(item) != 0 {
				itemS := strings.Split(item, itemSplit)
				account = append(account, itemS[index])
			}
		}
	} else {
		for _, item := range data {
			if len(item) != 0 {
				itemS := strings.Split(item, itemSplit)
				account = append(account, itemS[0])
			}
		}
	}
	var acc database.Accounts
	var accList []database.Accounts = make([]database.Accounts, 0)

	for _, item := range account {
		if hasPhone {
			data, err := acc.FindOnePhone(item)
			if err == nil {
				accList = append(accList, data)
			}

		} else {
			data, err := acc.FindOneAccount(item)
			if err == nil {
				accList = append(accList, data)
			}

		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "查询成功",
		"data":    accList,
	})
}

func isPhone(s, itemSplit string) (has bool, i int) {
	var x bool = false
	var index int = 0
	itemS := strings.Split(s, itemSplit)
	for i, item := range itemS {
		if len(item) == 11 {
			if isDigit(item) {
				x = true
				index = i
				break
			}
		}
	}
	return x, index
}
