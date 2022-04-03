package routing

import (
	authM "github.com/schiller-sql/littSQL/auth/delivery/middleware"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/schiller-sql/littSQL/auth"
)

type authHandler struct {
	usecase auth.Usecase
}

func ConfigureHandler(r *gin.RouterGroup, authMiddleware *authM.AuthMiddleware, usecase auth.Usecase) {
	handler := authHandler{usecase}

	group := r.Group("/auth")

	group.POST("/signup", handler.signup)
	group.POST("/login", authMiddleware.LoginHandler)
	group.POST("/logout", authMiddleware.LogoutHandler)
	group.GET("/refresh_token", authMiddleware.RefreshHandler)

	group.GET("/account", authMiddleware.JwtHandler, handler.getAccountDetails)
	group.DELETE("/account", authMiddleware.JwtHandler, handler.deleteAccount)

}

type teacherSignUp struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *authHandler) signup(c *gin.Context) {
	var req teacherSignUp
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "make sure the email is valid"})
		return
	}
	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password should be at least six characters long"})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only a teacher can delete their account"})
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
