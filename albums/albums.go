package albums

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"musica.com/app/utils"
)

type Album struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	ArtistID  string    `json:"artist_id"`
	AlbumArt  string    `json:"album_art"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAlbumRequest struct {
	Title    string  `json:"title"`
	ArtistID string  `json:"artist_id"`
	AlbumArt string  `json:"album_art"`
	Price    float64 `json:"price"`
}

func FindOne(c *gin.Context) {
	var album Album
	albumId := c.Param("id")

	err := utils.DB.First(&album, "id = ?", albumId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "An error occurred while fetching that album"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func ListAlbums(c *gin.Context) {
	// must declare what columns to return
	var albums []Album
	result := utils.DB.Find(&albums)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while querying albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var request CreateAlbumRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newAlbum := Album{
		ID:        utils.GenerateID("album"),
		Title:     request.Title,
		ArtistID:  request.ArtistID,
		AlbumArt:  request.AlbumArt,
		Price:     request.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := utils.DB.Create(&newAlbum)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
