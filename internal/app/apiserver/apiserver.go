package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/store/sqlstore"
	"github.com/sirupsen/logrus"
)

func Start(config *Config) error {
	db, err := newDB(config.DataBaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.NewStore(db)
	srv := NewServer(store)

	loglevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	srv.logger.SetLevel(loglevel)

	srv.logger.Info("start apiserver")
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseurl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseurl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return  nil, err
	}

	return db, nil
}