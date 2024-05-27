package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/database"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/encryption"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/hash"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/repository_impls/postgresql"
	v1 "github.com/myjinjin/sonic-odyssey-backend/internal/controller/http/v1"
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

	if err := db.InitMigrator("migrations"); err != nil {
		log.Fatalf("faied to initialize migrator: %v", err)
	}

	if err := db.MigrateUp(); err != nil {
		log.Fatalf("faied to run database migrations: %v", err)
	}

	encryptor, err := encryption.NewAESEncryptor("")
	if err != nil {
		log.Fatal("failed to create encryptor:", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFromAddress := os.Getenv("SMTP_FROM_ADDRESS")

	emailSender, err := email.NewSMTPEmailSender(smtpHost, smtpPort, smtpUsername, smtpPassword, smtpFromAddress)
	if err != nil {
		log.Fatal("failed to create email sender:", err)
	}

	userRepo := postgresql.NewUserRepository(db.GetDB())
	userUsecase := usecase.NewUserUsecase(userRepo, hash.BCryptPasswordHasher(), hash.SHA256EmailHasher(), encryptor, emailSender)

	router := v1.SetupRouter(userUsecase)

	err = router.Run(":8081")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
