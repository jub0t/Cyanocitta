package api

import (
	cserde "disco/serde"
	"disco/structs"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

// Serializes structs.Response Into Binary Message Pack Data
func (c *CustomContext) Success(resp structs.Response) {
	parsed, err := cserde.Se(resp)

	if err != nil {
		println("Unknwon error occured while parsing response", parsed)
	}

	c.String(200, string(parsed[:]))
}

func (c *CustomContext) Failure(resp structs.Response, status int) {
	parsed, err := cserde.Se(resp)

	if err != nil {
		println("Unknwon error occured while parsing response", parsed)
	}

	c.String(status, string(parsed[:]))
}

func Serdeware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return next(cc)
	}
}
