package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type (
    Tenant struct {
        ID     string
        AuthDB string
    }
)

func (t *Tenant) TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        t.ID = c.QueryParam("tenant")
        return next(c)
    }
}

func (t *Tenant) Connect(authDB string) echo.MiddlewareFunc {
    t.AuthDB = authDB
    return t.TenantMiddleware
}

func (t *Tenant) DB(dbName string) string {
    return fmt.Sprintf("tenant id: %s\ndb name: %s", t.ID, dbName)
}

func main() {
    authDB := "authDB"

	e := echo.New()

    // Debug mode
    e.Debug = true

    // Tenant
    t := &Tenant{}
	e.Use(t.Connect(authDB))

    // Handler
    h := handler{t}
    e.GET("/", h.Index)

    // Start server
	e.Logger.Fatal(e.Start(":8000"))
}
