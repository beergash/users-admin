// main
package main

import (
	"io/ioutil"
	"os"
	"users-admin/app"
	"users-admin/app/logger"
	"users-admin/config"
)

func init() {
}

// Start application
func main() {
	logger.InitLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		logger.Info.Printf("defined port %s", PORT)
		PORT = "8080"
	}
	logger.Info.Println("Welcome to Application")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + PORT)

}
