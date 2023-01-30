package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"musica.com/app/albums"
	"musica.com/app/utils"
)

func main() {
	// initialize config
	err := utils.InitConfig()
	if err != nil {
		log.Fatal("Config Error", err)
	}
	// initialize db connection
	err = utils.InitDBSession()
	if err != nil {
		log.Fatal("DB Error", err)
	}

	// initialize gin router
	router := gin.Default()

	router.GET("/albums", albums.ListAlbums)
	router.GET("/albums/:id", albums.FindOne)
	router.POST("/albums", albums.CreateAlbum)

	router.Run(":" + utils.Config.Port)
}
