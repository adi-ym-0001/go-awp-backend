package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetLocations(ctx echo.Context) error {
	locations, err := s.LocationUC.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, locations)
}
