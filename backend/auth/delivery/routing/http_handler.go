package routing

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/schiller-sql/littSQL/auth"
)

type authHandler struct {
	usecase auth.Usecase
}

func ConfigureHandler(r *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware, usecase auth.Usecase) {
	handler := authHandler{usecase}

	group := r.Group("/auth")

	group.POST("/signup", handler.signup)
	group.POST("/login", jwtMiddleware.LoginHandler)
	group.POST("/logout", jwtMiddleware.LogoutHandler)
	group.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	group.GET("/account", jwtMiddleware.MiddlewareFunc(), handler.getAccountDetails)
	group.DELETE("/account", jwtMiddleware.MiddlewareFunc(), handler.deleteAccount)

}

type teacherSignUp struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *authHandler) signup(c *gin.Context) {
	var req teacherSignUp
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MAKE SURE THE EMAIL IS VALID"})
		return
	}
	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PASSWORD SHOULD BE AT LEAST SIX CHARACTERS LONG"})
		return
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
		if participant.Name != nil {
			res["name"] = *participant.Name
		}
		c.JSON(200, res)
	}
}
