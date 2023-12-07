// cmd/myapp/main.go

package main

import (
	"bytetask-api/internal/app"
)

func main() {

	myApp := app.NewApp()
	defer myApp.Close()

	port := myApp.PORT
	if port == "" {
		port = "8000"
	}

	myApp.Run(port)
}
