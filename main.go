package main

import (
	"context"
	"github.com/amir-mirjalili/notes.git/handler"
	repository "github.com/amir-mirjalili/notes.git/repository/user"
	"github.com/amir-mirjalili/notes.git/service/user"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://amir:1379@localhost:5432/note_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())
	repo := repository.NewUserRepository(conn)
	service := user.NewService(repo)
	userHandler := handler.NewUserHandler(service)

	e := echo.New()
	e.GET("/", userHandler.List)

	e.Logger.Fatal(e.Start(":8080"))

}
