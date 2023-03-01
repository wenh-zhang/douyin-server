package config

type EtcdConfig struct {
	Host string
	Port int
}

type MinioConfig struct {
	Host            string
	Port            int
	AccessKeyID     string
	SecretAccessKey string
	UserSSL         bool
	Bucket          string
}

type AmqpConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}
