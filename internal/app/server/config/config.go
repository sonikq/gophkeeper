package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RunAddress string
	Postgres   PostgresConfig
	CtxTimeout time.Duration
}

type PostgresConfig struct {
	URI     string
	Workers int
}

const (
	defaultRunAddress          = "localhost:8080"
	defaultCtxTimeOut          = 50000
	defaultMode                = "debug" // (release or debug)
	defaultPostgresURI         = "postgres://postgres:postgres@127.0.0.1:5432/postgres"
	defaultPostgresPoolWorkers = 30
)

func Load() (Config, error) {
	mode := flag.String("mode", defaultMode, "mode defines which env file to use")
	runAddress := flag.String("run_address", defaultRunAddress, "run address defines on what port and host the server will be started")
	postgresURI := flag.String("postgres_uri", defaultPostgresURI, "postgres URI")
	postgresMaxCons := flag.Int("postgres_pool_workers", defaultPostgresPoolWorkers, "maximum number of Postgre workers used simultaneously")
	ctxTimeout := flag.Int("ctx_timeout", defaultCtxTimeOut, "context timeout")
	flag.Parse()

	switch *mode {
	case "debug":
		if err := godotenv.Load("internal/app/server/config/.env"); err != nil {
			return Config{}, err
		}
	case "release":

	default:
		log.Fatal("invalid mode: " + *mode)
	}

	var cfg = Config{}

	cfg.RunAddress = getEnvString(*runAddress, "RUN_ADDRESS")

	cfg.Postgres.URI = getEnvString(*postgresURI, "POSTGRES_URI")
	cfg.Postgres.Workers = getEnvInt(*postgresMaxCons, "POSTGRES_WORKERS")

	cfg.CtxTimeout = time.Millisecond * time.Duration(getEnvInt(*ctxTimeout, "CTX_TIMEOUT"))

	return cfg, nil
}

// getEnvString - function for determining the priority between flags and environment variables in string format
func getEnvString(flagValue string, envKey string) string {
	envValue, exists := os.LookupEnv(envKey)
	if exists {
		return envValue
	}
	return flagValue
}

// getEnvInt - function for determining the priority between flags and environment variables in int format
func getEnvInt(flagValue int, envKey string) int {
	envValue, exists := os.LookupEnv(envKey)
	if exists {
		intVal, err := strconv.Atoi(envValue)
		if err != nil {
			log.Printf("cant convert env-key: %s to int", envValue)
			return 1
		}

		return intVal
	}

	return flagValue
}
