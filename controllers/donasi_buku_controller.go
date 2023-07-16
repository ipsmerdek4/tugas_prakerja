package controllers

import (
	"net/http"
	"tugas_prakerja/config"
	"tugas_prakerja/models"

	"github.com/labstack/echo/v4"
)

func DeleteDonasiBukuController(c echo.Context) error {
	nim := c.Param("nim")
	var donasi_buku []models.T_Donasi_Buku

	result := config.DB.Where("nim = ?", nim).Delete(&donasi_buku)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal Menghapus Data", Status: false, Data: nil,
		})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{
			Message: "Berhasil Meghapus Data", Status: true, Data: donasi_buku,
		})
	}
}

func ItemDonasiBukuController(c echo.Context) error {
	nim := c.Param("nim")
	var donasi_buku []models.T_Donasi_Buku

	result := config.DB.First(&donasi_buku, nim)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal Mendapatkan Data", Status: false, Data: nil,
		})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{
			Message: "Berhasil Mendapatkan Data", Status: true, Data: donasi_buku,
		})
	}
}

func CreateDonasiBukuController(c echo.Context) error {
	donasi_buku := models.T_Donasi_Buku{}
	c.Bind(&donasi_buku)

	result := config.DB.Create(&donasi_buku)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal Menambah Data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil Menambah Data", Status: true, Data: donasi_buku,
	})

}

func DonasiBukuController(c echo.Context) error {
	var donasi_buku []models.T_Donasi_Buku

	result := config.DB.Find(&donasi_buku)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal Mendapatkan Data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil Mendapatkan Data", Status: true, Data: donasi_buku,
	})

}
