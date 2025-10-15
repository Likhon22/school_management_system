package main

import "school-management-system/internal/bootstrap"

func main() {

	app := bootstrap.NewApp()
	app.Run()
}
