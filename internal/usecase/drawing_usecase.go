package usecase

import (
	"context"

	"github.com/adi-ym-0001/go-awp-backend/internal/model"
	"github.com/adi-ym-0001/go-awp-backend/internal/repository"
)

type DrawingUsecase struct {
	Repo *repository.DrawingRepository // DBアクセスの依存を注入（DI）
}

func (uc *DrawingUsecase) GetAll(ctx context.Context) ([]model.Drawing, error) {
	// DBから図面情報を取得（Entity）
	raw, err := uc.Repo.FindAll(ctx)
	if err != nil {
		return nil, err // エラーはそのまま返却（ハンドラで処理）
	}

	var result []model.Drawing
	for _, d := range raw {
		// Entity → DTOへの変換
		result = append(result, model.Drawing{
			ID:      d.ID,
			Name:    d.Name,
			Type:    d.Type,
			Version: d.Version,
			Status:  d.Status,
		})
	}
	return result, nil
}
