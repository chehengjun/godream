package main

import (
	"godream/db/postgresql"
	"log"
	"net/http"
	"net/rpc"
)

const (
	URL = "127.0.0.1:5021"
)

func main() {
	var err error
	log.Println("\n\t\t\033[31m[This Is Operation Postgresql Service]\033[0m")

	err = rpc.Register(new(postgresql.RDB))
	if err != nil {
		log.Fatal("Register error:", err)
	}
	rpc.HandleHTTP()

	err = http.ListenAndServe(URL, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
