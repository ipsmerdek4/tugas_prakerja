package main

import (
	"os"
	"tugas_prakerja/config"
	"tugas_prakerja/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	// fmt.Println("tes")
	config.ConnDatabase()
	e := echo.New()
	e = routers.InitRoute(e)
	e.Start(":8001")
}

func getPORT() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}
