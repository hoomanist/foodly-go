package main

import (
	"context"
	"foodly/ent"
	"log"
  "database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
  "github.com/facebook/ent/dialect"
  entsql "github.com/facebook/ent/dialect/sql"
  _ "github.com/jackc/pgx/v4/stdlib"
)

var cursor *ent.Client
// Open new connection
func Open(databaseUrl string) *ent.Client {
    db, err := sql.Open("pgx", databaseUrl)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
    return ent.NewClient(ent.Driver(drv))
}

func main() {
  client := Open("postgresql://postgres:hooman86@127.0.0.1/foodly")
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	cursor = client
	r := gin.Default()
	r.MaxMultipartMemory = 1 << 20
	r.GET("/ping", ping)
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/submit/food", SubmitFood)
	r.POST("/upload/image", UploadImage)
	r.GET("/q/image/:name", ServeImage)
	r.GET("/q/restbycity", QueryRestaurants)
	r.GET("/q/foodbyRTi", FoodsByRest)
	r.POST("/submit/comment", CreateCamment)
	r.GET("/q/comment", QueryComment)
	r.POST("/vote/food", Vote)
  err := r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
