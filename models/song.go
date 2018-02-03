package main

import (
	"database/sql"
	"errors"
)

type song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json: "album"`
}

func (s *song) getSong(db *sql.DB) error {
	return db.QueryRow("SELECT title, artist, album FROM songs WHERE id=$1",
		s.ID).Scan(&s.Title, &s.Artist, &s.Album)
}

func (s *song) updateSong(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *song) deleteSong(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *song) createSong(db *sql.DB) error {
	return db.Exec("INSERT INTO songs ()")
}

func getSongs(db *sql.DB, start, count int) ([]song, error) {
	return nil, errors.New("Not implemented")
}
