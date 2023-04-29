package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var cfg Config

type Config struct {
	BasePath string `yaml:"base_path"`
	EnvPath  string `yaml:"env_path"`
	AppPort  string `yaml:"app_port"`

	Allowed struct {
		Hosts []string `yanml:"hosts"`
	} `yaml:"allowed"`

	PostgreSQL struct {
		Use      bool   `yaml:"use"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		DBName   string `yaml:"db_name"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"postgres"`
}

func InitConfig(args []string) (*Config, error) {
	var configPath string

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&configPath, "c", "config.yaml", "set path to config")
	err := flags.Parse(args[1:])
	if err != nil {
		return nil, err
	}

	clean := filepath.Clean(configPath)

	file, err := os.Open(clean)
	if err != nil {
		return nil, fmt.Errorf("fail to open config file in path \"%s\" with error %w", configPath, err)
	}

	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to parse config %w", err)
	}

	if err := godotenv.Load(cfg.EnvPath); err != nil {
		logrus.Errorf("error reading environment file: %s", err.Error())
	}

	return &cfg, nil
}
