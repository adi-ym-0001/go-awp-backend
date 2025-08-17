package handler

import "github.com/adi-ym-0001/go-awp-backend/internal/usecase"

type Server struct {
	LocationUC *usecase.LocationUsecase
	DrawingUC  *usecase.DrawingUsecase
}
