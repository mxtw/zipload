/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/mxtw/zipload/pkg/api"
	"github.com/mxtw/zipload/pkg/api/shorten"
	"github.com/spf13/cobra"
)

// shortenCmd represents the shorten command
var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		options := shorten.Options{
			ZeroWidthSpace: zeroWidthSpace,
			MaxViews:       maxViews,
		}

		url, err := shorten.Shorten(&client, args[0], vanity, options)
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
	rootCmd.AddCommand(shortenCmd)
	shortenCmd.Flags().StringVar(&vanity, "vanity", "", "vanity url")
	shortenCmd.Flags().BoolVar(&zeroWidthSpace, "zero-width-space", false, "url should use a zero width space")
	shortenCmd.Flags().UintVar(&maxViews, "max-views", 0, "maximum allowed views on link")
}
