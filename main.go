package main

import (
	"log"
	"net/http"
	"strconv"

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
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		todos = append(todos[:i], todos[i+1:]...)
		c.Redirect(http.StatusFound, "/")
	})

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

type Todos []Todo

var todos Todos = []Todo{
	{1, "Learn Go", false},
	{2, "Build a Todo App", false},
}

func (t Todos) Remove(i int) {
	h := len(t) - 1
	l := 0
	if h == l {
		if 
	}
}
