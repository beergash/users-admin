// main
package main

import (
	"came-users/app"
	"came-users/config"
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
