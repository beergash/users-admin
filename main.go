// main
package main

import (
	"came-users/app"
	_ "came-users/app/model"
	"came-users/config"
	"fmt"
	_ "time"
)

func main() {
	fmt.Println("Welcome to Came Application")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")

}
