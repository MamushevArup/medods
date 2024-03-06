package service

import (
	"github.com/MamushevArup/jwt-auth/internal/repository"
	"github.com/MamushevArup/jwt-auth/internal/service/token"
	"github.com/MamushevArup/jwt-auth/pkg/logger"
)

type Service struct {
	Token token.Generator
}

func NewService(repo *repository.Repo, lg *logger.Logger) *Service {
	return &Service{
		Token: token.NewTokenGenerator(repo, lg),
	}
}
