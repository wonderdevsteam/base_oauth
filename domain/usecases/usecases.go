package usecases

import (
	"time"

	"go.uber.org/zap"
)

type Usecases struct {
	logger *zap.Logger
	since  time.Time
}

func New(logger *zap.Logger) *Usecases {
	return &Usecases{
		logger: logger.Named("Usecases"),
		since:  time.Now(),
	}
}

func (uc *Usecases) Home() string {
	return "Home working"
}

func (uc *Usecases) GetAppUptime() time.Duration {
	return time.Since(uc.since)
}
