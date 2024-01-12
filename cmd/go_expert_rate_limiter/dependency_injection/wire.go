//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"github.com/aluferraz/go-expert-rate-limiter/internal/infra/persistence/rate_limit"
	"github.com/aluferraz/go-expert-rate-limiter/internal/infra/web/middleware"
	"github.com/aluferraz/go-expert-rate-limiter/internal/value_objects"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
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

func NewRateLimitMiddleware(client *redis.Client, requestLimits value_objects.RequestLimits) *middleware.RateLimiterMiddleware {
	wire.Build(
		rate_limit.NewRateLimitRepositoryRedis,
		middleware.NewRateLimiter,
	)
	return &middleware.RateLimiterMiddleware{}
}
