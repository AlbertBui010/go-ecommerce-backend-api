package initialize

import (
	"fmt"

	"github.com/albertbui010/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// fmt.Println("Server port::", viper.GetInt("server.port"))
	// fmt.Println("JWT key::", viper.GetString("security.jwt.key"))

	// Config structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	// fmt.Println("Config Port::", global.Config.Server.Port)
	// fmt.Println("Config DB user::", global.Config.Databases[0].User)
}
