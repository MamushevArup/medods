package repository

import (
	"github.com/MamushevArup/jwt-auth/internal/repository/mongo/auth"
	"github.com/MamushevArup/jwt-auth/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Auth auth.UserInserter
}

func NewRepo(lg *logger.Logger, db *mongo.Collection) *Repo {
	return &Repo{
		Auth: auth.NewAuth(lg, db),
	}
}
