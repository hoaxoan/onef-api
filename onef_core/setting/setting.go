package setting

import (
	"net/http"

	"github.com/hoaxoan/onef-api/onef_auth/base"
	"github.com/hoaxoan/onef-api/onef_core/model"
	"github.com/hoaxoan/onef-api/onef_core/setting"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
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
	JWTSecret struct {
		JWTKey string `mapstructure:"JWTEncodeKey"`
	} `mapstructure:"JWTSeret"`
}
