package routers

import (
	"tugas_prakerja/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/donasi_buku", controllers.DonasiBukuController)
	e.POST("/donasi_buku", controllers.CreateDonasiBukuController)
	e.GET("/donasi_buku/:nim", controllers.ItemDonasiBukuController)
	e.GET("/delete_donasi_buku/:nim", controllers.DeleteDonasiBukuController)
	return e
}
