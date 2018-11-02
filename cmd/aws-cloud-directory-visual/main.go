package main

import (
	"github.com/labstack/echo"
	"github.com/lornasong/aws-cloud-directory-visual/src/handlers"
)

func main() {

	e := echo.New()

	e.GET("/visualize", handlers.Visualize)
	e.Start(":8000")
}
