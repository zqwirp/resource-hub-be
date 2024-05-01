package config

import "os"

type Config struct {
	AppName string

	Env string

	LogPath     string
	LogFileName string

	HTTPPort string

	DBConnection string
	DBUser       string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	DBSSLMode    string
}

func NewConfig() *Config {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "goblok"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		logPath = "./"
	}

	logFileName := os.Getenv("LOG_PATH")
	if logFileName == "" {
		logFileName = "logs"
	}

	dbSslMode := os.Getenv("DB_SSLMODE")
	if dbSslMode == "" {
		dbSslMode = "disable"
	}

	c := Config{
		AppName: appName,

		Env: env,

		LogPath:     logPath,
		LogFileName: logFileName,

		HTTPPort: os.Getenv("HTTP_PORT"),

		DBConnection: os.Getenv("DB_CONNECTION"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    dbSslMode,
	}

	return &c
}
