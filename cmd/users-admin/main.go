// main
package main

import (
	"log"
	"os"
	"users-admin/app"
	"users-admin/config"
)

func init() {
}

// Start application
func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		log.Printf("defined port %s", PORT)
		PORT = "8080"
	}
	log.Println("Welcome to Application")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + PORT)

}
