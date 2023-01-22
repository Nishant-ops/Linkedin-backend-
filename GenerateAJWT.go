package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


var sampleSecretKey = []byte("uGwk9h/Sj29wYNoLKHF54elBgbim8Z87+51j0wklbkyy+dkLZhDiu7N+8K71Zquv")


func generateJWT(c *gin.Context){
token := jwt.New(jwt.SigningMethodHS256)
//c.JSON(http.StatusInternalServerError,gin.H{"message":token})
claims:=token.Claims.(jwt.MapClaims)
claims["exp"]=time.Now().Add(10*time.Minute)
claims["authorized"] = true
tokenString, err := token.SignedString(sampleSecretKey)
if err != nil {
	fmt.Print("ello->",err);
     c.JSON(http.StatusInternalServerError,gin.H{"message":err})
	 return
 }

c.JSON(http.StatusOK,gin.H{"token":tokenString})
}