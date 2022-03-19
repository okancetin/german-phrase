package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/okancetin/german-phrase/cmd/api/cmd/config"
	"github.com/okancetin/german-phrase/cmd/api/cmd/entity"
	"net/http"
	"os"
	"time"
)

import (
	"github.com/gin-gonic/gin"
)

var Phrases []*entity.Phrase

var phrases = []*entity.Phrase{
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

	c.IndentedJSON(http.StatusOK, getAllPhrasesFromRedis())
	//c.IndentedJSON(http.StatusOK, phrases)
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

func getAllPhrasesFromRedis() (phrase string) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	//ctx := context.Background()
	phrase = client.Get("1").Val()
	return
}

// getPhrasesFromRedis responds with the list of all phrases as JSON.
func getPhrasesFromRedis(c *gin.Context) {
	config.NewRedisClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"), time.Second*60000)
	c.IndentedJSON(http.StatusOK, phrases)
}

func getPhrasesFromDataSource() {
	redisAddress := fmt.Sprintf("%s", os.Getenv("REDIS_URL"))
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
