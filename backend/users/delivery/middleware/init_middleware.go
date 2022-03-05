package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/schiller-sql/littSQL/model"
	"github.com/schiller-sql/littSQL/users"
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

func NewUsersMiddleware(authusecase users.Usecase) *jwt.GinJWTMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "littSQL",
		SigningAlgorithm: viper.Get("JWT_SIGN_ALG").(string),
		IdentityKey:      "id",
		Key:              []byte(viper.Get("JWT_SECRET").(string)),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24 * 7,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req teacherLogin
			err := c.ShouldBindJSON(&req)
			if err != nil {
				var req participantLogin
				err := c.ShouldBindJSON(&req)
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
				return jwt.MapClaims{"is_teacher": false, "id": user.(*model.Participant).ID}
			case *model.Teacher:
				return jwt.MapClaims{"is_teacher": true, "id": user.(*model.Teacher).ID}
			default:
				panic("Authenticator should not give through a non Teacher and Participant")
			}
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
