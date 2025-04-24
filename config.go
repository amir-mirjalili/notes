package main

import "github.com/amir-mirjalili/notes.git/model"

type Config struct {
	PostgresDB model.Config `koanf:"postgres_db"`
}
