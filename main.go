package main

import (
	"fmt"
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
	e.Start(getPORT())
}

func getPORT() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return fmt.Sprintf(":%s", port)
}
