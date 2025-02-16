package api

import (
	"log"
	"net/http"
	"strconv"
)

// refer to https://zipline.diced.sh/docs/guides/upload-options
type Options struct {
	Format                  FormatFlag
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

func (o Options) ToHeaders() http.Header {
	headers := http.Header{}

	// use a helper function to only set headers when it makes sense to set
	// them especially bool headers will be respected by zipline even if they
	// are just set, true or false does not matter
	setHeader := func(key string, value interface{}) {
		switch v := value.(type) {
		case bool:
			if v {
				headers.Set(key, strconv.FormatBool(v))
			}
		case string:
			if v != "" {
				headers.Set(key, v)
			}
		case uint8:
			if v != 0 {
				headers.Set(key, strconv.FormatUint(uint64(v), 10))
			}
		case uint:
			if v != 0 {
				headers.Set(key, strconv.FormatUint(uint64(v), 10))
			}
		}
	}

	if o.Format.Value >= 0 {
		setHeader("Format", o.Format.String())
	}
	setHeader("Image-Compression-Percent", o.ImageCompressionPercent)
	setHeader("Expires-At", o.ExpiresAt)
	setHeader("Password", o.Password)
	setHeader("Zws", o.ZeroWidthSpace)
	setHeader("Embed", o.Embed)
	setHeader("Max-Views", o.MaxViews)
	setHeader("UploadText", o.UploadText)
	setHeader("X-Zipline-Filename", o.XZiplineFilename)
	setHeader("Override-Domain", o.OverrideDomain)
	setHeader("X-Zipline-Folder", o.XZiplineFolder)

	log.Println(headers)

	return headers
}
