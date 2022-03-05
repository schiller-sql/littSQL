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

	auth.GET("/account", jwtMiddleware.MiddlewareFunc(), handler.getAccountDetails)
	auth.DELETE("/account", jwtMiddleware.MiddlewareFunc(), handler.deleteAccount)

}

type teacherSignUp struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *authHandler) signup(c *gin.Context) {
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
}

func (h *authHandler) deleteAccount(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	if claims["is_teacher"].(bool) {
		id, _ := c.Get("id")
		err := h.usecase.DeleteTeacher(id.(int32))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ONLY A TEACHER CAN DELETE THEIR ACCOUNT"})
	}
}

func (h *authHandler) getAccountDetails(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id, _ := c.Get("id")
	if claims["is_teacher"].(bool) {
		teacher, err := h.usecase.GetTeacherAccountDetails(id.(int32))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"id": teacher.ID, "email": teacher.Email})
	} else {
		participant, err := h.usecase.GetParticipantAccountDetails(id.(int32))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := gin.H{"id": participant.ID, "name": nil}
		if participant.Name.Valid {
			res["name"] = participant.Name.String
		}
		c.JSON(200, res)
	}
}
