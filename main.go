package main

import (
	"github.com/tdaughton/foo/internal"
	"log"
)

func main() {
	log.Println("Listening on port 8080")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
