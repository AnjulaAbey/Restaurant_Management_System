package main

import (
	"github.com/gin-gonic/gin"
	"golang-restaurant-project/database"
	"golang-restaurant-project/routes"
	"log"
)

func main() {
	port := "8000"
	// Create the database connection
	config := database.NewEnvDBConfig()
	db, err := database.ConnectToDB(*config)
	if err != nil {
		log.Fatal(err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.FoodRoutes(router, db)
	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
