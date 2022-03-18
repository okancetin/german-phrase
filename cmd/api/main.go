package main

import "net/http"

import (
	"github.com/gin-gonic/gin"
)

type Phrase struct {
	Id          string      `json:"id"`
	Content     string      `json:"content"`
	Title       string      `json:"title"`
	Link        string      `json:"link"`
	Translation Translation `json:"translation"`
}

type Translation struct {
	En string `json:"en"`
	De string `json:"de"`
}

var Phrases []Phrase

var phrases = []Phrase{
	{Title: "Hello", Content: "Article Content"},
	{Title: "Hello 2", Content: "Article Content"},
}

func main() {
	router := gin.Default()
	router.GET("/phrases", getPhrases)
	router.GET("/phrases/:id", getPhraseByID)

	//router.Run("localhost:8080")
	router.Run()
}

// getPhrases responds with the list of all phrases as JSON.
func getPhrases(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, phrases)
}

// getPhraseByID locates the phrase whose ID value matches the id
func getPhraseByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range phrases {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "phrases not found"})
}
