package config

type EtcdConfig struct {
	Host string
	Port int
}

type MySQLConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type RPCConfig struct {
	Host string
	Port int
	Name string
}

type AmqpConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}
