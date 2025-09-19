package service

import (
	"context"
	"mall/internal/logger"
	"math/rand"
	"time"
)

func HealthCheck(ctx context.Context) (resp interface{}, err error) {
	r := rand.Intn(1000)

	time.Sleep(time.Microsecond * time.Duration(r))

	logger.WithContext(ctx).Infof("HealthCheck success, sleep %d microseconds", r)

	return r, nil
}

func HealthCheckV1(ctx context.Context) (resp interface{}, err error) {
	r := rand.Intn(1000)

	time.Sleep(time.Microsecond * time.Duration(r))

	logger.WithGoID().Infof("HealthCheckV1 success, sleep %d microseconds", r)

	return r, nil
}
