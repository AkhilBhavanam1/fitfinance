package main

import (
	"context"
	"fitfinance/db"
	"fitfinance/graph"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	mongoClient := InitializeDb()

	dbs, _ := mongoClient.ListDatabaseNames(context.Background(), bson.D{})

	e.Logger.Info("Db names", dbs)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DbClient: mongoClient}}))

	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func InitializeDb() *mongo.Client {

	helperClient := db.DBHelperClient{
		ConnString: "mongodb://localhost:32767/",
	}

	err := helperClient.ConnectToDB()
	if err != nil {
		log.Fatal("Cannot connect to db")
	}
	return helperClient.Client
}
