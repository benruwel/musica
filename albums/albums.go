package albums

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"musica.com/utils"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func FindOne(c *gin.Context) {
	var album Album
	albumId := c.Param("id")

	err := utils.Conn.QueryRow(context.Background(), "select (id, title, artist, price) from albums where id=$1", albumId).Scan(&album)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "An error occurred while fetching that album"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func ListAlbums(c *gin.Context) {
	// must declare what columns to return
	rows, _ := utils.Conn.Query(context.Background(), "select (id, title, artist, price) from albums")
	albums, err := pgx.CollectRows(rows, pgx.RowTo[Album])

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while querying albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err := utils.Conn.Exec(context.Background(), "insert into albums(id, title, artist, price) values ($1, $2, $3, $4)", utils.GenerateID("album"), newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
