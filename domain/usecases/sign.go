package usecases

import (
	"context"

	"go.uber.org/zap"
)

type Sign struct {
	logger *zap.Logger
}

func NewSign(logger *zap.Logger) *Sign {
	return &Sign{
		logger: logger.Named("Usecases"),
	}
}

func (sign *Sign) In(ctx context.Context)  {

}

func (sign *Sign) Out(ctx context.Context) {}