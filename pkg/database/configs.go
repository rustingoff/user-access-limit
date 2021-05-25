package database

type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

type RedisConfig struct {
	Host string
	Port string
}
