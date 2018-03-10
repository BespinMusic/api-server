// main_test.go

package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"."
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS songs
(
id SERIAL,
title TEXT NOT NULL,
artist TEXT NOT NULL,
album TEXT NOT NULL,
CONSTRAINT songs_pkey PRIMARY KEY (id)
)`

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func clearTable() {
	a.DB.Exec("DELETE FROM songs")
	a.DB.Exec("ALTER SEQUENCE songs_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/songs", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentSong(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/song/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Song not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Song not found'. Got '%s'", m["error"])
	}
}

func TestCreateSong(t *testing.T) {
	clearTable()

	payload := []byte(`{"title":"test song","album":"My Programmatic Romance"}`)

	req, _ := http.NewRequest("POST", "/song", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "test song" {
		t.Errorf("Expected song name to be 'test song'. Got '%v'", m["title"])
	}

	if m["album"] != "My Programmatic Romance" {
		t.Errorf("Expected song price to be 'My Programmatic Romance'. Got '%v'", m["album"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected song ID to be '1'. Got '%v'", m["id"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
