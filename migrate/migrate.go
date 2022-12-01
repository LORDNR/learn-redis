package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lordnr/learn-redis/config"
	"github.com/lordnr/learn-redis/database"
	"github.com/lordnr/learn-redis/models"
)

func init() {
	config.LoadEnvironment()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.Product{})
	mockData()
}

func mockData() error {
	var count int64
	database.DB.Model(&models.Product{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	var products []models.Product
	for i := 0; i < 300; i++ {
		products = append(products, models.Product{
			Name:     fmt.Sprintf("Product%v", i+1),
			Quantity: uint16(random.Intn(100)),
		})
	}
	// result := database.DB.Create(&products)
	// if result.Error != nil {
	// 	fmt.Println("err")
	// } else {
	// 	fmt.Println("create success")
	// }

	return database.DB.Create(&products).Error
}
