package services

import "fmt"

type Config struct {
	HTTP struct {
		Host string `env:"HOST"`
		Port string `env:"PORT"`
	}
	Database struct {
		Name                  string `env:"PG_DB"`
		Host                  string `env:"PG_HOST"`
		Port                  string `env:"PG_PORT"`
		Username              string `env:"PG_USER"`
		Password              string `env:"PG_PASSWORD"`
		PoolMaxConnections    int    `env:"PG_POOL_MAX_CONNECTIONS"`
		PoolMinConnections    int    `env:"PG_POOL_MIN_CONNECTIONS"`
		PoolHealthCheckPeriod string `env:"PG_POOL_HEALTH_CHECK_PERIOD"`
	}
	Migration struct {
		SeqDigits int    `env:"MIGRATION_SEQ_DIGITS"`
		Directory string `env:"MIGRATION_PG_MIGRATION_DIR"`
	}
}

func (c *Config) GetDbUrl(db string) string {
	switch db {
	case "app":
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%d&pool_min_conns=%d&pool_health_check_period=%s", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name, c.Database.PoolMaxConnections, c.Database.PoolMinConnections, c.Database.PoolHealthCheckPeriod)
	case "migrate":
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	}

	return ""
}

func (c *Config) GetHttpUrl() string {
	return fmt.Sprintf("%s:%s", c.HTTP.Host, c.HTTP.Port)
}
