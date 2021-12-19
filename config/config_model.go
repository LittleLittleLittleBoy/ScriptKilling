package config

type Config struct {
	Mysql MySQLConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Port     int
	Host     string
	Database string
}
