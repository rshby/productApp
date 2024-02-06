package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"productApp/app/config"
	"time"
)

func ConnectDB(cfg config.IConfig, log *logrus.Logger) *sql.DB {
	config := cfg.GetConfig()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("cant connect to DB : %v", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	log.Info("success connect to database")

	return db
}
