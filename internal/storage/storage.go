package storage

import (
	"database/sql"
	"log"
	"talkye/internal/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	Config *config.DBConfig
	DbConn *sql.DB
}

func NewStorage(config *config.DBConfig) *Storage {
	return &Storage{
		Config: config,
	}
}

func (s *Storage) Initialize() error {
	log.Printf("Initializing storage...")

	conn, err := sql.Open("postgres", s.Config.ToConnectionString())
	if err != nil {
		log.Printf("Error initializing database: %s", err)
		return err
	}

	s.DbConn = conn

	return nil
}

func (s *Storage) Close() {
	log.Printf("Closing storage...")
	if s.DbConn != nil {
		s.DbConn.Close()
	}
}

func (s *Storage) GetDatabase() (*sql.DB, error) {
	if s.DbConn == nil {
		err := s.Initialize()
		if err != nil {
			return nil, err
		}
	}

	return s.DbConn, nil
}
