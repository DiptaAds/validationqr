package main

import (
	"context"
	"crud_module/app/handler"
	"crud_module/app/infrastructure"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := infrastructure.ConnectM(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	qrhandler := handler.QRHandler{
		Database: db,
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Ok")
	})

	e.POST("qr", qrhandler.CreateQR)
	e.POST("product", qrhandler.CreateProduct)
	e.GET("getqrproduct", qrhandler.GetAllUser)
	e.PUT("putuser", qrhandler.Update)
	e.DELETE("deleteuser", qrhandler.Delete)
	e.Logger.Fatal(e.Start(":8080"))
}
