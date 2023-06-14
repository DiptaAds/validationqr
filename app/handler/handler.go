package handler

import (
	"crud_module/app/model"
	"fmt"

	"time"

	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QRHandler struct {
	Database *mongo.Database
}

var datainput map[string][]string

func (qr QRHandler) CreateQR(c echo.Context) error {

	var dataaQR model.QRCode
	err := c.Bind(&dataaQR)
	if err != nil {
		return err
	}
	uuids := make([]string, 0)
	for i := 0; i < dataaQR.Amount; i++ {
		uuid := uuid.New().String()
		uuids = append(uuids, uuid)
	}
	for _, uuid := range uuids {
		dataaQR.UUID = uuid
		dataaQR.CreatedAt = time.Now()
		_, err := qr.Database.Collection("qrku").InsertOne(c.Request().Context(), dataaQR)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]interface{}{
		"success": true,
		"data":    uuids,
	})
}
func (qr QRHandler) CreateProduct(c echo.Context) error {
	dataProduct := model.Product{}
	err := c.Bind(&dataProduct)
	if err != nil {
		return err
	}
	isrt, err := qr.Database.Collection("qrku").InsertOne(c.Request().Context(), dataProduct)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Failed input data"})
	}
	fmt.Println("Data berhasil masuk")

	return c.JSON(200, isrt)
}
func (qr QRHandler) GetAllUser(c echo.Context) error {

	qrData := make([]model.QRCode, 0)
	cursor, err := qr.Database.Collection("qrku").Find(c.Request().Context(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(c.Request().Context())
	for cursor.Next(c.Request().Context()) {
		var data model.QRCode
		err := cursor.Decode(&data)
		if err != nil {
			return err
		}
		qrData = append(qrData, data)
	}

	return c.JSON(200, map[string]interface{}{
		"success": true,
		"data":    qrData,
	})
}
func (qr QRHandler) Update(c echo.Context) error {
	dataaQR := model.QRCode{}
	err := c.Bind(&dataaQR)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Get the UUID from the URL parameter
	uuid := c.Param("uuid")

	// Set the UUID in the dataaQR struct
	dataaQR.UUID = uuid

	// Set UpdatedAt field to current time

	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": dataaQR}

	_, err = qr.Database.Collection("qrku").UpdateOne(c.Request().Context(), filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (qr QRHandler) Delete(c echo.Context) error {
	// Get the UUID from the URL parameter
	name := c.Param("name")

	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"DeletedAt": time.Now()}}

	_, err := qr.Database.Collection("qrku").UpdateOne(c.Request().Context(), filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
