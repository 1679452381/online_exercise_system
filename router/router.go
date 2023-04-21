package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"online_exercise_system/middleware"
	"online_exercise_system/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/hello", service.Hello)
	r.POST("/login", service.Login)
	r.POST("/email/code", service.SendEmailCode)
	r.POST("/register", service.Register)
	problem := r.Group("/problem")
	problem.GET("/list", service.ProblemList)
	//用户组
	auth := r.Group("/u", middleware.AuthCheck())
	auth.POST("/test", service.Hello)

	return r
}
