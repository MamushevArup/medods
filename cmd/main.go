package main

import (
	"context"
	"github.com/MamushevArup/jwt-auth/internal/config"
	"github.com/MamushevArup/jwt-auth/internal/handler"
	"github.com/MamushevArup/jwt-auth/internal/repository"
	"github.com/MamushevArup/jwt-auth/internal/service"
	"github.com/MamushevArup/jwt-auth/pkg/client/mongodb"
	"github.com/MamushevArup/jwt-auth/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// read environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init logger
	lg := logger.NewLogger()

	// init db connector
	mg := cfg.MongoDB
	db, err := mongodb.NewClient(context.TODO(), mg.Port, mg.Host, mg.Database, mg.Collection)
	if err != nil {
		log.Fatal(err)
	}
	_ = lg

	// init repository
	repo := repository.NewRepo(lg, db)

	// init services
	srv := service.NewService(repo, lg)

	// init handler
	hdl := handler.NewHandler(srv)

	go func() {
		if err = http.ListenAndServe(":"+cfg.Server.Port, hdl.InitRoute()); err != nil {
			lg.Fatalf("unable to start server %v", err)
		}
	}()

	select {}
	// two route we pass refresh token from the cookie using http only
	// /generate-token/:guid
	// /refresh/:guid/
}
