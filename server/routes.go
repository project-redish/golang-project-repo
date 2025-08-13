package server

import (
	_ "fmt"
	handle "go-initializer/handler"

	_ "github.com/gin-gonic/gin"
)

// RegisterRoute register routes
func (ws *WebServer) RegisterRoute() {

	ws.server.POST("/generate", handle.GenerateTemplate)
	ws.server.POST("/explore", handle.ExploreApp)
	ws.server.GET("/app-count", handle.AppCounter)
	ws.server.GET("/liveness", handle.Liveness)
	ws.server.GET("/libs", handle.GetSupportedLibraries)

	// Get user value
	//ws.server.GET("/user/:name", handle.GetUserValue)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := ws.server.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

}
