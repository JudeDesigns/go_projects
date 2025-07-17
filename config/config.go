package config

import "os"

type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

//Created a struct that fully matches our .env contents

func LoadConfig() AppConfig {

	//Created an Object of type AppConfig by populating the fields with the values from the .env file
	config := AppConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
	}
	//Returning the config object
	return config
}
