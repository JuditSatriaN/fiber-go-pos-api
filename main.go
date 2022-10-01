package main

import (
	"embed"
	"log"
	"time"

	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	apiRouter "github.com/fiber-go-pos-api/internal/app/router"
	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
	jwtPkg "github.com/fiber-go-pos-api/internal/pkg/jwt"
	serverPkg "github.com/fiber-go-pos-api/internal/pkg/server"
	goccyJson "github.com/goccy/go-json"
)

// Embed a schema directory
//go:embed schema/postgres/*
var embedSchemaFiles embed.FS

// Embed a private pem files
//go:embed schema/pem/private.pem
var embedPrivatePEMFile []byte

func main() {
	// Initialization App Config
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		AppName:      constant.AppName,
		JSONEncoder:  goccyJson.Marshal,
		JSONDecoder:  goccyJson.Unmarshal,
	})

	// Load Environment
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment, err:%v", err)
	}

	// Setting basic configuration
	app.Use(logger.New(), recover.New())

	// Setting CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Setting JWT RS256
	if err := jwtPkg.GenerateJWT(embedPrivatePEMFile); err != nil {
		log.Fatalf("rsa.GenerateKey, err:%v", err)
	}

	// Open Postgres Connection
	if err := postgresPkg.OpenConnection(); err != nil {
		log.Fatalf("failed to open connection postgre, err:%v", err)
	}

	// Initialize schema if not exists
	if err := postgresPkg.InitializeSchema(embedSchemaFiles); err != nil {
		log.Fatalf("failed to initialize schema postgre, err:%v", err)
	}

	// API Handler
	apiRouter.BuildAPIRouter(app)

	if err := app.Listen(serverPkg.GetAppPort()); err != nil {
		log.Fatalf("failed to run app, err:%v", err)
	}
}
