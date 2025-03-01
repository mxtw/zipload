package upload

import (
	"log"
	"net/http"
	"strconv"
)

// refer to https://zipline.diced.sh/docs/guides/upload-options
type Options struct {
	DeletesAt               string
	Domain                  string
	FileExtension           string
	Filename                string
	Folder                  uint
	Format                  Format
	ImageCompressionPercent uint8
	MaxViews                uint
	OriginalName            bool
	Password                string
}

func (o Options) toHeaders() http.Header {
	headers := http.Header{}

	// use a helper function to only set headers when it makes sense to set
	// them especially bool headers will be respected by zipline even if they
	// are just set, true or false does not matter
	setHeader := func(key string, value any) {
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

	formatFlag := FormatFlag{Value: o.Format}

	setHeader("x-zipline-deletes-at", o.DeletesAt)
	setHeader("x-zipline-domain", o.Domain)
	setHeader("x-zipline-file-extension", o.FileExtension)
	setHeader("x-zipline-filename", o.Filename)
	setHeader("x-zipline-folder", o.Folder)
	setHeader("x-zipline-format", formatFlag.String())
	setHeader("x-zipline-image-compression-percent", o.ImageCompressionPercent)
	setHeader("x-zipline-max-views", o.MaxViews)
	setHeader("x-zipline-original-name", o.OriginalName)
	setHeader("x-zipline-password", o.Password)

	log.Println(headers)

	return headers
}
