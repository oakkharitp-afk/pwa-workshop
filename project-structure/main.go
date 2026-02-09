package main

import (
	"os"
	"pwa-workshop/project-structure/database"

	"pwa-workshop/project-structure/handler"
	"pwa-workshop/project-structure/logger"
	myMid "pwa-workshop/project-structure/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	logger.Init("development")
	logger.SetLevel("debug")
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error to read env, %v", err)
	}
}
func main() {

	var (
		sysApiKey = os.Getenv("SYSTEM_API_KEY")
	)

	mgCli, err := database.ConnectMongoClient(os.Getenv("DB_MG_URI"))
	if err != nil {
		logger.Fatal("error to connect mongo database, ", err)
	}
	pgCli, err := database.ConnectSQLClient(os.Getenv("DB_PG_URI"), "postgres")
	if err != nil {
		logger.Fatal("error to connect postgres database, ", err)
	}

	db := database.NewDatabase(mgCli, pgCli)
	h := handler.NewHandler(db)

	e := echo.New()
	e.Use(middleware.Recover())

	// api key authentication
	e.Use(myMid.APIKeyMiddleware(sysApiKey))

	// api version 1
	v1 := e.Group("/api/1.0")
	// base paths
	var (
		pipes = "/pipes"
	)
	// path configuration
	v1.GET("/hello-worlds", h.HelloWorld)
	v1.POST(pipes, h.CreatePipe)
	v1.GET(pipes, h.GetPipes)
	v1.POST(pipes, h.CreatePipe)
	v1.POST(pipes, h.CreatePipe)

	if err := e.Start(":8080"); err != nil {
		logger.Fatal("error to start http server, ", err)
	}

}
