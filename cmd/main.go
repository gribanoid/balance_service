package main

import (
	"context"
	"fmt"
	"github.com/gribanoid/balance_service/config"
	"github.com/gribanoid/balance_service/internal/application/users"
	"github.com/gribanoid/balance_service/internal/repositories/postgres"
	"github.com/gribanoid/balance_service/internal/repositories/user"
	"github.com/jackc/pgx/v4"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

func main() {

	mainCtx := context.Background()
	var c config.Config
	if err := envconfig.Process("crypto-aggregator", &c); err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	// init pg pool connection
	pool, err := postgres.NewPgxPool(mainCtx, c.PostgresHost, pgx.LogLevelDebug)
	if err != nil {
		log.Fatalf("failed to init postgres: %v\n", err)
	}

	if err := pool.Ping(mainCtx); err != nil {
		log.Fatalf("postgres unavailable: %v\n", err)
	}

	// init repo
	userRepo := user.NewUserRepository(pool, time.Second*5)

	// create users service
	userService, err := users.NewUserService(userRepo)
	if err != nil {
		log.Fatalf("failed to init tron service: %v\n", err)
	}
	fmt.Println(userService)
}
