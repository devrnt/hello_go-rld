package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type Album struct {
	// specify what a field’s name should be when the struct’s contents are serialized into JSON
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(context *gin.Context) {
	// serialize the struct into JSON and add it to the response
	context.JSON(http.StatusOK, albums)
}

func postAlbum(context *gin.Context) {
	var newAlbum Album

	// bind the received JSON to newAlbum
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)

	context.JSON(http.StatusCreated, newAlbum)
}

func getAlbum(context *gin.Context) {
	var id = context.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			context.JSON(http.StatusOK, album)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// export function by capitalizing
func InstallAlbumRoutes(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)
}
