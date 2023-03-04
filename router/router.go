package router

import (
	"AccountSell/controller"
	Account "AccountSell/controller/Account"
	User "AccountSell/controller/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// InitRouter make router
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.StaticFS("/css", http.Dir("static/css"))
	router.StaticFS("/js", http.Dir("static/js"))
	router.StaticFile("/favicon.ico", "static/favicon.ico")
	router.LoadHTMLGlob("static/html/*")
	api := router.Group("/api/v1")
	{
		router.GET("/", controller.Index)
		router.GET("/postdata", controller.Index)
		router.GET("/userdata", controller.Index)
		api.GET("/AccountList", Account.AccountList)
		api.GET("/UserList", User.UserList)
		api.POST("/AddUser", User.AddUser)
		api.POST("/PostAccount", Account.PostAccount)
		api.POST("/PushAccount", Account.PushAccount)
		api.POST("/FindAccount", Account.FindAccount)
	}

	return router
}
