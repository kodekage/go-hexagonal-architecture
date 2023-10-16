package utils

import (
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kodekage/banking/internal/logger"
)

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func SqlClient() *sqlx.DB {
	sqlClient, err := sqlx.Open("mysql", "root:rootpw@/banking")
	if err != nil {
		logger.Info(err.Error())
		panic(err)
	}
	// See "Important settings" section.
	sqlClient.SetConnMaxLifetime(time.Minute * 3)
	sqlClient.SetMaxOpenConns(10)
	sqlClient.SetMaxIdleConns(10)

	return sqlClient
}
