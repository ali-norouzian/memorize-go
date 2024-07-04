package config

import (
	"fmt"
	"memorize/pkg/files"
	"path/filepath"
)

func NewConfig() (*Config, error) {
	var appConfig Config

	projectRoot, err := files.FindProjectRoot()
	if err != nil {
		return nil, err
	}

	configFilePath := filepath.Join(*projectRoot, configFolderName, configFileName)

	err = files.ReadJsonFile(configFilePath, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}

func (postgreDbConfig *PostgreDbConfig) GetDbConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		postgreDbConfig.Host,
		postgreDbConfig.Port,
		postgreDbConfig.Username,
		postgreDbConfig.DbName,
		postgreDbConfig.Password,
		postgreDbConfig.SslMode,
	)
}

func (config *Config) GetDbSetting() *PostgreDbConfig {
	return &config.PostgreDbConfig
}

func (config *Config) GetJwtSetting() *Jwt {
	return &config.Jwt
}
