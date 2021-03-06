package main

import (
	"database/sql"
	"errors"
)

type song struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Album string `json:"album"`
}

func (s *song) getSong(db *sql.DB) error {
	return db.QueryRow("SELECT nm, cd FROM songs WHERE id=$1",
		s.ID).Scan(&s.Title, &s.Album)
}

func (s *song) updateSong(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *song) deleteSong(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (s *song) createSong(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO songs(nm, cd) VALUES($1, $2) RETURNING id",
		s.Title, s.Album).Scan(&s.ID)

	if err != nil {
		return err
	}

	return nil
}

func getSongs(db *sql.DB, start, count int) ([]song, error) {
	rows, err := db.Query(
		"SELECT id, cd, nm FROM songs LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	songs := []song{}

	for rows.Next() {
		var s song
		if err := rows.Scan(&s.ID, &s.Album, &s.Title); err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}

	return songs, nil
}
