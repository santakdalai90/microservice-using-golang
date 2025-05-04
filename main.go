package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "payment",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		level.Fatal(logger).Log("exit", err)
	}
	defer db.Close()

}
