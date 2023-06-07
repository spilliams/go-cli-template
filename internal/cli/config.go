package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var env *config

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".foo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	raw := viper.AllSettings()
	rawBytes, err := json.Marshal(raw)
	cobra.CheckErr(err)

	err = json.Unmarshal(rawBytes, &env)
	cobra.CheckErr(err)
}

type config struct {
	Verbose bool `json:"verbose"`
}
