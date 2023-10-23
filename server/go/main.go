package main

import (
	"go/pg"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)


	pg.OpenPostgres(mux)

	http.ListenAndServe(":5000", handler)
}
