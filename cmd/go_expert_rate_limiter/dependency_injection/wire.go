//go:build wireinject
// +build wireinject

package dependency_injection

import (
	_ "github.com/google/wire"
)
/*
var setSampleRepositoryDependency = wire.NewSet(
	database.SampleRepository,
	wire.Bind(new(entity.SampleRepositoryInterface), new(*database.SampleRepository)),
)

func NewListAllOrdersUseCase(db *sql.DB) *usecase.MyUseCase {
	wire.Build(
		setSampleRepositoryDependency,
		usecase.NewUseCase,
	)
	return &usecase.MyUseCase{}
}
*/