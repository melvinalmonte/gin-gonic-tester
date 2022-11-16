package main

import (
	"fmt"
	"gin-gonic-tester/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to this small Golang API")
	})
	v1 := router.Group("")
	routes.Blueprint(v1)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("An error has occurred when starting server %s", err)
	}
}
