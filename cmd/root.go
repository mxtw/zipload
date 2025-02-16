package cmd

import (
	"log"
	"os"

	"github.com/mxtw/zipload/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "zipload",
	Short: "Simple API client to zipline",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	options api.Options
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zipload/zipload.yaml)")

	options = api.Options{
		Format:                  api.FormatFlag{},
		ImageCompressionPercent: 0,
		ExpiresAt:               "",
		Password:                "",
		ZeroWidthSpace:          false,
		Embed:                   false,
		MaxViews:                0,
		UploadText:              false,
		XZiplineFilename:        "",
		OriginalName:            false,
		OverrideDomain:          "",
		XZiplineFolder:          0,
	}

	var token string
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "zipline API token")
	rootCmd.MarkFlagRequired("token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	var host string
	rootCmd.PersistentFlags().StringVar(&host, "host", "", "zipline host")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	rootCmd.MarkFlagRequired("host")

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.zipload")
		viper.AddConfigPath("$XDG_CONFIG_HOME/zipload")
		viper.SetConfigName("zipload")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Can't read config:", err)
	}
}
