package main

import (
	"fmt"
	"log"
	"os"

	envconfig "github.com/agussyahrilmubarok/gohelp/config/env"
)

type AppConfig struct {
	AppName  string `mapstructure:"APP_NAME"`
	Port     int    `mapstructure:"APP_PORT"`
	Debug    bool   `mapstructure:"APP_DEBUG"`
	Database string `mapstructure:"DB_HOST"`
}

func main() {
	os.Setenv("APP_NAME", "EnvApp")
	os.Setenv("APP_PORT", "9090")
	os.Setenv("APP_DEBUG", "true")
	os.Setenv("DB_HOST", "localhost")

	cfg := envconfig.New()

	appName, _ := cfg.GetString("APP_NAME")
	port, _ := cfg.GetInt("APP_PORT")
	debug, _ := cfg.GetBool("APP_DEBUG")

	fmt.Println("App Name:", appName)
	fmt.Println("Port    :", port)
	fmt.Println("Debug   :", debug)

	var appConfig AppConfig
	if err := cfg.Unmarshal(&appConfig); err != nil {
		log.Fatalf("Failed to unmarshal env config: %v", err)
	}

	fmt.Println("\n=== Full Config Struct ===")
	fmt.Printf("%+v\n", appConfig)
	fmt.Println("Database Host:", appConfig.Database)
}
