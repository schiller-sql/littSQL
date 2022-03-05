package routing

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/users"
	"net/http"
)

type authHandler struct {
	usecase users.Usecase
}

func ConfigureHandler(r *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware, usecase users.Usecase) {
	handler := authHandler{usecase}

	auth := r.Group("/auth")

	auth.POST("/signup", handler.signup)
	auth.POST("/login", jwtMiddleware.LoginHandler)
	auth.POST("/logout", jwtMiddleware.LogoutHandler)

	auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	auth.DELETE("/account", jwtMiddleware.MiddlewareFunc(), handler.deleteAccount)

}

type teacherSignUp struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h authHandler) signup(c *gin.Context) {
	var req teacherSignUp
	err := c.BindJSON(&req)
	if err != nil {
		return
	}
	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PASSWORD SHOULD BE AT LEAST SIX CHARACTERS LONG"})
	}
	err = h.usecase.SignUpTeacher(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(200)
}

func (h authHandler) deleteAccount(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	if claims["is_teacher"].(bool) {
		err := h.usecase.DeleteTeacher(int32(claims["id"].(float64)))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "NEED TO BE A TEACHER"})
	}
}
