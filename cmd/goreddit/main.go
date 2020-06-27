package main

import (
	"log"
	"net/http"

	"github.com/nahuakang/goreddit/postgres"
	"github.com/nahuakang/goreddit/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// 32-byte CSRF Key
	csrfKey := []byte("01234567890123456789012345678901")
	h := web.NewHandler(store, csrfKey)
	http.ListenAndServe(":3000", h)
}
