package setting

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config ConfigSchema

func init() {
	viper.SetConfigName("setting")             // name of config file (without extension)
	viper.AddConfigPath("./onef_core/setting") // optionally look for config in the working directory
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

type ConfigSchema struct {
	Mongo struct {
		Uri      string `mapstructure:"Uri"`
		Host     string `mapstructure:"HostName"`
		Username string `mapstructure:"Username"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"MongoDB"`
	Mysql struct {
		Host     string `mapstructure:"Host"`
		Port     string `mapstructure:"Port"`
		DbName   string `mapstructure:"DbName"`
		Username string `mapstructure:"Username"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"Mysql"`
	JWTSecret struct {
		JWTKey string `mapstructure:"JWTEncodeKey"`
	} `mapstructure:"JWTSeret"`
}
