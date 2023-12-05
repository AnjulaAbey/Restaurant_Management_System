package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang-restaurant-project/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine, db *sql.DB) {
	incomingRoutes.GET("/foods/getFoods", controllers.GetFoods(db))
	incomingRoutes.GET("/foods/getFood/:food_id", controllers.GetFood(db))
	incomingRoutes.DELETE("/foods/deleteFood/:food_id", controllers.DeleteFood(db))
	incomingRoutes.PUT("/foods/updateFood/:food_id", controllers.UpdateFoodPrice(db))
	incomingRoutes.POST("/foods/createFood", controllers.CreateFood(db))
}
