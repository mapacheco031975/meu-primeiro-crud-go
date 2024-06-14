package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mapacheco031975/meu-primeiro-cru-go/src/controller/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro")
	}
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	router.Run(":8082")

	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
