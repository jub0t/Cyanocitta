package api

import (
	cserde "disco/serde"
	"disco/structs"

	"github.com/labstack/echo/v4"
)

type ModernContext struct {
	echo.Context
}

// Serializes structs.Response Into Binary Message Pack Data
func (c *ModernContext) Success(resp structs.Response) {
	parsed, err := cserde.Se(resp)

	if err != nil {
		println("Unknwon error occured while parsing response", parsed)
	}

	c.String(200, string(parsed[:]))
}

// Serializes structs.Response Into Binary Message Pack Data
func (c *ModernContext) SuccessAny(resp structs.ResponseAny) error {
	parsed, err := cserde.Se(resp)

	if err != nil {
		println("Unknwon error occured while parsing response", parsed)
	}

	c.String(200, string(parsed[:]))
	return nil
}

func (c *ModernContext) Failure(resp structs.Response, status int) error {
	parsed, err := cserde.Se(resp)

	if err != nil {
		println("Unknwon error occured while parsing response", parsed)
	}

	c.String(status, string(parsed[:]))
	return nil
}

func Serdeware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &ModernContext{c}
		return next(cc)
	}
}
