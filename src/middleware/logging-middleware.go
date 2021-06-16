package middleware

import (
	"context"

	"cryptoserver-clean-app/model"
	"cryptoserver-clean-app/service"
	"cryptoserver-clean-app/util"
	"time"

	"github.com/sirupsen/logrus"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(service.Service) service.Service

/*LoggingMiddleware ...*/
//LoggingMiddleware creates MiddleWare
func LoggingMiddleware(logger *logrus.Entry) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   service.Service
	logger *logrus.Entry
}

func (mw loggingMiddleware) GetSymbol(ctx context.Context, symbol string) (model.GetSymbolResponseBody, error) {

	mw.logger = mw.logger.WithFields(logrus.Fields{"api": "/symbol"})
	mw.logger = mw.logger.WithFields(logrus.Fields{"Symbol": symbol})
	mw.logger.WithFields(logrus.Fields{"start": time.Now()}).Info("Started")

	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{"took": time.Since(begin).String()}).Info("Completed")
	}(time.Now())

	ctx = util.WithLogger(ctx, mw.logger)
	return mw.next.GetSymbol(ctx, symbol)
}

