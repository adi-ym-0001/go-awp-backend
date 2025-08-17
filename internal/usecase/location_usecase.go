package usecase

import (
	"context"

	// OpenAPI定義に基づくDTO
	"github.com/adi-ym-0001/go-awp-backend/internal/model"
	"github.com/adi-ym-0001/go-awp-backend/internal/repository" // DBアクセス層
)

// LocationUsecase はロケーション関連のビジネスロジックを提供する
type LocationUsecase struct {
	Repo *repository.LocationRepository // DBアクセスの依存を注入（DI）
}

// GetAll はロケーションと図面情報を取得し、DTOに変換して返す
func (uc *LocationUsecase) GetAll(ctx context.Context) ([]model.LocationWithDrawings, error) {
	// DBからロケーション＋図面情報を取得（Entity）
	rawLocations, err := uc.Repo.FindAllWithDrawings(ctx)
	if err != nil {
		return nil, err // エラーはそのまま返却（ハンドラで処理）
	}

	var result []model.LocationWithDrawings
	for _, loc := range rawLocations {
		var drawings []model.Drawing
		for _, d := range loc.Drawings {
			// Entity → DTOへの変換（ポインタ型に合わせて & を使用）
			drawings = append(drawings, model.Drawing{
				Id:      d.Id,
				Name:    d.Name,
				Type:    d.Type,
				Version: d.Version,
				Status:  d.Status,
			})
		}
		// ロケーション情報もDTOに変換
		// &はアドレスを返してる
		result = append(result, model.LocationWithDrawings{
			Id:       &loc.Id,
			Name:     &loc.Name,
			Floor:    &loc.Floor,
			Area:     &loc.Area,
			Drawings: &drawings, // ポインタスライスとして渡す
		})
	}
	return result, nil
}

// ポインタ型をそのままコピーする関数
func copyPtr(s *string) *string {
	if s == nil {
		return nil
	}
	str := *s
	return &str
}

func (uc *LocationUsecase) UpdateDrawings(ctx context.Context, locationId string, drawings []model.Drawing) error {
	var entities []model.Drawing
	for _, d := range drawings {
		entities = append(entities, model.Drawing{
			Id:      copyPtr(d.Id), // ポインタから値を取得
			Name:    copyPtr(d.Name),
			Type:    copyPtr(d.Type),
			Version: copyPtr(d.Version),
			Status:  copyPtr(d.Status),
		})
	}
	return uc.Repo.UpdateDrawings(ctx, locationId, entities)
}
