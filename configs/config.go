package configs

import (
	"os"
	"sync"
)

type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()
	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Database.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Database.Name = getEnv("NAME", "project_ecommerce_k2")
	defaultConfig.Database.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = getEnv("USERNAME", "root")
	defaultConfig.Database.Password = getEnv("PASSWORD", "root")

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./configs")
	// // log.Info(viper.ReadInConfig())
	// if err := viper.ReadInConfig(); err != nil {
	// 	// fmt.Println("kok mlaku")
	// 	log.Info("error in open file")
	// 	return &defaultConfig
	// }

	// var finalConfig AppConfig

	// if err := viper.Unmarshal(&finalConfig); err != nil {
	// 	log.Info("error in extract external config, must use default config")
	// 	return &defaultConfig
	// }

	// // fmt.Println(defaultConfig)

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "user" {
		// fmt.Println(value)
		return value
	}

	return fallback

}
