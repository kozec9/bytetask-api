// cmd/myapp/main.go

package main

import (
	"bytetask-api/internal/app"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	myApp := app.NewApp()
	defer myApp.Close()

	myApp.Run(port)
}
