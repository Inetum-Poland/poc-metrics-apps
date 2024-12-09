package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"inetum.com/metrics-go-app/internal/config"
)

// Priority of configuration sources:
// - command-line flags
// - environment variables
// - config file values
// - defaults set on command-line flags

var (
	envPrefix                  = "APP"
	replaceHyphenWithCamelCase = false
)

func RootCommand() *cobra.Command {
	mainCmd := &cobra.Command{
		Short: "App Metrics",
		Long:  `Demonstrate how go-otel works`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	// WebApp
	mainCmd.Flags().StringVar(&config.C.WebApp.Name, "webapp_name", "App", "Name of the WebApp")
	mainCmd.Flags().IntVar(&config.C.WebApp.Port, "webapp_port", 8080, "Port to listen on")
	// OpenTelemetry
	mainCmd.Flags().StringVar(&config.C.Otel.Host, "otel_host", "127.0.0.1", "OpenTelemetry URL")
	mainCmd.Flags().IntVar(&config.C.Otel.Port, "otel_port", 4317, "OpenTelemetry Port")
	// MongoDB
	mainCmd.Flags().StringVar(&config.C.Mongo.Host, "mongo_host", "127.0.0.1", "MongoDB URL")
	mainCmd.Flags().IntVar(&config.C.Mongo.Port, "mongo_port", 27017, "MongoDB Port")
	mainCmd.Flags().StringVar(&config.C.Mongo.User, "mongo_user", "root", "MongoDB User")
	mainCmd.Flags().StringVar(&config.C.Mongo.Pass, "mongo_pass", "Password123", "MongoDB Password")

	return mainCmd
}

func initializeConfig(cmd *cobra.Command) {
	v := viper.New()

	v.SetConfigName("config.yml")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	v.ReadInConfig()

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	bindFlags(cmd, v)
}

// Bind each cobra flag to its associated viper configuration environment variable
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := f.Name

		if replaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
