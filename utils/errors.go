package utils

import "github.com/labstack/echo"

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewError create new error object
func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}
