package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/schiller-sql/littSQL/auth"
	"github.com/schiller-sql/littSQL/model"
	"github.com/spf13/viper"
)

type participantLogin struct {
	AccessCode string `json:"access_code" binding:"required"`
}

type teacherLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewAuthMiddleware(usecase auth.Usecase) *AuthMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "littSQL",
		SigningAlgorithm: viper.Get("JWT_SIGN_ALG").(string),
		IdentityKey:      "id",
		Key:              []byte(viper.Get("JWT_SECRET").(string)),
		Timeout:          time.Hour * 24,     // TODO: Add config in viper
		MaxRefresh:       time.Hour * 24 * 7, // TODO: Add config in viper
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var jsonData gin.H
			data, _ := ioutil.ReadAll(c.Request.Body)
			c.Set(gin.BodyBytesKey, data)
			if err := json.Unmarshal(data, &jsonData); err != nil {
				return nil, err
			}
			_, accessCodeExists := jsonData["access_code"]
			_, emailExists := jsonData["email"]
			_, passwordExists := jsonData["password"]
			loginFormExists := emailExists && passwordExists
			if !accessCodeExists && !loginFormExists {
				return nil, fmt.Errorf("need property 'access_code' to login as a student " +
					"or properties 'email' and 'password' to login as a teacher")
			}
			if accessCodeExists && (emailExists || passwordExists) {
				return nil, fmt.Errorf("cannot login as a student and teacher at the same time," +
					"use property 'access_code' to login as a student " +
					"or properties 'email' and 'password' to login as a teacher")
			}
			if accessCodeExists {
				var req participantLogin
				if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
					return nil, fmt.Errorf("make sure the access code is a string")
				}
				if len(req.AccessCode) != 6 {
					return nil, fmt.Errorf("the length of the access code has to be 6")
				}
				participant, err := usecase.LogInParticipant(req.AccessCode)
				if err != nil {
					return nil, err
				}
				return participant, nil
			} else {
				var req teacherLogin
				err := c.ShouldBindBodyWith(&req, binding.JSON)
				if err != nil {
					return nil, fmt.Errorf("make sure the email is a valid email " +
						"and the password is at least six characters long")
				}
				teacher, err := usecase.LogInTeacher(req.Email, req.Password)
				if err != nil {
					return nil, err
				}
				return teacher, nil
			}
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
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
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
	return &AuthMiddleware{
		LoginHandler:   jwtMiddleware.LoginHandler,
		LogoutHandler:  jwtMiddleware.LogoutHandler,
		RefreshHandler: jwtMiddleware.RefreshHandler,
		JwtHandler:     jwtMiddleware.MiddlewareFunc(),
		IsTeacherValidator: func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			if !claims["is_teacher"].(bool) {
				c.JSON(http.StatusForbidden, gin.H{"error": "you have to be a teacher to access this resource"})
			}
		},
		IsStudentValidator: func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			if claims["is_teacher"].(bool) {
				c.JSON(http.StatusForbidden, gin.H{"error": "you have to be a student to access this resource"})
			}
		},
	}
}
