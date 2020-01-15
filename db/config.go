package db

import (
	"fmt"
	"github.com/spf13/viper"
)

// DatabaseConfig represents the config of the database
type DatabaseConfig struct {
	URI      string
	Username string
	Password string
}

// InitDatabaseConfig gets the configuration from the Viper file and returns a DatabaseConfig object
func InitDatabaseConfig() (*DatabaseConfig, error) {
	dbConfig := &DatabaseConfig{
		URI:      viper.GetString("db.uri"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
	}

	if len(dbConfig.URI) == 0 {
		return nil, fmt.Errorf("Database URL must be set")
	}

	if len(dbConfig.Username) == 0 || len(dbConfig.Password) == 0 {
		return nil, fmt.Errorf("Database username and password must be set")
	}

	return dbConfig, nil
}
