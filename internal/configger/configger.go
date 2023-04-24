package configger

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	log.Println("Config loaded successfully...")
	log.Println("Getting environment variables...")
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
}

func getEnvOrPanic(env string) string {
	if !strings.Contains(env, ":") {
		log.Fatalf("Log format variable %s is incorrect. ':' missing", env)
	}

	varSplit := strings.Split(env, ":")
	envVar := varSplit[0]
	defaultVar := varSplit[1]

	res := os.Getenv(envVar)
	if len(res) == 0 {
		res = defaultVar
	}
	return res
}
