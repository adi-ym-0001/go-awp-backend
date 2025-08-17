package repository

import (
	"context"

	"github.com/adi-ym-0001/go-awp-backend/internal/model"
	"gorm.io/gorm"
)

type DrawingRepository struct {
	DB *gorm.DB // GORMのDBインスタンス（DI）
}

func (r *DrawingRepository) FindAll(ctx context.Context) ([]model.Drawing, error) {
	var drawings []model.Drawing
	err := r.DB.WithContext(ctx).Find(&drawings).Error
	return drawings, err // エラーはユースケース層で処理
}
