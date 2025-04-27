package repository

import (
	"context"
	"github.com/amir-mirjalili/notes.git/notes/params"
	"github.com/jackc/pgx/v5"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteRepository struct {
	Conn *pgx.Conn
}

func NewNoteRepository(conn *pgx.Conn) *NoteRepository {
	return &NoteRepository{conn}
}

func (r *NoteRepository) Create(note params.AddNotRequest) error {
	r.Conn.Exec(context.Background(), "insert into notes (user_id,title,content) VALUES ($1,$2,$3) ", note.UserID, note.Title, note.Content)
	return nil
}

func (r *NoteRepository) Update(id int, note Note) (Note, error) {
	r.Conn.Exec(context.Background(), "update notes set title=$2,content=$3  where id = $1", id, note.Title, note.Content)
	note.ID = id
	return note, nil
}

func (r *NoteRepository) Delete(id int) error {
	r.Conn.Exec(context.Background(), "delete from notes where id = $1", id)
	return nil
}

func (r *NoteRepository) GetByID(id int) (Note, error) {
	var note Note

	row := r.Conn.QueryRow(context.Background(), "select * from notes where id = $1", id)
	err := row.Scan(&note.ID, &note.UserID, &note.Title, &note.Content)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepository) GetList() ([]Note, error) {
	rows, err := r.Conn.Query(context.Background(), "SELECT id, user_id, title, content FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
