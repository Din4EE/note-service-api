package config

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DB   DBConfig   `yaml:"db"`
	HTTP HTTPConfig `yaml:"http"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type DBConfig struct {
	DSN            string `yaml:"dsn"`
	MaxConnections int32  `yaml:"max_connections"`
}

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type GRPCConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config{}
	tmpl, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %s", err.Error())
	}

	t, err := template.New("config").Parse(string(tmpl))
	if err != nil {
		log.Fatalf("failed to parse config file: %s", err.Error())
	}

	env := map[string]string{
		"DB_DSN":    "host=" + os.Getenv("POSTGRES_HOST") + " port=" + os.Getenv("POSTGRES_PORT") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=" + os.Getenv("POSTGRES_SSLMODE"),
		"HTTP_HOST": os.Getenv("HTTP_SERVER_HOST"),
		"HTTP_PORT": os.Getenv("HTTP_SERVER_PORT"),
		"GRPC_HOST": os.Getenv("GRPC_SERVER_HOST"),
		"GRPC_PORT": os.Getenv("GRPC_SERVER_PORT"),
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, env); err != nil {
		log.Fatalf("failed to execute config file: %s", err.Error())
	}

	if err = yaml.Unmarshal(buf.Bytes(), &config); err != nil {
		log.Fatalf("failed to unmarshal config file: %s", err.Error())
	}

	return config, nil
}

func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	cfg, err := pgxpool.ParseConfig(c.DB.DSN)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = c.DB.MaxConnections

	return cfg, nil
}
