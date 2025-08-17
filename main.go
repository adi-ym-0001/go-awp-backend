package main

import (
	"log"

	api "github.com/adi-ym-0001/go-awp-backend/internal/gen"    // OpenAPIから自動生成されたDTOとルーティング定義
	"github.com/adi-ym-0001/go-awp-backend/internal/handler"    // HTTPハンドラ（APIエンドポイントの実装）
	"github.com/adi-ym-0001/go-awp-backend/internal/repository" // DBアクセス層（永続化責任）
	"github.com/adi-ym-0001/go-awp-backend/internal/usecase"    // ユースケース層（ビジネスロジック）
	"github.com/labstack/echo/v4"                               // Echoフレームワーク（Webサーバ）
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres" // PostgreSQLドライバ
	"gorm.io/gorm"            // GORM ORMライブラリ
)

func main() {
	// DB接続の初期化（DSNは未設定なので、環境変数や設定ファイルから取得すべき）
	db, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		log.Fatal(err) // 接続失敗時はログ出力して終了（運用上の心理的安全性）
	}

	// Repositoryの初期化（永続化責任を持つ）
	locationRepo := &repository.LocationRepository{DB: db}
	drawingRepo := &repository.DrawingRepository{DB: db}

	// Usecaseの初期化（ビジネスロジックを持つ）
	locationUC := &usecase.LocationUsecase{Repo: locationRepo}
	drawingUC := &usecase.DrawingUsecase{Repo: drawingRepo}

	// Handlerの初期化（HTTPエンドポイントの実装）
	server := &handler.Server{
		LocationUC: locationUC,
		DrawingUC:  drawingUC,
	}

	// Echoのインスタンス作成とルーティング登録
	e := echo.New()

	// ✅ CORSミドルウェアの追加（ここがポイント）
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	api.RegisterHandlers(e, server) // OpenAPI定義に基づくルーティング登録

	// サーバ起動（ポート8080）
	e.Logger.Fatal(e.Start(":8080"))
}
