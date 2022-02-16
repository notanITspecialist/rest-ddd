package main

import "rest-ddd/internal/app"

func main() {
	application := app.NewApplications()
	application.Run()
}
