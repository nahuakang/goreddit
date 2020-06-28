package main

import (
	"log"
	"net/http"

	"github.com/nahuakang/goreddit/postgres"
	"github.com/nahuakang/goreddit/web"
)

func main() {
	dsn := "postgres://postgres:secret@localhost/postgres?sslmode=disable"
	store, err := postgres.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	sessions, err := web.NewSessionManager(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 32-byte CSRF Key
	csrfKey := []byte("01234567890123456789012345678901")
	h := web.NewHandler(store, sessions, csrfKey)
	http.ListenAndServe(":3000", h)
}
