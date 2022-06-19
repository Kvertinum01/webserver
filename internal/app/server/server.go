package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer(config *Config) error {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "index.html",
			gin.H{"Title": "Обо мне"},
		)
	})

	router.GET("/contacts", func(c *gin.Context) {
		c.HTML(
			http.StatusOK, "contacts.html",
			gin.H{"Title": "Контакты"},
		)
	})

	return router.Run(config.Addr)
}
