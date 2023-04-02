package server

import (
	"bookstore/Infra/database"
	"bookstore/Infra/database/postgresql"
	"bookstore/app/api/dependency"

	"github.com/labstack/echo/v4"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	e := echo.New()
	dsn := postgresql.PostgreSQLConnect()
	db, err := database.DB("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dependency.Dependency(db, e)

	e.GET("/", healthCheck)

	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.JSON(200, "OK")
}
