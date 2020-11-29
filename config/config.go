package config

import (
	"time"
	"fmt"

	"github.com/spf13/viper"
)

//Configtoml is
func Configtoml() (string,string,string,string,time.Duration,string) {

	viper.SetConfigName("config")                
	viper.SetConfigType("toml")                  
	viper.AddConfigPath("/Github/icapeg-client") 
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() 
	if err != nil {             
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	file := viper.GetString("configuration.download_file")
	host := viper.GetString("icap_configuration.host")
	port := viper.GetString("icap_configuration.port")
	service := viper.GetString("icap_configuration.service")
	timeout := viper.GetDuration("icap_timeout.timeout")
	filepath := viper.GetString("configuration.file_path")
	return file,host,port,service,timeout,filepath

}
