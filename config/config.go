package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Configtoml is
func Configtoml() (string,string,string) {

	viper.SetConfigName("config")                // name of config file (without extension)
	viper.SetConfigType("toml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/Github/icapeg-client") // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	file := viper.GetString("configuration.download_file")
	host := viper.GetString("icap_configuration.host")
	port := viper.GetString("icap_configuration.port")
	return file,host,port

}
