package types

import (
	"context"

	"go.uber.org/zap"
)

type Context struct {
	Ctx    context.Context
	Logger *zap.Logger
}
