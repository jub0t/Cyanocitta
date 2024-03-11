package api

import (
	database "disco/db"
	"disco/structs"

	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		headers := ctx.Request().Header
		token := headers.Get("Token")

		if len(token) <= 0 {
			return ctx.JSON(401, structs.Response{
				Success: false,
				Message: "Invalid Token, Unauthenticated",
			})
		}

		user, err := database.GetUserByToken(database.DB, token)
		if err != nil {
			return ctx.JSON(500, structs.Response{
				Success: false,
				Message: "Unauthorized - Please login",
			})
		}

		if user.ID != 0 {
			ctx.Set("User", user)
		} else {
			return ctx.JSON(401, structs.Response{
				Success: false,
				Message: "Unauthenticated",
			})
		}

		return next(ctx)
	}
}
