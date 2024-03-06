package token

import (
	"context"
	"errors"
	"github.com/MamushevArup/jwt-auth/internal/models"
	"github.com/MamushevArup/jwt-auth/internal/repository"
	"github.com/MamushevArup/jwt-auth/pkg/logger"
	"github.com/MamushevArup/jwt-auth/utils"
)

// Generator injected to the service manager
type Generator interface {
	Generate(ctx context.Context, guid string) (*models.GenerateResponse, error)
	UpdateToken(ctx context.Context, guid, refresh string) (*models.GenerateResponse, error)
}

// Gen contains repo interface implementation and log errors
type Gen struct {
	repo *repository.Repo
	lg   *logger.Logger
}

func NewTokenGenerator(repo *repository.Repo, lg *logger.Logger) Generator {
	return &Gen{repo: repo, lg: lg}
}

func (g *Gen) UpdateToken(ctx context.Context, guid, refresh string) (*models.GenerateResponse, error) {

	ex, err := g.repo.Auth.UniqueUser(ctx, guid)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, errors.New("user doesn't exist")
	}

	tokenHash, err := g.repo.Auth.UserToken(ctx, guid)
	if err != nil {
		g.lg.Errorf("can't fetch token %v", err)
		return nil, err
	}

	ok, err := utils.CheckToken(refresh, tokenHash)
	if err != nil {
		g.lg.Errorf("check failed due to %v", err)
		return nil, err
	}

	if !ok {
		return nil, errors.New("refresh token is corrupted")
	}

	access, refreshToken, hashedRefresh, err := g.generateTokens(guid)
	if err != nil {
		g.lg.Errorf("fail with tokens %v", err)
		return nil, err
	}

	err = g.repo.Auth.UpdateRefreshToken(ctx, guid, hashedRefresh)
	if err != nil {
		g.lg.Errorf("fail update %v", err)
		return nil, err
	}
	r := &models.GenerateResponse{
		AccessToken:  access,
		RefreshToken: refreshToken,
	}
	return r, nil
}

// Generate access and refresh token and call method to the repo to store a refresh and guid of user
// return access and refresh into cookie http-only
func (g *Gen) Generate(ctx context.Context, guid string) (*models.GenerateResponse, error) {
	ex, err := g.repo.Auth.UniqueUser(ctx, guid)
	if err != nil {
		g.lg.Errorf("error uniqueness check %v", err)
		return nil, err
	}
	if ex {
		return nil, errors.New("user exist")
	}

	accessToken, refresh, hash, err := g.generateTokens(guid)
	if err != nil {
		g.lg.Errorf("fail to generate tokens %v", err)
		return nil, err
	}

	if err = g.repo.Auth.InsertUser(ctx, guid, hash); err != nil {
		g.lg.Errorf("unable insert user %v", err)
		return nil, err
	}

	response := &models.GenerateResponse{
		AccessToken:  accessToken,
		RefreshToken: refresh,
	}

	return response, nil
}

// generateTokens return all about tokens
func (g *Gen) generateTokens(guid string) (string, string, string, error) {
	accessToken, err := utils.GenerateAccess(guid)
	if err != nil {
		g.lg.Errorf("due to %v", err)
		return "", "", "", err
	}

	refresh := utils.GenerateRefresh()

	hash, err := utils.HashToken(refresh)
	if err != nil {
		g.lg.Errorf("hash failed %v", err)
		return "", "", "", err
	}
	return accessToken, refresh, hash, nil
}
