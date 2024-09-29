package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var defaults = map[string]interface{}{
	"debug": false,
}

func main() {
	v := viper.GetViper()

	err := initializeConfig(v)

	if err != nil {
		log.Fatalln("Failed to initialize config: " + err.Error())
	}

	println(fmt.Sprintf("Debug: %s", v.GetBool("debug")))
}

func initializeConfig(v *viper.Viper) error {
	v.SetEnvPrefix("video_storybook")

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	return v.ReadInConfig()
}
