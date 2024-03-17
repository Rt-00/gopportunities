package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	r := gin.Default()

	// Initializa Routes
	initializeRoutes(r)

	// Listen and serve on 0.0.0.0:8000
	r.Run()
}
