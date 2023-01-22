package main

import (
	"fmt"
    "github.com/gin-gonic/gin"
	
)

func main() {
	fmt.Println("Hello world");
	handleFunc()
}
func handleFunc(){
	r:=gin.Default()
	r.GET("jwt",generateJWT)
	r.GET("/signup",handlePage)
	r.Run()
}