package cmd

import (
	"fmt"
	"log"

	"github.com/mxtw/zipload/pkg/api"
	"github.com/mxtw/zipload/pkg/api/user/urls"
	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Shorten a url",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		options := urls.Options{
			MaxViews: maxViews,
			Domain:   domain,
			Password: password,
		}

		url, err := urls.Url(&client, args[0], vanity, options)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println(url)
	},
}

var (
	vanity string
)

func init() {
	rootCmd.AddCommand(urlCmd)
	urlCmd.Flags().StringVar(&vanity, "vanity", "", "vanity url")
	urlCmd.Flags().StringVar(&password, "password", "", "protect link with a password")
	urlCmd.Flags().StringVar(&domain, "domain", "", "override default domain")
	urlCmd.Flags().UintVar(&maxViews, "max-views", 0, "maximum allowed views on link")
}
