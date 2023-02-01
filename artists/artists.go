package artists

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"musica.com/app/utils"
)

type Artist struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	CountryCode  string    `json:"country_code"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type createArtistRequest struct {
	Name         string `json:"name"`
	CountryCode  string `json:"country_code"`
	ProfileImage string `json:"profile_image"`
}

func FindOne(c *gin.Context) {
	var artist Artist
	artistId := c.Param("id")

	err := utils.DB.First(&artist, "id = ?", artistId).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "An error occurred while fetching that artist"})
		return
	}
	c.IndentedJSON(http.StatusOK, artist)
}

func CreateArtist(c *gin.Context) {
	var request createArtistRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newArtist := Artist{
		ID:           utils.GenerateID("artist"),
		Name:         request.Name,
		CountryCode:  request.CountryCode,
		ProfileImage: request.ProfileImage,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := utils.DB.Create(&newArtist).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occurred while creating that artist"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newArtist)
}
