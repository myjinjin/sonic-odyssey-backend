package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/auth"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/database"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/email"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/encryption"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/logging"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/repository_impls/postgresql"
	v1 "github.com/myjinjin/sonic-odyssey-backend/internal/controller/http/v1"
	"github.com/myjinjin/sonic-odyssey-backend/internal/usecase"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logging.Log().Fatal("failed to load .env file: %v", zap.Error(err))
	}
	db, err := database.NewDB(
		database.WithHost(os.Getenv("DB_HOST")),
		database.WithPort(os.Getenv("DB_PORT")),
		database.WithUsername(os.Getenv("DB_USERNAME")),
		database.WithPassword(os.Getenv("DB_PASSWORD")),
		database.WithDBName(os.Getenv("DB_NAME")),
	)
	if err != nil {
		logging.Log().Fatal("faied to connect to the database: %v", zap.Error(err))
	}
	defer db.Close()

	if err := db.InitMigrator("migrations"); err != nil {
		logging.Log().Fatal("faied to initialize migrator: %v", zap.Error(err))
	}

	if err := db.MigrateUp(); err != nil {
		logging.Log().Fatal("faied to run database migrations: %v", zap.Error(err))
	}

	encryptor, err := encryption.NewAESEncryptor(os.Getenv("DB_ENCRYPTION_KEY"))
	if err != nil {
		logging.Log().Fatal("failed to create encryptor: ", zap.Error(err))
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFromAddress := os.Getenv("SMTP_FROM_ADDRESS")

	emailSender, err := email.NewSMTPEmailSender(smtpHost, smtpPort, smtpUsername, smtpPassword, smtpFromAddress)
	if err != nil {
		logging.Log().Fatal("failed to create email sender: ", zap.Error(err))
	}

	userRepo := postgresql.NewUserRepository(db.GetDB())
	passwordResetRepo := postgresql.NewPasswordResetFlowRepository(db.GetDB())
	userUsecase := usecase.NewUserUsecase(userRepo, passwordResetRepo, encryptor, emailSender)
	userJwt := auth.NewUserJWT(userRepo)

	jwtAuth, err := auth.NewJWTMiddleware(
		auth.WithKey([]byte(os.Getenv("JWT_SECRET_KEY"))),
		auth.WithPayloadFunc(userJwt.PayloadFunc),
		auth.WithIdentityHandler(userJwt.IdentityHandler),
		auth.WithAuthenticator(userJwt.Authenticator),
		auth.WithAuthorizator(userJwt.Authorizator),
		auth.WithUnauthorized(userJwt.Unauthorized),
		auth.WithLoginResponse(userJwt.LoginResponse),
	)
	if err != nil {
		logging.Log().Fatal("failed to create jwt auth middleware: ", zap.Error(err))
	}
	router := v1.SetupRouter(userUsecase, jwtAuth)

	err = router.Run(":8081")
	if err != nil {
		logging.Log().Fatal("failed to start server: %v", zap.Error(err))
	}
}
