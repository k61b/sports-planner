package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kayraberktuncer/sports-planner/pkg/common/db"
	"github.com/kayraberktuncer/sports-planner/pkg/exercises"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	router := gin.Default()
	dbHandler := db.Init()

	exercises.RegisterRoutes(router, dbHandler)

	router.Run(port)
}

//https://medium.com/swlh/create-go-service-the-easy-way-de827d7f07cf
