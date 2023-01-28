package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var sampleSecretKey = []byte("uGwk9h/Sj29wYNoLKHF54elBgbim8Z87+51j0wklbkyy+dkLZhDiu7N+8K71Zquv")

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email"`
	UserName string             `json:"username"`
	Password string             `json:"password"`
}

func generateJWT(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		fmt.Print("ello->", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Add("Access-Control-Max-Age", "1800")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "content-type")
	c.Writer.Header().Add("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, PATCH, OPTIONS")
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("auth")
		fmt.Print("camehre", auth)
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No jwt token in header"})
			return
		}
		token, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("DDu")
			}
			return "", errors.New("DDu")
		})

		if token == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No jwt token in header"})
			return
		}
	}
}

func loginAUser(c *gin.Context) {

}
