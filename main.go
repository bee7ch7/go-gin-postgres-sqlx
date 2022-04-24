package main

import (
	"log"

	"github.com/bee7ch7/go-gin-postgres-sqlx/controllers"
	"github.com/bee7ch7/go-gin-postgres-sqlx/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()
	router.POST("/account", controllers.CreateAccount)
	router.GET("/accounts", controllers.GetAccounts)
	router.GET("/account/:id", controllers.GetAccount)

	log.Fatal(router.Run("0.0.0.0:8080"))

}
