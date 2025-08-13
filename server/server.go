package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_"github.com/gin-contrib/cors"
)

//WebServer struct
type WebServer struct {
	// gin server reference
	server *gin.Engine
	// channel for cleanup notification
	cleanUp chan int
}

//Create this will load default server config
func Create(cleanup chan int) *WebServer {
	server := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	
	server.Use(CORSMiddleware())
	
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	
	server.Use(gin.Recovery())

	ginServer := &WebServer{
		server:  server,
		cleanUp: cleanup,
	}
	return ginServer
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func (ws *WebServer)Run(port string){
	ws.server.Run(port)
}
