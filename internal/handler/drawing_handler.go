package handler

import (
	"net/http"

	"github.com/adi-ym-0001/go-awp-backend/internal/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetDrawings(ctx echo.Context) error {
	drawings, err := s.DrawingUC.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, drawings)
}

func (s *Server) PutLocationsIdDrawings(ctx echo.Context, id string) error {
	var drawings []model.Drawing
	if err := ctx.Bind(&drawings); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストボディの解析に失敗しました"})
	}

	if err := s.LocationUC.UpdateDrawings(ctx.Request().Context(), id, drawings); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
