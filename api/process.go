package api

import (
	"disco/prom"
	"disco/structs"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ProcessResourcesRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*ModernContext)
		pid := ctx.Param("pid")
		status, err := prom.GetProcessStatus(pid)

		if err != nil {
			fmt.Printf("%s", err)
			return ctx.Failure(structs.Response{
				Success: false,
				Message: "Process not found, could not fetch resource data",
			}, 404)
		}

		return ctx.SuccessAny(structs.ResponseAny{
			Data:    status,
			Success: true,
		})
	}
}
