package repository

import (
	"context"

	"github.com/adi-ym-0001/go-awp-backend/internal/model" // DBモデル（Entity）
	"gorm.io/gorm"
)

// LocationRepository はロケーション情報のDBアクセス責任を持つ
type LocationRepository struct {
	DB *gorm.DB // GORMのDBインスタンス（DI）
}

// FindAllWithDrawings はロケーションと関連する図面情報を取得する
func (r *LocationRepository) FindAllWithDrawings(ctx context.Context) ([]model.Location, error) {
	var locations []model.Location
	// GORMのPreloadを使って Drawings を同時取得（N+1問題の回避）
	err := r.DB.WithContext(ctx).Preload("Drawings").Find(&locations).Error
	return locations, err // エラーはユースケース層で処理
}

// Update はロケーションの図面情報を更新する
func (r *LocationRepository) UpdateDrawings(ctx context.Context, locationId string, drawings []model.Drawing) error {
	var location model.Location
	if err := r.DB.WithContext(ctx).First(&location, "id = ?", locationId).Error; err != nil {
		return err
	}
	return r.DB.Model(&location).Association("Drawings").Replace(drawings)
}
