package main

import (
	"DatingApp/src/handler"
	"DatingApp/src/middleware"
	"DatingApp/src/models"
	"DatingApp/src/repositories"
	"DatingApp/src/services"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	env := models.SetEnv()

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", env.DB_USER, env.DB_PASS, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	db, err := sql.Open(env.DB_TYPE, dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	repo := repositories.Init(repositories.Param{Db: db})

	srv := services.Init(services.Param{Repositories: repo})

	midlwre := middleware.Init(middleware.InitParam{Service: srv})

	hndlr := handler.Init(handler.InitParam{Service: srv, Middleware: midlwre})

	hndlr.Run()

}
