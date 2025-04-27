package main

import (
	"context"
	noteHandler "github.com/amir-mirjalili/notes.git/notes/handler"
	noteRepository "github.com/amir-mirjalili/notes.git/notes/repository"
	noteService "github.com/amir-mirjalili/notes.git/notes/service"
	userHandler "github.com/amir-mirjalili/notes.git/users/handler"
	userRepository "github.com/amir-mirjalili/notes.git/users/repository"
	userService "github.com/amir-mirjalili/notes.git/users/service"

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

	userRepo := userRepository.NewUserRepository(conn)
	userServe := userService.NewService(userRepo)
	userHandler := userHandler.NewUserHandler(userServe)

	noteRepo := noteRepository.NewNoteRepository(conn)
	noteServe := noteService.NewService(noteRepo)
	noteHandle := noteHandler.NewNoteHandler(noteServe)

	e := echo.New()

	e.GET("/users", userHandler.List)

	e.POST("/notes", noteHandle.Create)

	e.Logger.Fatal(e.Start(":8080"))

}
