package config

import (
	"github.com/bells307/everydaypic/pkg/minio"
	"github.com/bells307/everydaypic/pkg/mongodb"
	"github.com/spf13/viper"
)

type Config struct {
	Listen  string                  `mapstructure:"LISTEN"`
	MongoDB mongodb.MongoDBConfig   `mapstructure:"MONGODB"`
	MinIO   minio.MinIOClientConfig `mapstructure:"MINIO"`
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
