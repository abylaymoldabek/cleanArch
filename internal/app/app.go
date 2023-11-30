package app

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gitlab.com/sunrise-dev/utest/utest-auth-app/config"
	"gitlab.com/sunrise-dev/utest/utest-auth-app/internal/repositories"
	authenticationUsecase "gitlab.com/sunrise-dev/utest/utest-auth-app/internal/usecase/authentication"
	"gitlab.com/sunrise-dev/utest/utest-auth-app/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Подключение к БД
	// pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	// }
	// defer pg.Close()

	// Подключение к БД
	db, err := sqlx.Connect("mysql", "sundev:sundev@(localhost:3306)/utest")
	if err != nil {
		l.Fatal(fmt.Errorf("DB connect error: %w", err))
	}
	if err := db.Ping(); err != nil {
		l.Fatal(fmt.Errorf("DB ping error: %w", err))
	}
	log.Println("DB connected")

	// Инициализация UseCase
	authenticationUseCase := authenticationUsecase.New(
		repositories.NewUserRepoSQL(db),
	)

	r := mux.NewRouter()
	r.HandleFunc("/users", authenticationUseCase.GetAll()).Methods("GET")
	http.ListenAndServe(":80", r)
}
