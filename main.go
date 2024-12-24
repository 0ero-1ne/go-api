package main

import (
	"fmt"
	"simpapi/db"
	"simpapi/envresolver"
	"simpapi/models"
)

func main() {
	dns := envresolver.GetKey("dns")

	db := db.OpenDB(dns)
	connection, err := db.DB()

	if err != nil {
		panic("Get DB error")
	}

	defer connection.Close()

	db.AutoMigrate(&models.Country{}, &models.City{}, &models.User{})

	var resultCountries []models.Country

	db.Find(&resultCountries)

	for k, v := range resultCountries {
		fmt.Println("Row = ", k)
		fmt.Println("ID =", v.ID)
		fmt.Println("Title =", v.Title)
		fmt.Println("CreatedAt =", v.CreatedAt)
		fmt.Println("UpdatedAt =", v.UpdatedAt)
		fmt.Println("-----")
	}
}
