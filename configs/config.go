package configs

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

type conf struct {
	WebserverPort string `mapstructure:"WEBSERVER_PORT"`
	RpsLimitIp    string `mapstructure:"RPS_LIMIT_IP"`
	RpsLimitToke  string `mapstructure:"RPS_LIMIT_TOKEN"`
}

func defaultAndBindings() error {
	defaultConfigs := map[string]string{
		"WEBSERVER_PORT":  "8080",
		"RPS_LIMIT_IP":    "5",
		"RPS_LIMIT_TOKEN": "10",
	}
	for envKey, envValue := range defaultConfigs {
		err := viper.BindEnv(envKey)
		if err != nil {
			return err
		}
		viper.SetDefault(envKey, envValue)
	}
	return nil

}
func LoadConfig(workdir string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	_, err := os.Stat(path.Join(workdir, ".env"))
	if err == nil {
		viper.SetConfigType("env")
		viper.AddConfigPath(workdir)
		viper.SetConfigFile(".env")
		err = viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	viper.AutomaticEnv()
	err = defaultAndBindings()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
