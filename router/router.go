package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações default do Gin
	r := gin.Default()
	// Definindo uma Rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Listen and serve on 0.0.0.0:8000
	r.Run()
}
