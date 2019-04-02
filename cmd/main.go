// main
package main

import (
	"users-admin/app"
	"users-admin/config"
	"log"
)

func init() {
}

func main() {
	log.Println("Welcome to Came Application")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")

}
