package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig      `map:"app"`
	Db  DatabaseConfig `map:"db"`
}

type AppConfig struct {
	Host 			string `map:"host"`
	Port 			int    `map:"port"`
	JWTSecret 		string    `map:"jwtSecret"`
}

type DatabaseConfig struct {
	Host     string `map:"host"`
	Port     int    `map:"port"`
	Username string `map:"username"`
	Password string `map:"password"`
	Name string `map:"name"`
}

var Configs Config

func init() {
	viper.SetConfigFile("../../configs/config.yaml")
	readConfigErr := viper.ReadInConfig()
	if readConfigErr != nil {
		fmt.Println(readConfigErr)
		log.Fatal("Reading Config Error")
	}
	
	err := viper.Unmarshal(&Configs)
	if err != nil {
		log.Fatal("Viper Unmarshal Error:", err)
	}
}