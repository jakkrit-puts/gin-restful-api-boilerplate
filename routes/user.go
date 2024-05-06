package routes

import (
	userctrl "example.com/jakkrit/ginbackendapi/controllers/user"
	"example.com/jakkrit/ginbackendapi/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")
	// routerGroup := rg.Group("/users").Use(middlewares.AuthJWT())

	routerGroup.POST("/register", userctrl.Register)
	routerGroup.POST("/login", userctrl.Login)
	routerGroup.GET("/", middlewares.AuthJWT(), userctrl.GetAll)
	routerGroup.GET("/:id", middlewares.AuthJWT(), userctrl.GetById)
	routerGroup.GET("/search", middlewares.AuthJWT(), userctrl.SearchByName)
	routerGroup.GET("/me", middlewares.AuthJWT(), userctrl.GetProfile)

}
