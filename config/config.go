package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var config *Configuration

type Configuration struct {
	Port         int    `env:"PORT" envDefault:"8080"`
	Version      string `env:"VERSION" envDefault:"1.0.0"`
	Postgres     PostgresConfig
	Redis        RedisConfig
	GRPCEndpoint GRPCEndpointConfig
	Debug        bool `env:"DEBUG" envDefault:"false"`
	Migrate      bool `env:"MIGRATE" envDefault:"true"`
	CacheEnabled bool `env:"CACHE_ENABLED" envDefault:"true"`
	// AMQPConnectionString string `env:"AMQP_CONNECTION_STRING" envDefault:"amqp://guest:guest@localhost:5672/"`
	// FrontendURL          string `env:"FRONTEND_URL" envDefault:"https://develop.metub.dev"`
	// JwtSecret            string `env:"JWT_SECRET"`
	// UploadDir            string `env:"UPLOAD_DIR"`
	// UploadURL            string `env:"UPLOAD_URL"`
}

type PostgresConfig struct {
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port     string `env:"DATABASE_PORT" envDefault:"5432"`
	User     string `env:"DATABASE_USER" envDefault:"postgres"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE_NAME" envDefault:"postgres"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	Database int    `env:"REDIS_DB" envDefault:"0"`
}

type GRPCEndpointConfig struct {
	Identity     string `env:"IDENTITY_ENDPOINT"`
	DataParser   string `env:"DATA_PARSER_ENDPOINT"`
	MasterData   string `env:"MASTERDATA_ENDPOINT"`
	EContract    string `env:"ECONTRACT_ENDPOINT"`
	Asset        string `env:"ASSET_ENDPOINT"`
	Notification string `env:"NOTIFICATION_ENDPOINT"`
	Email        string `env:"EMAIL_ENDPOINT"`
	Media        string `env:"MEDIA_ENDPOINT"`
	Impex        string `env:"IMPEX_ENDPOINT"`
}

// newConfig will read the config data from given .env file
func newConfig(files ...string) *Configuration {
	err := godotenv.Load(files...) // Loading config from env file

	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{Postgres: PostgresConfig{}, Redis: RedisConfig{}}
	// Parse env to configuration
	if err = env.Parse(&cfg); err != nil {
		fmt.Printf("Config error: %+v\n", err)
	}

	// posgres config
	if err = env.Parse(&cfg.Postgres); err != nil {
		fmt.Printf("Postgres Config error: %+v\n", err)
	}

	// redis config
	if err = env.Parse(&cfg.Redis); err != nil {
		fmt.Printf("Redis Config error: %+v\n", err)
	}

	// grpc endpoints config
	if err = env.Parse(&cfg.GRPCEndpoint); err != nil {
		fmt.Printf("GRPC Endpoint Config error: %+v\n", err)
	}

	return &cfg
}

func GetInstance() *Configuration {
	if config == nil {
		config = newConfig(".env", "base.env")
	}
	return config
}

func LoadConfig() (*Configuration, error) {
	config := GetInstance()
	return config, nil
}
