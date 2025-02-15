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
			Format:                  format.Value,
			ImageCompressionPercent: imageCompressionPercent,
			ExpiresAt:               expiresAt,
			Password:                password,
			ZeroWidthSpace:          zeroWidthSpace,
			Embed:                   embed,
			MaxViews:                maxViews,
			UploadText:              uploadText,
			XZiplineFilename:        ziplineFilename,
			OriginalName:            keepOriginalName,
			OverrideDomain:          overrideDomain,
			XZiplineFolder:          ziplineFolder,
		}

		urls, err := upload.Upload(&client, args[0], options)
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, url := range urls {
			fmt.Println(url)
		}
	},
}

var (
	format                  upload.FormatFlag
	imageCompressionPercent uint8
	expiresAt               string
	password                string
	zeroWidthSpace          bool
	embed                   bool
	maxViews                uint
	uploadText              bool
	ziplineFilename         string
	keepOriginalName        bool
	overrideDomain          string
	ziplineFolder           uint
)

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().Var(&format, "format", "format of the file name (one of random, uuid, date, name, gfycat)")
	uploadCmd.Flags().Uint8Var(&imageCompressionPercent, "image-compression-percent", 0, "image compression percentage")
	uploadCmd.Flags().StringVar(&expiresAt, "expires-at", "", "when the link should expire, e.g. '1d', '2 months', etc.")
	uploadCmd.Flags().StringVar(&password, "password", "", "choose password to protect the link")
	uploadCmd.Flags().BoolVar(&zeroWidthSpace, "zero-width-space", false, "url should use a zero width space")
	uploadCmd.Flags().BoolVar(&embed, "embed", true, "make file embeddable in e.g. discord")
	uploadCmd.Flags().UintVar(&maxViews, "max-views", 0, "maximum allowed views on link")
	uploadCmd.Flags().BoolVar(&uploadText, "upload-text", false, "always upload as text/plain")
	uploadCmd.Flags().StringVar(&ziplineFilename, "zipline-filename", "", "override filename in zipline")
	uploadCmd.Flags().BoolVar(&keepOriginalName, "keep-original-name", false, "keep original filename")
	uploadCmd.Flags().StringVar(&overrideDomain, "override-domain", "", "override the domain used for the link")
	uploadCmd.Flags().UintVar(&ziplineFolder, "zipline-folder", 0, "id of folder to save file in")
}
