package main

import (
	"fmt"

	"github.com/azacdev/go-react/api/controllers"
	"github.com/azacdev/go-react/config"
	"github.com/azacdev/go-react/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDB()
	database.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(cors.Default())
	fmt.Println("Wubba lubba dub dub")

	r.GET("/entries", controllers.GetEntreies)
	r.GET("/entry/:id", controllers.GetEntry)
	r.GET("/ingredient/:ingredient", controllers.GetEntryByIngredient)

	r.POST("/entry/create", controllers.CreateEntry)

	r.PUT("/entry/update/:id", controllers.UpdateEntry)
	r.PUT("/ingredient/update/:id", controllers.UpdateIngredient)
	r.DELETE("/entry/delete/:id", controllers.DeleteEntry)

	r.Run()
}
