package main

import "rest-ddd/pkg/app"

func main() {
	application := app.NewApplications()
	application.Run()
}
