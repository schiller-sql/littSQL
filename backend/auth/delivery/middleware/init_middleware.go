package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/schiller-sql/littSQL/auth"
	"github.com/schiller-sql/littSQL/model"
	"github.com/spf13/viper"
	"time"
)

type teacherLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type participantLogin struct {
	AccessCode string `json:"access_code" binding:"required"`
}

func NewAuthMiddleware(authusecase auth.Usecase) *jwt.GinJWTMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "littSQL",
		SigningAlgorithm: viper.Get("JWT_SIGN_ALG").(string),
		IdentityKey:      "id",
		Key:              []byte(viper.Get("JWT_SECRET").(string)),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24 * 7,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req teacherLogin
			err := c.ShouldBindBodyWith(&req, binding.JSON)
			if err != nil {
				var req participantLogin
				err := c.ShouldBindBodyWith(&req, binding.JSON)
				if err != nil {
					return nil, err
				}
				participant, err := authusecase.LogInParticipant(req.AccessCode)
				if err != nil {
					return nil, err
				}
				return participant, nil
			}
			teacher, err := authusecase.LogInTeacher(req.Email, req.Password)
			if err != nil {
				return nil, err
			}
			return teacher, nil
		},
		PayloadFunc: func(user interface{}) jwt.MapClaims {
			switch user.(type) {
			case *model.Participant:
				return jwt.MapClaims{"is_teacher": false, "id": int(user.(*model.Participant).ID)}
			case *model.Teacher:
				return jwt.MapClaims{"is_teacher": true, "id": int(user.(*model.Teacher).ID)}
			default:
				panic("Authenticator should not give through a non Teacher and Participant")
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return int32(claims["id"].(float64))
		},
	})
	if err != nil {
		panic(err)
	}
	err = jwtMiddleware.MiddlewareInit()
	if err != nil {
		panic(err)
	}
	return jwtMiddleware
}
