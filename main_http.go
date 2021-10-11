package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	echoMid "github.com/labstack/echo/v4/middleware"
)

var DefaultPort uint16 = 8080

func (s *Service) HttpServerMain() {
	e := echo.New()

	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	// hGroup := e.Group("/api/health")
	// v1Group := e.Group("/api/v1")

	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	}

	port = DefaultPort

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))

}
