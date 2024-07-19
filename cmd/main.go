package main

import (
	"zg_backend/internal/app"
)

// @title ZG Backend API
// @version 1
// @description Backend app for ZmeyGorynych project. Written for fun :)
// @host localhost:8080
// @license none
// @BasePath /api/v1
func main() {
	app.NewApp().Run()
}
