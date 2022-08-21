package service

import (
	"github.com/dokyan1989/postgresdemo/server/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	repo   repository.Repository
	logger *zap.Logger
}

func NewService(repo repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
