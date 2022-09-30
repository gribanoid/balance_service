package config

type Config struct {
	PostgresHost string `envconfig:"POSTGRES_HOST" default:"user=postgres password=qwerty host=db port=5432 dbname=db pool_max_conns=10"`
}
