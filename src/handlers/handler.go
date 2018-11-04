package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lornasong/aws-cloud-directory-visual/src/visual"
)

const (
	rootNodeID = "/"
)

// FindRoot returns the node profile for the root node
func FindRoot(v *visual.Visual) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		out, err := v.GenerateProfile(rootNodeID)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, out)
	}
}

// Find returns the node profile for the id specified
func Find(v *visual.Visual) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		out, err := v.GenerateProfile(id)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, out)
	}
}
