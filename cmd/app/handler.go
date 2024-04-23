package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
    handler struct {
        *Tenant
    }
)

func (h *handler) Index(c echo.Context) error {
    return c.String(http.StatusOK, h.DB("user"))
}
