package main

import (
	"ads/CSV"
	"ads/controllers"
	"ads/database"
	"ads/midlleware"

	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDatabase()
}
func main() {
	r := gin.Default()

	rows := CSV.ReadFile("role.csv")
	rows2 := CSV.ReadFile("account.csv")
	//
	CSV.InsertDataCSV(rows)
	CSV.InsertDataAccount(rows2)

	api := r.Group("api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginAccount)
		//api.POST("/", controllers.AddProduct)
		//api.PATCH("/:id", controllers.EditProduct)
		//api.POST("/login", controllers.Login)
		//api.POST("/create", controllers.Register)
		api.GET("/info", midlleware.RequireAuthen, controllers.GetInforUser)
		api.GET("/list-account", midlleware.RequireAuthen, controllers.ListUser)
		api.GET("/account:id", midlleware.RequireAuthen, controllers.DetailUser)
		api.DELETE("/account:id", midlleware.RequireAuthen, controllers.DeleteUser)

	}

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
