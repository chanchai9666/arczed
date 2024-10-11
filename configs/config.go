package configs

import "github.com/caarlos0/env/v11"

type Config struct {
	Port         int    `env:"PORT" envDefault:"8080"`
	AppEnv       string `env:"APP_ENV" envDefault:"local"`
	DbHost       string `env:"DB_HOST" envDefault:"localhost"`
	DbPort       string `env:"DB_PORT" envDefault:"5432"`
	DbDatabase   string `env:"DB_DATABASE" envDefault:"arczedxxxx"`
	DbUsername   string `env:"DB_USERNAME" envDefault:"admin"`
	DbPassword   string `env:"DB_PASSWORD" envDefault:"1234"`
	JwtSECRETKEY string `env:"SECRET_KEY" envDefault:"the_secret_key"`
}

// LoadConfig โหลดค่า environment variables ลงใน config struct
func LoadConfig() (*Config, error) {
	var config Config
	// ใช้ env.Parse เพื่อโหลดค่า environment variables ลงใน struct
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
