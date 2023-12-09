package main

import (
	"fmt"
	"form/internal/server"
)

func main() {

	server := server.NewServer()

	fmt.Println("listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
