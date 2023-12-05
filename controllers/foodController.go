package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang-restaurant-project/models"
	"net/http"
	"strconv"
)

func GetFood(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId, err := strconv.Atoi(c.Param("food_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}
		var food models.Food
		err = db.QueryRow("SELECT * FROM Food WHERE food_id=?", foodId).Scan(&food.FoodId, &food.Name, &food.FoodType, &food.Price)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving food item"})
			}
			return
		}
		c.JSON(http.StatusOK, food)
	}
}
func GetFoods(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query all food items from the database
		foods, err := db.Query("SELECT * FROM Food")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving food items"})
			return
		}
		defer foods.Close()

		var foodList []models.Food
		for foods.Next() {
			var food models.Food
			if err := foods.Scan(&food.FoodId, &food.Name, &food.FoodType, &food.Price); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning food items"})
				return
			}
			foodList = append(foodList, food)
		}
		c.JSON(http.StatusOK, foodList)
	}
}

func DeleteFood(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId, err := strconv.Atoi(c.Param("food_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}
		// Check if the food item exists before deleting
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM Food WHERE food_id=?", foodId).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for food item existence"})
			return
		}

		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}
		_, err = db.Exec("DELETE FROM Food WHERE food_Id=?", foodId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting food item"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": "Item deleted successfully"})
	}
}

func UpdateFoodPrice(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId, err := strconv.Atoi(c.Param("food_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}

		// Check if the food item exists before updating
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM Food WHERE food_id=?", foodId).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for food item existence"})
			return
		}

		if count == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}

		// request body
		var requestBody struct {
			Price string `json:"price"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Update the price of the food item
		_, err = db.Exec("UPDATE Food SET price=? WHERE food_id=?", requestBody.Price, foodId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating food item price"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Food item price updated"})
	}
}

func CreateFood(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			Name     string `json:"name"`
			FoodType string `json:"type"`
			Price    string `json:"price"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
		// Need to know the number of rows in the database
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM Food").Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for food item existence"})
			return
		}
		// Insert the food item into the database (Since id is incrementing, make the id equals to the count)
		_, err = db.Exec("INSERT INTO Food (food_id,name, food_type, price) VALUES (?,?, ?, ?)", count, requestBody.Name, requestBody.FoodType, requestBody.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting food item"})
			return
		}
		c.JSON(http.StatusCreated, nil)
	}
}
