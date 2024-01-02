package main

import (
	"fmt"
	_ "github.com/MrRytis/key-store/docs"
	"github.com/MrRytis/key-store/internal/handler"
	"github.com/MrRytis/key-store/internal/service"
	"github.com/MrRytis/key-store/internal/storage"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"
)

// @title Key store API
// @version 1.0
// @description This is a simple key store application.

// @contact.name Rytis Jančeris
// @contact.email rytis.janceris@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
func main() {
	e := echo.New()

	s := storage.NewStorage()
	h := handler.NewHandler(s)

	checker := service.NewExpirationChecker(s)
	go checker.Start()

	e.GET("/api/v1/store", h.GetAllValues)
	e.GET("/api/v1/store/:key", h.GetValueByKey)
	e.POST("/api/v1/store", h.StoreValue)
	e.DELETE("/api/v1/store/:key", h.DeleteValue)

	e.GET("/docs/*", echoSwagger.WrapHandler)

	fmt.Println("Starting server on port", os.Getenv("PORT"))
	fmt.Printf("Swagger docs available at http://localhost:%s/docs/index.html", os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
