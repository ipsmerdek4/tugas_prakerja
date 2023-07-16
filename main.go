package main

import (
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
