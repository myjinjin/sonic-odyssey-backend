package database

import (
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	postgresmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db       *gorm.DB
	migrator *migrate.Migrate
}

func NewDB(opts ...Option) (*Database, error) {
	cfg := &DBConfig{
		SSLMode:         "disable",
		MaxOpenConns:    100,
		MaxIdleConns:    10,
		ConnMaxLifetime: 30 * time.Minute,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.Host == "" {
		return nil, fmt.Errorf("database host is required")
	}
	if cfg.Port == "" {
		return nil, fmt.Errorf("database port is required")
	}
	if cfg.Username == "" {
		return nil, fmt.Errorf("database username is required")
	}
	if cfg.Password == "" {
		return nil, fmt.Errorf("database password is required")
	}
	if cfg.DBName == "" {
		return nil, fmt.Errorf("database name is required")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &Database{db: db}, nil
}

func (d *Database) InitMigrator() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	driver, err := postgresmigrate.WithInstance(sqlDB, &postgresmigrate.Config{})
	if err != nil {
		return err
	}
	d.migrator, err = migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) MigrateUp() error {
	if err := d.migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (d *Database) MigrateDown() error {
	if err := d.migrator.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func (d *Database) Close() error {
	if d.migrator != nil {
		d.migrator.Close()
	}

	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type Option func(*DBConfig)

func WithHost(host string) Option {
	return func(cfg *DBConfig) {
		cfg.Host = host
	}
}

func WithPort(port string) Option {
	return func(cfg *DBConfig) {
		cfg.Port = port
	}
}

func WithUsername(username string) Option {
	return func(cfg *DBConfig) {
		cfg.Username = username
	}
}

func WithPassword(password string) Option {
	return func(cfg *DBConfig) {
		cfg.Password = password
	}
}

func WithDBName(dbName string) Option {
	return func(cfg *DBConfig) {
		cfg.DBName = dbName
	}
}

func WithSSLMode(sslMode string) Option {
	return func(cfg *DBConfig) {
		cfg.SSLMode = sslMode
	}
}

func WithMaxOpenConns(maxOpenConns int) Option {
	return func(cfg *DBConfig) {
		cfg.MaxOpenConns = maxOpenConns
	}
}

func WithMaxIdleConns(maxIdleConns int) Option {
	return func(cfg *DBConfig) {
		cfg.MaxIdleConns = maxIdleConns
	}
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) Option {
	return func(cfg *DBConfig) {
		cfg.ConnMaxLifetime = connMaxLifetime
	}
}
