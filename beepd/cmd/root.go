package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xbt573/beepsuite/beepd/internal/app"
)

var (
	configFile string
	config     Config
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Config file (.yaml)")
	rootCmd.PersistentFlags().StringP("listen", "l", "0.0.0.0:3000", "Address to which server listens to")

	viper.BindPFlag("listen", rootCmd.PersistentFlags().Lookup("listen"))

	cobra.OnInitialize(func() {
		if configFile != "" {
			viper.SetConfigFile(configFile)
		} else {
			viper.AddConfigPath(".")
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
		}

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				panic(err)
			}
		}

		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
}

var rootCmd = &cobra.Command{
	Use: "beepd",
	RunE: func(cmd *cobra.Command, args []string) error {
		a := app.New()

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		return a.Serve(ctx, config.Listen)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
