package main

import (
	"fmt"
	"github.com/grandleemon/go-test-app.git/internal/router"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	fmt.Println("Starting server on port 8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
