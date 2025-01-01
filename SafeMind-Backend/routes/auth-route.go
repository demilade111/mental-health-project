package routes

import (
	"mental-health-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.Login)
		auth.GET("/verify", controllers.VerifyEmail) 
		auth.POST("/reset-password", controllers.ResetPassword) 
		auth.PUT("/reset-password", controllers.RequestResetPassword) 
	}
}
