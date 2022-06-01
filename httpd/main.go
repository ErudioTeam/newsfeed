package main

import (
	"net/http"
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")
		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func setupNewsfeedRouter() *gin.Engine {
	feed := newsfeed.New()

	r := gin.Default()
	r.GET("/ping", handler.PingGetv1())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	return r
}

func setupNewsRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	//r.GET("/ping", handler.PingGet)

	// Pingv1 test
	r.GET("/ping", handler.PingGetv1())

	// add news feed

	//r.GET("/news", handler.Newsfeed())
	return r
}

func main() {
	/* 	fmt.Println("news feeder")

	   	feed := newsfeed.New()

	   	fmt.Println(feed)
	   	feed.Add(newsfeed.Item{"Hello", "Erudio learner"})
	   	fmt.Println(feed) */

	/* 	r := gin.Default()
	   	r.GET("/ping", func(c *gin.Context) {
	   		c.JSON(200, gin.H{
	   			"message": "pong",
	   		})
	   	})
	   	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	*/

	/* 	r := setupNewsRouter()
	   	// Listen and Server in 0.0.0.0:8080
	   	r.Run(":8081") */

	r := setupNewsfeedRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
