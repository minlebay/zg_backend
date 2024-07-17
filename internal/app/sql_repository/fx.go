package sql_repository

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"zg_backend/internal/model"
)

type SqlRepository interface {
	Start()
	Stop()
	GetAll(db *gorm.DB) ([]*model.Message, error)
	GetById(db *gorm.DB, uuid string) (*model.Message, error)
	GetDbs() []*gorm.DB
}

func NewModule() fx.Option {

	return fx.Module(
		"repo",
		fx.Provide(
			NewRepositoryConfig,
			fx.Annotate(
				NewMySQLRepository,
				fx.As(new(SqlRepository)),
			),
		),
		fx.Invoke(
			func(lc fx.Lifecycle, r SqlRepository) {
				lc.Append(fx.StartStopHook(r.Start, r.Start))
			},
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("repo")
		}),
	)
}
