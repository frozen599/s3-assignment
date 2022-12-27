package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

// Config struct for describe configuration of the app.
type Config struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	DBHost       string
	DBPort       int
	DBName       string
	DBUser       string
	DBPassword   string
}

var (
	once       sync.Once // create sync.Once primitive
	instance   *Config   // create nil Config struct
	dbInstance *pg.DB
)

// NewConfig function to prepare config variables from .env file and return config.
func NewConfig() *Config {
	// Configuring config one time.
	once.Do(func() {
		curDir, err := os.Getwd()
		if err != nil {
			panic("cannot load current directory")
		}

		err = godotenv.Load(curDir + "/.env")
		if err != nil {
			panic(err)
		}
		// Server host (should be string):
		host := os.Getenv("SERVER_HOST")
		// Server port (should be int):
		port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
		if err != nil {
			panic("wrong server port (check your .env)")
		}
		// Server read timeout (should be int):
		readTimeout, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
		if err != nil {
			panic("wrong server read timeout (check your .env)")
		}
		// Server write timeout (should be int):
		writeTimeout, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
		if err != nil {
			panic("wrong server write timeout (check your .env)")
		}
		// Server idle timeout (should be int):
		idleTimeout, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))
		if err != nil {
			panic("wrong server idle timeout (check your .env)")
		}

		dbHost := os.Getenv("DB_HOST")

		dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			panic("cannot convert db port")
		}

		dbName := os.Getenv("DB_NAME")

		dbUser := os.Getenv("DB_USER")

		dbPassword := os.Getenv("DB_PASSWORD")

		// Set all variables to the config instance.
		instance = &Config{
			Addr:         fmt.Sprintf("%s:%d", host, port),
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
			IdleTimeout:  time.Duration(idleTimeout) * time.Second,
			DBHost:       dbHost,
			DBPort:       dbPort,
			DBName:       dbName,
			DBUser:       dbUser,
			DBPassword:   dbPassword,
		}

	})
	// Return configured config instance.
	return instance
}

func InitDB(cfg *Config) *pg.DB {
	dbInstance = pg.Connect(&pg.Options{
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Database: cfg.DBName,
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
	})
	return dbInstance
}

func GetDB() *pg.DB {
	if dbInstance == nil {
		panic("must call InitDB first")
	}
	return dbInstance
}
