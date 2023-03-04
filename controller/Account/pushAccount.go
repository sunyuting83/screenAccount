package controller

import (
	"AccountSell/database"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// Node node
type Node struct {
	UserID string `form:"userid" json:"userid" xml:"userid"  binding:"required"`
}

func PushAccount(c *gin.Context) {
	var form Node
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
	file, _, err := c.Request.FormFile("files")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "1上传文件失败",
		})
		return
	}

	b, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "1上传文件失败",
		})
		return
	}

	UserID, _ := strconv.Atoi(form.UserID)

	linSplit := "\r\n"
	if !strings.Contains(string(b), "\r") {
		linSplit = "\n"
	}
	if !strings.Contains(string(b), "\n") {
		linSplit = "\r"
	}

	itemSplit := "\t"

	data := strings.Split(string(b), linSplit)

	var account []database.Accounts

	for _, item := range data {
		if len(item) > 0 {
			itemS := strings.Split(item, itemSplit)
			PriceLen := len(itemS) - 1
			Price, _ := strconv.Atoi(itemS[PriceLen])

			if isDigit(itemS[1]) && len(itemS[1]) == 11 {
				Gold, _ := strconv.Atoi(itemS[3])
				account = append(account, database.Accounts{
					UserID:   uint(UserID),
					Account:  itemS[0],
					Phone:    itemS[1],
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
		"message": "上传文件成功",
	})
}

func isDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}
