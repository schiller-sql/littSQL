package http_handler

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/schiller-sql/littSQL/auth"
)

type HttpHandler struct {
	usecase auth.Usecase
}

func NewHttpHandler(usecase auth.Usecase) HttpHandler {
	return HttpHandler{usecase: usecase}
}

type teacherSignUp struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *HttpHandler) Signup(c *gin.Context) {
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

func (h *HttpHandler) DeleteAccount(c *gin.Context) {
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

func (h *HttpHandler) GetAccountDetails(c *gin.Context) {
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
