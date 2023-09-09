package main

import (
	"log"

	"github.com/highcloud100/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
