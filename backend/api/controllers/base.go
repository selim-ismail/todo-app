package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
	"todo-app/api/middlewares"
)

type Server struct {
	Context  context.Context
	Database *mongo.Database
	Router   *gin.Engine
}

var errList = make(map[string]string)

func (server *Server) Initialize(
	dbDriver string,
	dbUser string,
	dbPass string,
	dbName string,
) {
	fmt.Println("initializing server...")

	if dbDriver == "mongodb" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbUri := fmt.Sprintf("mongodb+srv://%s:%s@todoappcluster0-tbuqv.mongodb.net/test?retryWrites=true&w=majority", dbUser, dbPass)

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))

		if err != nil {
			log.Fatal(err)
		}

		server.Database = client.Database(dbName)
	} else {
		log.Fatal("Unknown Database Driver")
	}

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
