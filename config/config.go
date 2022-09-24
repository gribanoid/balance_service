package config

type Config struct {
	PostgresHost string `envconfig:"POSTGRES_HOST" default:"user=postgres password=qwerty host=localhost port=5432 dbname=cryptopayments pool_max_conns=10"`
}
