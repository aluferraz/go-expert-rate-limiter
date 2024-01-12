package go_expert_rate_limiter

import (
	"github.com/aluferraz/go-expert-rate-limiter/cmd/go_expert_rate_limiter/dependency_injection"
	"github.com/aluferraz/go-expert-rate-limiter/configs"
	"github.com/aluferraz/go-expert-rate-limiter/internal/infra/web/web"
	"github.com/aluferraz/go-expert-rate-limiter/internal/value_objects"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func handleErr(err error) {
	if err != nil {
		log.Error().Err(err).Msg("")
		panic(err)
	}
}

func Bootstap() {

	workdir, err := os.Getwd()
	handleErr(err)
	appConfig, err := configs.LoadConfig(workdir)
	if err != nil {
		panic(err)
	}

	webserver := web.NewWebServer(appConfig.WebserverPort)
	webserver.AddHandler("/", http.MethodGet, func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Up and Running"))
		handleErr(err)
	})
	rdb := redis.NewClient(&redis.Options{
		Addr:     appConfig.RedisURI,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	limits := value_objects.NewRequestLimit(appConfig.IPThrottling, appConfig.APIThrottling, appConfig.Expiration)
	middlewareRequestLimiter := dependency_injection.NewRateLimitMiddleware(rdb, limits)

	webserver.AddMiddleware(middlewareRequestLimiter.RateLimiter)
	log.Info().Msgf("Listening on: %s", appConfig.WebserverPort)
	err = webserver.Start()
	if err != nil {
		panic(err)
	}

}
