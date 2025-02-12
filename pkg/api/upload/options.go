package upload

import (
	"log"
	"net/http"
	"strconv"
)

// refer to https://zipline.diced.sh/docs/guides/upload-options
type Options struct {
	Format                  Format
	ImageCompressionPercent uint8
	ExpiresAt               string
	Password                string
	ZeroWidthSpace          bool
	Embed                   bool
	MaxViews                uint
	UploadText              bool
	XZiplineFilename        string
	OriginalName            bool
	OverrideDomain          string
	XZiplineFolder          uint
}

// i dont really like this setup but i guess its functional
func (o Options) toHeaders() http.Header {
	headers := http.Header{}

	formatFlag := FormatFlag{Value: o.Format}

	headers.Set("Format", formatFlag.String())
	headers.Set("Image-Compression-Percent", strconv.FormatUint(uint64(o.ImageCompressionPercent), 10))
	headers.Set("Expires-At", o.ExpiresAt)
	headers.Set("Password", o.Password)
	headers.Set("Zws", strconv.FormatBool(o.ZeroWidthSpace))
	headers.Set("Embed", strconv.FormatBool(o.Embed))
	headers.Set("Max-Views", strconv.FormatUint(uint64(o.MaxViews), 10))
	headers.Set("UploadText", strconv.FormatBool(o.UploadText))
	headers.Set("X-Zipline-Filename", o.XZiplineFilename)
	headers.Set("Override-Domain", o.OverrideDomain)
	headers.Set("X-Zipline-Folder", strconv.FormatUint(uint64(o.XZiplineFolder), 10))

	log.Println(headers)

	return headers
}
