package api

import (
	"github.com/ignaciovila/tuiter/src/tweet"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GinServer struct {
	manager *tweet.TweetManager
}

func NewGinServer(manager *tweet.TweetManager) *GinServer {
	return &GinServer{manager}
}

func (server *GinServer) StartGinServer() {
	router := gin.Default()

	router.GET("/listTweets", server.listTweets)
	router.GET("/listTweets/:user", server.listTweets)
	router.POST("publishTweet", server.publishTweet)
	router.POST("publishImageTweet", server.publishImageTweet)
	router.POST("publishQuoteTweet", server.publishQuoteTweet)

	go router.Run()
}

func (server *GinServer) listTweets(c *gin.Context){
	tweets := server.manager.GetTweets()

	for _, tw := range tweets {
		c.String(http.StatusOK, tw.String() + "\n")
	}
}

func (server *GinServer) publishTweet(c *gin.Context) {
	var twt GinTweet
	c.Bind(&twt)

	tweetToPublish := tweet.NewTextTweet(twt.User, twt.Text)

	id, err := server.manager.PublishTweet(tweetToPublish)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
	}else{
		c.JSON(http.StatusOK, struct {
			Id int
		}{id})
	}
}

func (server *GinServer) publishImageTweet(c *gin.Context) {
	var twt GinTweet
	c.Bind(&twt)

	tweetToPublish := tweet.NewImageTweet(twt.User, twt.Text, twt.Url)

	id, err := server.manager.PublishTweet(tweetToPublish)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
	}else{
		c.JSON(http.StatusOK, struct {
			Id int
		}{id})
	}
}

func (server *GinServer) publishQuoteTweet(c *gin.Context) {
	var twt GinTweet
	c.Bind(&twt)

	pId, _ := strconv.Atoi(twt.Id)
	qTwt := server.manager.GetTweetById(pId)

	tweetToPublish := tweet.NewQuoteTweet(twt.User, twt.Text, qTwt)

	id, err := server.manager.PublishTweet(tweetToPublish)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
	}else{
		c.JSON(http.StatusOK, struct {
			Id int
		}{id})
	}
}