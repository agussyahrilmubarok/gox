package main

import (
	"fmt"
	"log"

	viperconfig "github.com/agussyahrilmubarok/gohelp/config/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type AppConfig struct {
	AppName  string
	Port     int
	Debug    bool
	Database DatabaseConfig
}

func main() {
	cfg, err := viperconfig.New("./config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	appName, _ := cfg.GetString("app_name")
	port, _ := cfg.GetInt("port")
	debug, _ := cfg.GetBool("debug")

	fmt.Println("App Name:", appName)
	fmt.Println("Port    :", port)
	fmt.Println("Debug   :", debug)

	var appConfig AppConfig
	if err := cfg.Unmarshal(&appConfig); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	fmt.Println("\n=== Full Config Struct ===")
	fmt.Printf("%+v\n", appConfig)

	fmt.Println("\nDatabase Host:", appConfig.Database.Host)
	fmt.Println("Database Port:", appConfig.Database.Port)
}
