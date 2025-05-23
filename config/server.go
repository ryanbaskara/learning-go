package config

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ryanbaskara/learning-go/handler"
	"github.com/ryanbaskara/learning-go/repository"
	"github.com/ryanbaskara/learning-go/usecase"
)

type Server struct {
	HttpServer *http.Server
}

func NewServer() (*Server, error) {
	cfg, err := loadServerConfig()
	if err != nil {
		return nil, err
	}

	mysqlDB := newMysqlDatabase(&cfg.DatabaseConfig)

	repo := repository.NewRepository(mysqlDB)
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	httpServer := &http.Server{
		Addr:         cfg.ServerHost,
		Handler:      handler.RegisterHandler(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return &Server{
		HttpServer: httpServer,
	}, nil
}

func newMysqlDatabase(cfg *DatabaseConfig) *sqlx.DB {
	var sqlxDB *sqlx.DB
	var err error

	sqlxDB, err = sqlx.Connect("mysql", cfg.databaseSourceName())

	if err != nil {
		panic(err)
	}

	sqlxDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlxDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlxDB.SetConnMaxIdleTime(cfg.MaxIdleTime)
	sqlxDB.SetConnMaxLifetime(cfg.MaxLifetime)

	return sqlxDB
}
