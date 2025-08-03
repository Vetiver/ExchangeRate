package domain

import (
	"exchangeRate/config"
	"exchangeRate/constants"
	"exchangeRate/pkg/ext/cbr"
	"fmt"

	"go.uber.org/zap"
)

type DomainService struct {
	logger      *zap.Logger
	config      *config.Config
	RecordCache *cbr.Cache
}

func NewExchangeRateService(logger *zap.Logger, config *config.Config, cache *cbr.Cache) *DomainService {
	return &DomainService{
		logger:      logger,
		config:      config,
		RecordCache: cache,
	}
}

func (r *DomainService) ExchangeRate(date string, val string, config *config.Config) (Record, error) {
	cbrADDR := config.BaseUrlCBR
	records, ok := r.RecordCache.Get(date, val)
	if ok {
		return Record{Date: records[0].Date, Nominal: records[0].Nominal, Value: records[0].Value, VunitRate: records[0].VunitRate}, nil
	}

	newRec, err := cbr.GetCurrentCurrencyDynamics(date, val, cbrADDR)
	if err != nil {
		r.logger.Error("error GetCurrentCurrencyDynamics", zap.Error(err))
		return Record{}, err
	}

	r.RecordCache.Set(date, val, newRec)

	if len(newRec) == 0 {
		return Record{}, fmt.Errorf(constants.EmptyDynamics)
	}
	domainRecord := Record{Date: newRec[0].Date, Nominal: newRec[0].Nominal, Value: newRec[0].Value, VunitRate: newRec[0].VunitRate}

	return domainRecord, nil
}
