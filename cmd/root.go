package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/joyrex2001/asciime/backend"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "asciime",
	Short: "ASCIIme is an application that converts images to ascii art.",
	Long:  ``,
	Run:   backend.Main,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().String("healthz-addr", ":8088", "Webserver /healthz port")
	rootCmd.PersistentFlags().String("listen-addr", ":8080", "Webserver listen address")
	rootCmd.PersistentFlags().Bool("enable-tls", false, "Enable TLS on webserver")
	rootCmd.PersistentFlags().String("key-file", "", "TLS keyfile")
	rootCmd.PersistentFlags().String("cert-file", "", "TLS certificate file")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose mode")
	viper.BindPFlag("generic.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("health.listen-addr", rootCmd.PersistentFlags().Lookup("healthz-addr"))
	viper.BindPFlag("web.listen-addr", rootCmd.PersistentFlags().Lookup("listen-addr"))
	viper.BindPFlag("web.enable-tls", rootCmd.PersistentFlags().Lookup("enable-tls"))
	viper.BindPFlag("web.cert-file", rootCmd.PersistentFlags().Lookup("cert-file"))
	viper.BindPFlag("web.key-file", rootCmd.PersistentFlags().Lookup("key-file"))
	viper.BindEnv("health.listen-addr", "HEALTH_LISTEN_ADDR")
	viper.BindEnv("web.listen-addr", "WEB_LISTEN_ADDR")
	viper.BindEnv("web.enable-tls", "WEB_ENABLE_TLS")
	viper.BindEnv("web.cert-file", "WEB_CERT_FILE")
	viper.BindEnv("web.key-file", "WEB_KEY_FILE")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "kopenvoor" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		// fmt.Printf("not using config file: %s\n", err)
	} else {
		fmt.Printf("using config: %s\n", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
