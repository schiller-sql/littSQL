package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
	LoginHandler       func(*gin.Context)
	LogoutHandler      func(*gin.Context)
	RefreshHandler     func(*gin.Context)
	JwtHandler         func(*gin.Context)
	IsTeacherValidator func(*gin.Context)
	IsStudentValidator func(*gin.Context)
}
