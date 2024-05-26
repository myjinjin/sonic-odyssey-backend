package testhelper

import (
	"fmt"
	"os"

	"github.com/myjinjin/sonic-odyssey-backend/infrastructure/database"
)

func SetupTestDB() (*database.Database, error) {
	testDB, err := database.NewDB(
		database.WithHost(os.Getenv("DB_HOST")),
		database.WithPort(os.Getenv("DB_PORT")),
		database.WithUsername(os.Getenv("DB_USERNAME")),
		database.WithPassword(os.Getenv("DB_PASSWORD")),
		database.WithDBName(os.Getenv("DB_NAME")),
	)
	testDB.GetDB().Debug()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the test database: %v", err)
	}

	if err := testDB.InitMigrator("../../../../migrations"); err != nil {
		return nil, fmt.Errorf("failed to initialize migrator for test database: %v", err)
	}

	if err := testDB.MigrateUp(); err != nil {
		return nil, fmt.Errorf("failed to run database migrations for test database: %v", err)
	}

	return testDB, nil
}
