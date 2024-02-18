package main

import (
	"fmt"
	"html/template"
	"net/http"

	"frontendmasters.com/go/museum/api"
	"frontendmasters.com/go/museum/data"
	"github.com/gin-gonic/gin"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Service error"))
	}
	html.Execute(w, data.GetAll())
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
		
	err := http.ListenAndServe(":8080", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}