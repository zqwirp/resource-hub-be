package config

import "os"

type Config struct {
	AppName string

	Env string

	LogPath string
	LogFile string

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

	logFile := os.Getenv("LOG_PATH")
	if logFile == "" {
		logFile = "logs"
	}

	dbSslMode := os.Getenv("DB_SSLMODE")
	if dbSslMode == "" {
		dbSslMode = "disable"
	}

	c := Config{
		AppName: appName,

		Env: env,

		LogPath: logPath,
		LogFile: logFile,

		HTTPPort: os.Getenv("HTTP_PORT"),
	}

	return &c
}
