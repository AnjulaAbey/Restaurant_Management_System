package models

// Food Structure should have an ID, Name, Type and Price
type Food struct {
	FoodId   string `json:"ID"`
	Name     string `json:"Name"`
	FoodType string `json:"Type"`
	Price    string `json:"Price"`
}
