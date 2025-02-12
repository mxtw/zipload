package cmd

import (
	"fmt"
	"log"

	"github.com/mxtw/zipload/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload a file to a zipline server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		host := viper.GetString("host")
		client := api.NewClient(token, host)
		urls, err := client.Upload(args[0])
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, url := range urls {
			fmt.Println(url)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
