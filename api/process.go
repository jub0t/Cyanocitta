package api

import (
	"disco/prom"
	"disco/structs"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ProcessResourcesRoute() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pid := ctx.Param("pid")
		status, err := prom.GetProcessStatus(pid)

		if err != nil {
			fmt.Printf("%s", err)
			return ctx.JSON(404, structs.Response{
				Success: false,
				Message: "Process not found, could not fetch resource data",
			})
		}

		return ctx.JSON(200, structs.ResponseAny{
			Data:    status,
			Success: true,
		})
	}
}
