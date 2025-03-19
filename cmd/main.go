package main

import (
	"fmt"

	server "arczed/internal/interface"
)

func main() {
	server := server.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
