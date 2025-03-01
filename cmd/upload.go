package cmd

import (
	"fmt"
	"log"

	"github.com/mxtw/zipload/pkg/api"
	"github.com/mxtw/zipload/pkg/api/upload"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload a file to a zipline server",
	Long: `upload a file to a zipline server. 

    refer to https://zipline.diced.sh/docs/guides/upload-options for more info on the upload options`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		options := upload.Options{
			DeletesAt:               deletesAt,
			Domain:                  domain,
			FileExtension:           fileExtension,
			Filename:                filename,
			Folder:                  folder,
			Format:                  format.Value,
			ImageCompressionPercent: imageCompressionPercent,
			MaxViews:                maxViews,
			OriginalName:            keepOriginalName,
			Password:                password,
		}

		files, err := upload.Upload(&client, args[0], options)
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, file := range files {
			fmt.Println(file.Url)
		}
	},
}

var (
	deletesAt               string
	domain                  string
	fileExtension           string
	filename                string
	folder                  uint
	format                  upload.FormatFlag
	imageCompressionPercent uint8
	keepOriginalName        bool
	maxViews                uint
	password                string
)

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVar(&deletesAt, "deletes-at", "", "when the link should expire, e.g. '1d', '2 months', etc.")
	uploadCmd.Flags().StringVar(&domain, "domain", "", "override the domain used for the link")
	uploadCmd.Flags().StringVar(&fileExtension, "file-extension", "", "override file extension in zipline")
	uploadCmd.Flags().StringVar(&filename, "filename", "", "override filename in zipline")
	uploadCmd.Flags().UintVar(&folder, "folder", 0, "id of folder to save file in")
	uploadCmd.Flags().Var(&format, "format", "format of the file name (one of random, uuid, date, name, gfycat)")
	uploadCmd.Flags().Uint8Var(&imageCompressionPercent, "image-compression-percent", 0, "image compression percentage")
	uploadCmd.Flags().UintVar(&maxViews, "max-views", 0, "maximum allowed views on link")
	uploadCmd.Flags().BoolVar(&keepOriginalName, "keep-original-name", false, "keep original filename")
	uploadCmd.Flags().StringVar(&password, "password", "", "choose password to protect the link")
}
