package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func LoadConfig(Vars []string) {
	variablesToValidate := Vars
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.AddConfigPath(".")
	viper.AllowEmptyEnv(false)
	replacer := strings.NewReplacer("__", ".")
	viper.SetEnvKeyReplacer(replacer)
	for _, variable := range variablesToValidate {
		err := viper.BindEnv(variable)
		if err != nil {
			println("variable", variable, err.Error())
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			fmt.Println("found config file but failed to load it", err)
		}
	}
	for _, variable := range variablesToValidate {
		if viper.GetString(variable) == "" {
			log.Fatal(fmt.Sprintf("variable %s not defined", variable))
		}
	}
}
