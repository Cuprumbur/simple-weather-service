package configuration

type Config struct {
	Port int
	DB   *DbInfo
}

type DbInfo struct {
	User string
	Pass string
	Name string
}

func NewConfig() *Config {
	return &Config{
		Port: 3001,
		DB: &DbInfo{
			User: "root",
			Pass: "my-secret-pw",
			Name: "my_db",
		},
	}
}
