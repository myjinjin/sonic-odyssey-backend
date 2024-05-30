package tests

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/database"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/logging"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/repository_impls/postgresql"
	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/repository_impls/postgresql/testhelper"
	"github.com/myjinjin/sonic-odyssey-backend/internal/domain/repositories"
	"go.uber.org/zap"
)

var (
	userRepo repositories.UserRepository
	flowRepo repositories.PasswordResetFlowRepository
	testdb   *database.Database
	logger   logging.Logger
)

func init() {
	var err error
	logger, err = logging.NewZapLogger(true)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	err = godotenv.Load("../../../../.env.test")
	if err != nil {
		logger.Fatal("failed to load .env.test file", zap.Error(err))
	}
}

func TestMain(m *testing.M) {
	var err error
	testdb, err = testhelper.SetupTestDB()
	if err != nil {
		logger.Fatal("failed to set up test database", zap.Error(err))
	}
	defer testdb.Close()

	userRepo = postgresql.NewUserRepository(testdb.GetDB())
	flowRepo = postgresql.NewPasswordResetFlowRepository(testdb.GetDB())
	code := m.Run()

	os.Exit(code)
}
