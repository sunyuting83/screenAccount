package main

import (
	orm "AccountSell/database"
	"AccountSell/router"
	"AccountSell/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	orm.InitDB()
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)
	defer orm.Eloquent.Close()

	app := router.InitRouter()
	utils.OpenURI("http://localhost:13088/")

	app.Run(strings.Join([]string{":", "13088"}, ""))
}
