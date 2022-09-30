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
	"strconv"

	"log"
	"net/http"
	"time"
)

func main() {
	mainCtx := context.Background()
	var c config.Config
	if err := envconfig.Process("balance-service", &c); err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}
	fmt.Println(c.PostgresHost)
	// init pg pool connection
	pool, err := postgres.NewPgxPool(mainCtx, c.PostgresHost, pgx.LogLevelDebug)
	if err != nil {
		log.Fatalf("failed to init postgres: %v\n", err)
	}
	time.Sleep(5 * time.Second)
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
	fmt.Println("ready to accept requests")

	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from docker")
	})
	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		err = userService.CreateUser(mainCtx, userID)
	})
	http.HandleFunc("/user/balance", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		balance, err := userService.GetBalance(mainCtx, userID)
		if err != nil {
			// TODO
		}
		fmt.Fprintf(w, "баланс пользователя %v", balance)
	})
	http.HandleFunc("/user/deposit", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		amount := r.URL.Query().Get("amount")
		a, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			// TODO
		}
		err = userService.Deposit(mainCtx, userID, int(a))
		if err != nil {
			// TODO
		}
		fmt.Fprintf(w, "success deposit")
	})
	http.HandleFunc("/user/withdrawal", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		amount := r.URL.Query().Get("amount")
		a, err := strconv.ParseInt(amount, 10, 64)
		if err != nil {
			// TODO
		}
		err = userService.Deposit(mainCtx, userID, int(a))
		if err != nil {
			// TODO
		}
		fmt.Fprintf(w, "success deposit")
	})
	http.HandleFunc("/user/send", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		userService.CreateUser(mainCtx, userID)
	})
	http.ListenAndServe(":8080", nil)
}
