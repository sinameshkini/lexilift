package cmd

import (
	"fmt"
	"lexilift/app"
	"lexilift/internal/config"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $COMMAND_HOME/config.yml)")
}

func initConfig() (conf *config.Config) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs/")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		if err = viper.Unmarshal(&conf); err != nil {
			log.Fatalln(err.Error())
		}

	} else {
		conf = &config.DefaultConfig
	}

	return
}

var (
	cfgFile string

	serverCmd = &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "Run web service",
		Long:    `Restful API web service for Lexilift web application.`,
		Run:     server,
	}
)

func server(cmd *cobra.Command, args []string) {
	conf := initConfig()

	if err := app.Server(conf); err != nil {
		log.Fatalln(err.Error())
	}
}
