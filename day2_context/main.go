package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.POST("/", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{
			"username": context.PostForm("username"),
			"passwd":   context.PostForm("pwd"),
		})
	})

	r.Run(":8888")
}
