package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"musica.com/app/albums"
	"musica.com/app/utils"
)

func main() {

	// initialize config
	error := utils.InitConfig()
	if error != nil {
		log.Fatal(error)
	}

	// initialize db connection
	error = utils.InitDBSession()
	if error != nil {
		log.Fatal(error)
	}

	// initialize gin router
	router := gin.Default()

	router.GET("/albums", albums.ListAlbums)
	router.GET("/albums/:id", albums.FindOne)
	router.POST("/albums", albums.CreateAlbum)

	router.Run(":" + utils.Config.Port)
}
