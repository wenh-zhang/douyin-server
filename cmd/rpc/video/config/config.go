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

type RPCConfig struct {
	Host string
	Port int
	Name string
}