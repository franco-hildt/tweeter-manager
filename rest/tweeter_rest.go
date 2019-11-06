package rest

import (
	"net/http"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/service"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	TweetManager *service.TweetManager
}

var ginServer GinServer

func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	ginServer.TweetManager = tweetManager

	return &ginServer
}

func (*GinServer) StartGinServer() {
	router := gin.Default()

	router.GET("/tweet", getTweets)
	router.POST("/tweet", saveTweet)

	go router.Run()
}

func getTweets(c *gin.Context) {
	c.JSON(http.StatusOK, ginServer.TweetManager.GetTweets())
}

func saveTweet(c *gin.Context) {
	var tweet domain.TextTweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ginServer.TweetManager.PublishTweet(tweet)

}
