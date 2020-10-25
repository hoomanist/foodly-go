package main

import (
	"context"
	"foodly/ent"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var curser *ent.Client

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	curser = client
	r := gin.Default()
	r.MaxMultipartMemory = 1 << 20
	r.GET("/ping", register)
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/submit/food", SubmitFood)
	r.Run(":5000")
}
