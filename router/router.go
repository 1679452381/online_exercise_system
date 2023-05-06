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
	//用户提交排行
	r.GET("/rank_list", service.GetRankList)
	//用户提交列表
	r.GET("/submit/list", service.SubmitList)

	problem := r.Group("/problem")
	//问题列表
	problem.GET("/list", service.ProblemList)
	//问题详情
	problem.GET("/detail", service.ProblemDetail)

	//用户组
	auth := r.Group("/u", middleware.AuthCheck())
	auth.GET("/detail", service.UserDetail)
	//管理员创建问题
	auth.POST("/problem/add", middleware.AuthAdminCheck(), service.CreateProblem)
	auth.POST("/test", service.Hello)

	return r
}
