package postgres

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fiber-go-pos-api/internal/app/constant"
	filePkg "github.com/fiber-go-pos-api/internal/pkg/file"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbPG *sqlx.DB

// OpenConnection custom package to open connection
func OpenConnection() error {
	var err error

	// Initialize variable
	dbHost := os.Getenv(constant.DBHostEnvKey)
	dbName := os.Getenv(constant.DBNameEnvKey)
	dbUser := os.Getenv(constant.DBUserEnvKey)
	dbPassword := os.Getenv(constant.DBPasswordEnvKey)
	dbPort, err := strconv.Atoi(os.Getenv(constant.DBPortEnvKey))
	if err != nil {
		return errors.New("port must be number")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	dbPG, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}

	dbPG.SetMaxIdleConns(10)
	dbPG.SetConnMaxLifetime(5 * time.Minute)
	dbPG.SetMaxOpenConns(50)
	return nil
}

// GetPgConn custom package to get postgres connection
func GetPgConn() *sqlx.DB {
	return dbPG
}

// BeginTxx custom package to begin transaction
func BeginTxx(ctx *fiber.Ctx) (*sqlx.Tx, error) {
	return dbPG.BeginTxx(ctx.Context(), nil)
}

// InitializeSchema : Function to set up postgres schema
func InitializeSchema(embedSchemaFiles embed.FS) error {
	contents, err := filePkg.GetAllContentFiles(&embedSchemaFiles, "")
	if err != nil {
		return err
	}

	for _, content := range contents {
		if _, err := dbPG.Exec(content); err != nil {
			return err
		}
	}

	return nil
}
