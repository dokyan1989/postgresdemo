package service

import (
	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type Service struct {
	db     *dbr.Connection
	logger *zap.Logger
}

func NewService(db *dbr.Connection, logger *zap.Logger) *Service {
	return &Service{
		db:     db,
		logger: logger,
	}
}
