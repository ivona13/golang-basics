package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: "20.99"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: "21.99"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: "18.99"},
}

func main() {
	r := gin.Default()
	setupRouter(r)
	applyMiddleware(r)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

func setupRouter(r *gin.Engine) {
	r.GET("/ping", pingServer)
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumByID)
}

func pingServer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Panic(err)
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Before request %s\n", c.Request.URL.Path)

		c.Next()

		log.Printf("After request %s: status %d\n", c.Request.URL.Path, c.Writer.Status())
	}
}

func applyMiddleware(r *gin.Engine) {
	r.Use(loggingMiddleware())
	r.Use(gin.Recovery())
}
