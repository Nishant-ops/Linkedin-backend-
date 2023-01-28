package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

const db = "userDB"
const colName = "users"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri:=os.Getenv("MONGODB")
	fmt.Print("NISNAT CAME ,", uri)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Nishant", err, os.Getenv("MONGODB"))
	}

	collection = (*mongo.Collection)(client.Database(db).Collection(colName))
	title:="null"
	collection.InsertOne(context.TODO(),title)
}
func main() {
	fmt.Println("Hello world")
	handleFunc()
}
func handleFunc() {
	r := gin.Default()
	public := r.Group("/create")
	public.POST("/join", generateJWT)
	public.POST("/login", loginAUser)
	r.GET("/signup", handlePage)
	r.Use(AuthMiddleWare())
	r.Run()
}
