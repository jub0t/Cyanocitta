package api

import (
	database "disco/db"
	"disco/structs"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TokenMiddleware(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		headers := ctx.Request().Header
		token := headers.Get("Token")

		if len(token) <= 0 {
			return ctx.JSON(401, structs.Response{
				Success: false,
				Message: "Invalid Token, Unauthenticated",
			})
		}

		if user, err := database.GetUserByToken(db, token); err != nil {
			return ctx.JSON(502, structs.Response{
				Success: false,
				Message: "Internal Server Error",
			})
		} else {
			if (user != structs.User{}) {
				ctx.Set("User", &user)
			} else {
				return ctx.JSON(401, structs.Response{
					Success: false,
					Message: "Unauthenticated",
				})
			}
		}

		return next(ctx)
	}
}
