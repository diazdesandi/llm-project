package routes

import (
	"fmt"

	"github.com/diazdesandi/llm-project/backend/internal/auth"
	"github.com/diazdesandi/llm-project/backend/internal/middlewares"
	"github.com/diazdesandi/llm-project/backend/internal/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Body struct {
	Name string `path:"name" maxLength:"30" required:"true" example:"John Doe"`
}

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func SetupRoutes(mh model.Handler, ah auth.Handler) *gin.Engine {

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middlewares.Metrics())
	// r.Use(middlewares.CORS())

	// Metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Check backend is running, sample endpoint
	r.GET("/greeting/:name", func(c *gin.Context) {
		name := c.Param("name")
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", name)
		c.JSON(200, resp)
	})

	// model/generate
	// modelGroup := r.Group("/model")
	// {
	// 	modelGroup.GET("/generate", mh.Handler)
	// }
	r.POST("/model/generate", mh.Handler)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", ah.SignupHandler)
		authGroup.POST("/login", ah.SignInHandler)
	}

	return r
}
