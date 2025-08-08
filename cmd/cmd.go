package cmd

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/jasonkwh/coffeed/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:     "coffeed",
		Short:   "coffee schedule daemon",
		Version: buildVersion(),
		Run:     root,
	}
	cfgFile string
	cfg     config.Config

	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Put all the config in a common struct
	if err := viper.Unmarshal(&cfg); err == nil {
		fmt.Printf("failed to unmarshal config: %v\n", err)
	}
}

func initZapLogger() (*zap.Logger, error) {
	zlCfg := zap.NewProductionConfig()

	// set the internal logger to INFO because we need all internal logs
	zlCfg.Level.SetLevel(zapcore.Level(cfg.LogLevel))

	return zlCfg.Build()
}

func gracefulClose(services []io.Closer) error {
	var errs error

	for _, item := range services {
		err := item.Close()
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	return errs
}

func buildVersion() string {
	return fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date)
}
