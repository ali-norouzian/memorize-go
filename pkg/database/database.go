package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memorize/config"
	"strings"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(lc fx.Lifecycle, config *config.Config, listOfDbModels []any) (*gorm.DB, error) {
	postgreDbConfig := config.GetDbSetting()
	err := createDatabaseIfNotExists(config)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(postgreDbConfig.GetDbConnectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Database connected")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Database disconnected")
			dbSQL, err := db.DB()
			if err != nil {
				return err
			}
			return dbSQL.Close()
		},
	})

	// Migrate the schema
	err = db.AutoMigrate(listOfDbModels...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createDatabaseIfNotExists(config *config.Config) error {
	// Connect to PostgreSQL without specifying a database
	postgreDbConfig := config.GetDbSetting()
	connStr := strings.Replace(postgreDbConfig.GetDbConnectionString(),
		postgreDbConfig.DbName,
		"postgres",
		1)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// postgreDbConfig.DbName = strings.ToLower(postgreDbConfig.DbName)
	// Check if the database exists
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", postgreDbConfig.DbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return err
	}

	// Create the database if it doesn't exist
	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\";", postgreDbConfig.DbName))
		if err != nil {
			return err
		}
		log.Printf("Database %s created successfully.", postgreDbConfig.DbName)
	} else {
		log.Printf("Database %s already exists.", postgreDbConfig.DbName)
	}

	return nil
}
