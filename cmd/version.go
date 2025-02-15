package cmd

import (
	"fmt"
	"log"

	"github.com/mxtw/zipload/pkg/api"
	"github.com/mxtw/zipload/pkg/api/version"
	"github.com/spf13/cobra"
)

// im intentionally not setting version on rootCmd for now so i
// can have my custom handling of the version subcommand
// TODO figure out a better way of doing this
var Version string = "dev"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get client and server versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("zipload version: %s\n", Version)

		if !clientOnly {
			client := api.NewClient()
			serverVersion, err := version.Version(&client)
			if err != nil {
				log.Fatalf("could not get server version")
			}

			fmt.Printf("zipline server version: %s\n", serverVersion.Versions.Current)

			if serverVersion.Update {
				fmt.Printf("zipline server version %s is available\n", serverVersion.Versions.Stable)
			}
		}
	},
}

var clientOnly bool

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolVar(&clientOnly, "client-only", false, "if set, only get client version")
}
