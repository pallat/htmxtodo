package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/todos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "todos.html", todos)
	})
	r.DELETE("/todos", func(ctx *gin.Context) {
		b, _ := httputil.DumpRequest(ctx.Request.Clone(ctx.Request.Context()), true)
		fmt.Println(string(b))
	})
	// c.Redirect(http.StatusFound, "/")

	log.Println("Starting server on :8080")
	r.Run()
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var todos = []Todo{
	{1, "Learn Go", false},
	{2, "Build a Todo App", false},
}
