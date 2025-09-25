package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port            string
	BaseUrl         string
	AuthToken       string
	DbUrl           string
	BoxOfficeUrl    string
	BoxOfficeApiKey string
}

var Conf = &Config{}

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	Conf.Port = os.Getenv("PORT")
	Conf.BaseUrl = os.Getenv("BASE_URL")
	Conf.AuthToken = os.Getenv("AUTH_TOKEN")
	Conf.DbUrl = os.Getenv("DB_URL")
	Conf.BoxOfficeUrl = os.Getenv("BOXOFFICE_URL")
	Conf.BoxOfficeApiKey = os.Getenv("BOXOFFICE_API_KEY")

	return nil
}
