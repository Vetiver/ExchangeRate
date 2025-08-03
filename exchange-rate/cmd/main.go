package main

import (
	"exchangeRate/config"
	"exchangeRate/internal/domain"
	"exchangeRate/pkg"
	"exchangeRate/pkg/ext/cbr"
	httpInfra "exchangeRate/transport"
	httpHandler "exchangeRate/transport/handlers"
	"net/http"
	_ "net/http/pprof"
	"sync"

	"go.uber.org/zap"
)

func main() {
	logger := pkg.CreateLogger()
	defer logger.Sync()
	config := config.NewConfig()
	cache:= cbr.NewCache()
	wg := sync.WaitGroup{}
	wg.Add(2)

	domainService := domain.NewExchangeRateService(logger, config, cache)
	httpHandlers := httpHandler.NewBaseHandler(logger, domainService, config)
	httpServer := httpInfra.NewHttpServer(logger, config.HTTPAddr)

	go func() {
		defer wg.Done()
		if config.ProfilerAddr != "" {
			err := http.ListenAndServe(config.ProfilerAddr, nil)
			if err != nil {
				logger.Error("fail to start pprof", zap.Error(err))
			}
		}
	}()

	go func() {
		defer wg.Done()
		httpServer.StartHTTPServer(httpHandlers)
		logger.Error("HTTP server down")
	}()

	wg.Wait()
}
