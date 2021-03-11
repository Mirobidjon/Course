package main

import (
	"github.com/Mirobidjon/course"
	"github.com/Mirobidjon/course/pkg/handler"
	"github.com/Mirobidjon/course/pkg/repository"
	"github.com/Mirobidjon/course/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("dport"),
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		DbName:   os.Getenv("dbname"),
		SslMode:  os.Getenv("sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed initsialized database : %s", err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("env file port not found!")
	}
	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handle := handler.NewHandler(serv)

	svr := new(course.Server)

	if err := svr.Run(port, handle.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server : %s", err.Error())
	}
}
