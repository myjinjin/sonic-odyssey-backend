package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/database"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/password"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/repository_impls/postgresql"
	"github.com/myjinjin/sonic-odyssey-backend/internal/controller/http"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
	db, err := database.NewDB(
		database.WithHost(os.Getenv("DB_HOST")),
		database.WithPort(os.Getenv("DB_PORT")),
		database.WithUsername(os.Getenv("DB_USERNAME")),
		database.WithPassword(os.Getenv("DB_PASSWORD")),
		database.WithDBName(os.Getenv("DB_NAME")),
	)
	if err != nil {
		log.Fatalf("faied to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.InitMigrator(); err != nil {
		log.Fatalf("faied to initialize migrator: %v", err)
	}

	if err := db.MigrateUp(); err != nil {
		log.Fatalf("faied to run database migrations: %v", err)
	}

	userRepo := postgresql.NewUserRepository(db.GetDB(), os.Getenv("DB_ENCRYPTION_KEY"))
	userUsecase := usecase.NewUserUsecase(userRepo, password.BCryptPasswordHasher())

	router := http.SetupRouter(userUsecase)

	err = router.Run(":8081")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
