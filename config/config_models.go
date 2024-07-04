package config

type Config struct {
	PostgreDbConfig PostgreDbConfig `json:"postgreDbConfig"`
	Jwt             Jwt             `json:"jwt"`
}

type PostgreDbConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	DbName   string `json:"dbName"`
	SslMode  string `json:"sslMode"`
}

type Jwt struct {
	Secret              string `json:"secret"`
	ExpirationTimeInDay int    `json:"expirationTimeInDay"`
}
